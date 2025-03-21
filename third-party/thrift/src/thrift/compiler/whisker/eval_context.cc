/*
 * Copyright (c) Meta Platforms, Inc. and affiliates.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

#include <thrift/compiler/whisker/eval_context.h>

#include <algorithm>
#include <iterator>

#include <fmt/core.h>

namespace w = whisker::make;

namespace whisker {

namespace {

object::ptr find_property(
    diagnostics_engine& diags,
    const object::ptr& self,
    const ast::identifier& identifier) {
  using result = object::ptr;
  return self->visit(
      [](null) -> result { return nullptr; },
      [](i64) -> result { return nullptr; },
      [](f64) -> result { return nullptr; },
      [](const string&) -> result { return nullptr; },
      [](boolean) -> result { return nullptr; },
      [](const array&) -> result { return nullptr; },
      [&](const native_object::ptr& o) -> result {
        if (auto map_like = o->as_map_like()) {
          return map_like->lookup_property(identifier.name);
        }
        return nullptr;
      },
      [](const native_function::ptr&) -> result { return nullptr; },
      [&](const native_handle<>& h) -> result {
        // Recurse through the prototype chain until the first matching
        // descriptor is found. This is similar to JavaScript:
        //   https://developer.mozilla.org/en-US/docs/Web/JavaScript/Guide/Inheritance_and_the_prototype_chain
        prototype<>::ptr proto = h.proto();
        while (proto != nullptr) {
          if (auto* descriptor = proto->find_descriptor(identifier.name)) {
            return detail::variant_match(
                *descriptor,
                [&](const prototype<>::property& prop) -> object::ptr {
                  return prop.function->invoke(native_function::context{
                      identifier.loc,
                      diags,
                      self,
                      {} /* positional arguments */,
                      {} /* named arguments */,
                  });
                },
                [&](const prototype<>::fixed_object& fixed) -> object::ptr {
                  return fixed.value;
                });
          }
          proto = proto->parent();
        }
        return nullptr;
      },
      [&](const map& m) -> result {
        if (auto it = m.find(identifier.name); it != m.end()) {
          return manage_as_static(it->second);
        }
        return nullptr;
      });
}

/**
 * A class representing the bag of properties at the global scope (even before
 * the root scope).
 *
 * This could be a w::map but for debugging purposes, a native_object with a
 * custom print_to function is beneficial.
 */
class global_scope_object
    : public native_object,
      public native_object::map_like,
      public std::enable_shared_from_this<global_scope_object> {
 public:
  explicit global_scope_object(map properties)
      : properties_(std::move(properties)) {}

  native_object::map_like::ptr as_map_like() const override {
    return shared_from_this();
  }

  object::ptr lookup_property(std::string_view identifier) const override {
    if (auto property = properties_.find(identifier);
        property != properties_.end()) {
      return manage_as_static(property->second);
    }
    return nullptr;
  }

  void print_to(tree_printer::scope scope, const object_print_options& options)
      const override {
    scope.println("<global scope> (size={})", properties_.size());
    for (const auto& [key, value] : properties_) {
      auto element_scope = scope.open_transparent_property();
      element_scope.println("'{}'", key);
      whisker::print_to(value, element_scope.open_node(), options);
    }
  }

 private:
  map properties_;
};

} // namespace

object::ptr eval_context::lexical_scope::lookup_property(
    diagnostics_engine& diags, const ast::identifier& identifier) {
  if (auto local = locals_.find(identifier.name); local != locals_.end()) {
    return local->second;
  }
  return find_property(diags, this_ref_, identifier);
}

eval_context::eval_context(diagnostics_engine& diags, object::ptr globals)
    : diags_(diags),
      global_scope_(std::move(globals)),
      stack_({lexical_scope(global_scope_)}) {}

eval_context::eval_context(diagnostics_engine& diags, map globals)
    : eval_context(
          diags,
          manage_owned<object>(
              w::make_native_object<global_scope_object>(std::move(globals)))) {
}

/* static */ eval_context eval_context::with_root_scope(
    diagnostics_engine& diags, object::ptr root_scope, map globals) {
  eval_context result{diags, std::move(globals)};
  result.push_scope(root_scope);
  return result;
}

eval_context::~eval_context() noexcept = default;

std::size_t eval_context::stack_depth() const {
  // The global scope is always on the stack but should not count towards the
  // depth (it's at depth 0).
  return stack_.size() - 1;
}

void eval_context::push_scope(object::ptr object) {
  stack_.emplace_back(std::move(object));
}

void eval_context::pop_scope() {
  // The global scope cannot be popped.
  assert(stack_depth() > 0);
  stack_.pop_back();
}

const object::ptr& eval_context::global_scope() const {
  return global_scope_;
}

expected<std::monostate, eval_name_already_bound_error>
eval_context::bind_local(std::string name, object::ptr value) {
  assert(!stack_.empty());
  if (auto [_, inserted] =
          stack_.back().locals().insert(std::pair{name, std::move(value)});
      !inserted) {
    return unexpected(eval_name_already_bound_error(std::move(name)));
  }
  return {};
}

expected<eval_context::lookup_result, eval_context::lookup_error>
eval_context::lookup_object(const ast::variable_lookup& lookup) {
  assert(!stack_.empty());
  using result = expected<lookup_result, lookup_error>;

  return detail::variant_match(
      lookup.chain,
      [&](ast::variable_lookup::this_ref) -> result {
        return lookup_result::without_parent(stack_.back().this_ref());
      },
      [&](const std::vector<ast::identifier>& path) -> result {
        const auto make_searched_scopes = [&]() -> std::vector<object::ptr> {
          std::vector<object::ptr> result;
          result.reserve(stack_.size());
          // Searching happened in reverse order.
          std::transform(
              stack_.rbegin(),
              stack_.rend(),
              std::back_inserter(result),
              [](const auto& scope) { return scope.this_ref(); });
          return result;
        };

        using component_iterator = std::vector<ast::identifier>::const_iterator;
        const auto make_success_path =
            [](component_iterator begin,
               component_iterator end) -> std::vector<std::string> {
          std::vector<std::string> result;
          result.reserve(std::distance(begin, end));
          std::transform(
              begin,
              end,
              std::back_inserter(result),
              [](const auto& component) { return component.name; });
          return result;
        };

        object::ptr current = nullptr;
        // Crawl up through the scope stack since names can be shadowed.
        for (auto scope = stack_.rbegin(); scope != stack_.rend(); ++scope) {
          try {
            if (auto result = scope->lookup_property(diags_, path.front())) {
              current = result;
              break;
            }
          } catch (const native_object::fatal_error& err) {
            return unexpected(eval_scope_lookup_error(
                path.front().name,
                make_searched_scopes(),
                err.what() /* cause */));
          }
        }

        if (current == nullptr) {
          return unexpected(eval_scope_lookup_error(
              path.front().name, make_searched_scopes()));
        }
        object::ptr parent = manage_as_static(whisker::make::null);

        for (auto component = std::next(path.begin()); component != path.end();
             ++component) {
          try {
            object::ptr next = find_property(diags_, current, *component);
            if (next == nullptr) {
              return unexpected(eval_property_lookup_error(
                  current, /* missing_from */
                  make_success_path(path.begin(), component),
                  component->name /* missing_name */));
            }
            parent = current;
            current = next;
          } catch (const native_object::fatal_error& err) {
            return unexpected(eval_property_lookup_error(
                current, /* missing_from */
                make_success_path(path.begin(), component),
                component->name /* missing_name */,
                err.what() /* cause */));
          }
        }

        assert(current != nullptr);
        assert(parent != nullptr);
        return lookup_result{current, parent};
      });
}

} // namespace whisker
