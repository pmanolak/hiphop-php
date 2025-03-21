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

#pragma once

#include <map>
#include <string>
#include <vector>

namespace apache::thrift::compiler::test {

struct check_compile_options {
  // Specifies whether to add Thrift annotation and lib directories to the
  // includes search path.
  bool add_standard_includes = true;

  // Specifies whether to use global resolution (~ via t_global_scope) when
  // resolving identifiers or only rely on the local resolution (~ via
  // program_scope).
  bool use_global_resolution = true;
};

void check_compile(
    const std::string& source,
    std::vector<std::string> args = std::vector<std::string>(),
    check_compile_options options = {});
void check_compile(
    const std::map<std::string, std::string>& name_contents_map,
    const std::string& file_name,
    std::vector<std::string> args = std::vector<std::string>(),
    check_compile_options options = {});

} // namespace apache::thrift::compiler::test
