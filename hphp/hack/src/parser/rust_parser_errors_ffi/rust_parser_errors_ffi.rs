// Copyright (c) Facebook, Inc. and its affiliates.
//
// This source code is licensed under the MIT license found in the
// LICENSE file in the "hack" directory of this source tree.

// Copyright (c) Facebook, Inc. and its affiliates.
//
// This source code is licensed under the MIT license found in the
// LICENSE file in the "hack" directory of this source tree.

use std::collections::HashSet;

use bumpalo::Bump;
use ocamlrep::FromOcamlRep;
use ocamlrep::ptr::UnsafeOcamlPtr;
use ocamlrep_ocamlpool::ocaml_ffi;
use oxidized::experimental_features;
use oxidized::namespace_env::Mode;
use oxidized::parser_options::ParserOptions;
use parser_core_types::source_text::SourceText;
use parser_core_types::syntax_by_ref::positioned_syntax::PositionedSyntax;
use parser_core_types::syntax_error::SyntaxError;
use parser_core_types::syntax_tree::SyntaxTree;

// "only_for_parser_errors" because it sets only a subset of options relevant to parser errors,
// leaving the rest default
unsafe fn parser_options_from_ocaml_only_for_parser_errors(
    ocaml_opts: UnsafeOcamlPtr,
) -> (ParserOptions, (bool, bool, bool, bool)) {
    let ocaml_opts = ocaml_opts.as_usize() as *const usize;

    // Keep in sync with src/options/parserOptions.ml
    let hhvm_compat_mode = bool::from_ocaml(*ocaml_opts.add(0)).unwrap();
    let hhi_mode = bool::from_ocaml(*ocaml_opts.add(1)).unwrap();
    let codegen = bool::from_ocaml(*ocaml_opts.add(2)).unwrap();
    let mut parser_options = ParserOptions::default();

    let po_disable_lval_as_an_expression = bool::from_ocaml(*ocaml_opts.add(3)).unwrap();
    let po_disable_legacy_soft_typehints = bool::from_ocaml(*ocaml_opts.add(4)).unwrap();
    let po_const_static_props = bool::from_ocaml(*ocaml_opts.add(5)).unwrap();
    let po_disable_legacy_attribute_syntax = bool::from_ocaml(*ocaml_opts.add(6)).unwrap();
    let po_const_default_func_args = bool::from_ocaml(*ocaml_opts.add(7)).unwrap();
    let po_abstract_static_props = bool::from_ocaml(*ocaml_opts.add(8)).unwrap();
    let po_disallow_func_ptrs_in_constants = bool::from_ocaml(*ocaml_opts.add(9)).unwrap();
    let po_enable_xhp_class_modifier = bool::from_ocaml(*ocaml_opts.add(10)).unwrap();
    let po_disable_xhp_element_mangling = bool::from_ocaml(*ocaml_opts.add(11)).unwrap();
    let po_disable_xhp_children_declarations = bool::from_ocaml(*ocaml_opts.add(12)).unwrap();
    let po_const_default_lambda_args = bool::from_ocaml(*ocaml_opts.add(13)).unwrap();
    let po_allow_unstable_features = bool::from_ocaml(*ocaml_opts.add(14)).unwrap();
    let po_interpret_soft_types_as_like_types = bool::from_ocaml(*ocaml_opts.add(15)).unwrap();
    let tco_is_systemlib = bool::from_ocaml(*ocaml_opts.add(16)).unwrap();
    let po_disallow_static_constants_in_default_func_args =
        bool::from_ocaml(*ocaml_opts.add(17)).unwrap();
    let use_legacy_experimental_feature_config = bool::from_ocaml(*ocaml_opts.add(18)).unwrap();
    let po_experimental_features =
        oxidized::s_map::SMap::<experimental_features::FeatureStatus>::from_ocaml(
            *ocaml_opts.add(19),
        )
        .unwrap();
    let consider_unspecified_experimental_features_released =
        bool::from_ocaml(*ocaml_opts.add(20)).unwrap();
    let enable_class_pointer_hint = bool::from_ocaml(*ocaml_opts.add(21)).unwrap();
    let use_oxidized_by_ref_decls = bool::from_ocaml(*ocaml_opts.add(22)).unwrap();
    let use_oxidized_by_ref_decls2 = bool::from_ocaml(*ocaml_opts.add(23)).unwrap();

    parser_options.disable_lval_as_an_expression = po_disable_lval_as_an_expression;
    parser_options.disable_legacy_soft_typehints = po_disable_legacy_soft_typehints;
    parser_options.const_static_props = po_const_static_props;
    parser_options.disable_legacy_attribute_syntax = po_disable_legacy_attribute_syntax;
    parser_options.const_default_func_args = po_const_default_func_args;
    parser_options.abstract_static_props = po_abstract_static_props;
    parser_options.disallow_func_ptrs_in_constants = po_disallow_func_ptrs_in_constants;
    parser_options.enable_xhp_class_modifier = po_enable_xhp_class_modifier;
    parser_options.disable_xhp_element_mangling = po_disable_xhp_element_mangling;
    parser_options.disable_xhp_children_declarations = po_disable_xhp_children_declarations;
    parser_options.const_default_lambda_args = po_const_default_lambda_args;
    parser_options.allow_unstable_features = po_allow_unstable_features;
    parser_options.interpret_soft_types_as_like_types = po_interpret_soft_types_as_like_types;
    parser_options.is_systemlib = tco_is_systemlib;
    parser_options.disallow_static_constants_in_default_func_args =
        po_disallow_static_constants_in_default_func_args;
    parser_options.use_legacy_experimental_feature_config = use_legacy_experimental_feature_config;
    parser_options.experimental_features = po_experimental_features;
    parser_options.consider_unspecified_experimental_features_released =
        consider_unspecified_experimental_features_released;
    parser_options.enable_class_pointer_hint = enable_class_pointer_hint;
    parser_options.use_oxidized_by_ref_decls = use_oxidized_by_ref_decls;
    parser_options.use_oxidized_by_ref_decls2 = use_oxidized_by_ref_decls2;
    (
        parser_options,
        (hhvm_compat_mode, hhi_mode, codegen, tco_is_systemlib),
    )
}

ocaml_ffi! {
    fn drop_tree_positioned(ocaml_tree: usize) {
        unsafe {
            let pair = Box::from_raw(ocaml_tree as *mut (usize, usize));
            let _ = Box::from_raw(pair.0 as *mut SyntaxTree<'_, PositionedSyntax<'_>, ()>);
            let _ = Box::from_raw(pair.1 as *mut Bump);
        }
    }

    fn rust_parser_errors_positioned(
        source_text: SourceText<'_>,
        ocaml_tree: usize,
        ocaml_parser_options: UnsafeOcamlPtr,
    ) -> Vec<SyntaxError> {
        let (parser_options, (hhvm_compat_mode, hhi_mode, codegen, is_systemlib)) =
            unsafe { parser_options_from_ocaml_only_for_parser_errors(ocaml_parser_options) };
        let (tree, _arena) = unsafe {
            // A rust pointer of (&SyntaxTree, &Arena) is passed to Ocaml in rust_parser_ffi::parse,
            // Ocaml passes it back here
            // PLEASE ENSURE TYPE SAFETY MANUALLY!!!
            let pair = Box::from_raw(ocaml_tree as *mut (usize, usize));
            let tree = <SyntaxTree<'_, PositionedSyntax<'_>, ()>>::ffi_pointer_into_boxed(
                pair.0,
                &source_text,
            );
            let arena = Box::from_raw(pair.1 as *mut Bump);
            (tree, arena)
        };

        let (errors, _, _) = rust_parser_errors::parse_errors(
            &tree,
            parser_options,
            hhvm_compat_mode,
            hhi_mode,
            if codegen {
                Mode::ForCodegen
            } else {
                Mode::ForTypecheck
            },
            is_systemlib,
            HashSet::default(),
        );
        errors
    }
}
