(*
 * Copyright (c) 2017, Facebook, Inc.
 * All rights reserved.
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the "hack" directory of this source tree.
 *
 *)

open Hh_prelude
module Env = Tast_env

class virtual iter =
  object (self)
    inherit [_] Aast.iter as super

    (* Entry point *)
    method go ctx program = self#on_list (fun () -> self#go_def ctx) () program

    method go_def ctx x = self#on_def (Env.def_env ctx x) x

    method! on_expr_ env expr_ =
      match expr_ with
      | Aast.Invalid _ -> ()
      | _ -> super#on_expr_ env expr_

    method! on_fun_ env x = super#on_fun_ (Env.restore_fun_env env x) x

    method! on_method_ env x =
      let env =
        if
          String.equal
            (snd x.Aast.m_name)
            Naming_special_names.Members.__construct
        then
          Env.set_inside_constructor env
        else if x.Aast.m_static then
          Env.set_static env
        else
          env
      in
      super#on_method_ (Env.restore_method_env env x) x

    method! on_class_var env cv =
      let env =
        if cv.Aast.cv_is_static then
          Env.set_static env
        else
          env
      in
      super#on_class_var env cv

    method! on_Assign env lhs bop rhs =
      Option.iter ~f:(self#on_bop env) bop;
      self#on_expr (Env.set_val_kind env Typing_defs.Lval) lhs;
      self#on_expr env rhs

    method! on_Is env e h =
      let env = Env.set_allow_wildcards env in
      super#on_Is env e h

    method! on_As env e =
      let env = Env.set_allow_wildcards env in
      super#on_As env e

    method! on_expression_tree
        env Aast.{ et_class; et_runtime_expr; et_free_vars = _ } =
      self#on_id env et_class;
      let env = Env.inside_expr_tree env et_class in
      self#on_expr env et_runtime_expr

    method! on_ET_Splice env e =
      let env = Env.outside_expr_tree env in
      super#on_ET_Splice env e
  end

class virtual ['state] iter_with_state =
  object (self)
    inherit [_] Aast.iter as super

    (* Entry point *)
    method go (state : 'state) ctx program =
      self#on_list (fun () -> self#go_def ctx state) () program

    method go_def ctx state x = self#on_def (Env.def_env ctx x, state) x

    method! on_expr_ env expr_ =
      match expr_ with
      | Aast.Invalid _ -> ()
      | _ -> super#on_expr_ env expr_

    method on_fun_with_env (env, state) x = super#on_fun_ (env, state) x

    method! on_fun_ (env, state) x =
      self#on_fun_with_env (Env.restore_fun_env env x, state) x

    method on_method_with_env (env, state) x = super#on_method_ (env, state) x

    method! on_method_ (env, state) x =
      let env =
        if
          String.equal
            (snd x.Aast.m_name)
            Naming_special_names.Members.__construct
        then
          Env.set_inside_constructor env
        else if x.Aast.m_static then
          Env.set_static env
        else
          env
      in
      self#on_method_with_env (Env.restore_method_env env x, state) x

    method on_class_var_with_env (env, state) cv =
      super#on_class_var (env, state) cv

    method! on_class_var (env, state) cv =
      let env =
        if cv.Aast.cv_is_static then
          Env.set_static env
        else
          env
      in
      self#on_class_var_with_env (env, state) cv

    method! on_Assign (env, state) lhs bop rhs =
      Option.iter ~f:(self#on_bop (env, state)) bop;
      self#on_expr (Env.set_val_kind env Typing_defs.Lval, state) lhs;
      self#on_expr (env, state) rhs

    method! on_Is (env, state) e h =
      let env = Env.set_allow_wildcards env in
      super#on_Is (env, state) e h

    method! on_As (env, state) e =
      let env = Env.set_allow_wildcards env in
      super#on_As (env, state) e

    method! on_expression_tree
        (env, state) Aast.{ et_class; et_runtime_expr; et_free_vars = _ } =
      self#on_id (env, state) et_class;
      let env = Env.inside_expr_tree env et_class in
      self#on_expr (env, state) et_runtime_expr

    method! on_ET_Splice (env, state) e =
      let env = Env.outside_expr_tree env in
      super#on_ET_Splice (env, state) e
  end

class virtual ['a] reduce =
  object (self)
    inherit [_] Aast.reduce as super

    (* Entry point *)
    method go ctx program : 'a =
      self#on_list (fun () -> self#go_def ctx) () program

    method go_def ctx x = self#on_def (Env.def_env ctx x) x

    method! on_expr_ env expr_ =
      match expr_ with
      | Aast.Invalid _ -> self#zero
      | _ -> super#on_expr_ env expr_

    method! on_fun_ env x = super#on_fun_ (Env.restore_fun_env env x) x

    method! on_method_ env x =
      let env =
        if
          String.equal
            (snd x.Aast.m_name)
            Naming_special_names.Members.__construct
        then
          Env.set_inside_constructor env
        else if x.Aast.m_static then
          Env.set_static env
        else
          env
      in
      super#on_method_ (Env.restore_method_env env x) x

    method! on_class_var env cv =
      let env =
        if cv.Aast.cv_is_static then
          Env.set_static env
        else
          env
      in
      super#on_class_var env cv

    method! on_Assign env lhs bop rhs =
      let op = Option.map ~f:(self#on_bop env) bop in
      let e1 = self#on_expr (Env.set_val_kind env Typing_defs.Lval) lhs in
      let e2 = self#on_expr env rhs in
      match op with
      | Some op -> self#plus e1 (self#plus op e2)
      | None -> self#plus e1 e2

    method! on_Is env e h =
      let env = Env.set_allow_wildcards env in
      super#on_Is env e h

    method! on_As env e =
      let env = Env.set_allow_wildcards env in
      super#on_As env e

    method! on_expression_tree
        env Aast.{ et_class = cls; et_runtime_expr; et_free_vars = _ } =
      let et_class = self#on_id env cls in
      let env = Env.inside_expr_tree env cls in
      let et_runtime_expr = self#on_expr env et_runtime_expr in
      self#plus et_class et_runtime_expr

    method! on_ET_Splice env e =
      let env = Env.outside_expr_tree env in
      super#on_ET_Splice env e
  end

class virtual map =
  object (self)
    inherit [_] Aast.map as super

    method on_'ex _ ex = ex

    method on_'en _ en = en

    (* Entry point *)
    method go ctx program : Tast.program =
      self#on_list (fun () -> self#go_def ctx) () program

    method go_def ctx x = self#on_def (Env.def_env ctx x) x

    method! on_expr_ env expr_ =
      match expr_ with
      | Aast.Invalid _ -> expr_
      | _ -> super#on_expr_ env expr_

    method! on_fun_ env x = super#on_fun_ (Env.restore_fun_env env x) x

    method! on_method_ env x =
      let env =
        if
          String.equal
            (snd x.Aast.m_name)
            Naming_special_names.Members.__construct
        then
          Env.set_inside_constructor env
        else if x.Aast.m_static then
          Env.set_static env
        else
          env
      in
      super#on_method_ (Env.restore_method_env env x) x

    method! on_class_var env cv =
      let env =
        if cv.Aast.cv_is_static then
          Env.set_static env
        else
          env
      in
      super#on_class_var env cv

    method! on_Assign env lhs bop rhs =
      let rhs = self#on_expr env rhs in
      let lhs = self#on_expr (Env.set_val_kind env Typing_defs.Lval) lhs in
      let bop = Option.map ~f:(self#on_bop env) bop in
      Aast.Assign (lhs, bop, rhs)

    method! on_Is env e h =
      let env = Env.set_allow_wildcards env in
      super#on_Is env e h

    method! on_As env e =
      let env = Env.set_allow_wildcards env in
      super#on_As env e

    method! on_expression_tree
        env Aast.{ et_class; et_runtime_expr; et_free_vars } =
      let et_class = self#on_id env et_class in
      let et_runtime_expr =
        let env = Env.inside_expr_tree env et_class in
        self#on_expr env et_runtime_expr
      in
      Aast.{ et_class; et_runtime_expr; et_free_vars }

    method! on_ET_Splice env e =
      let env = Env.outside_expr_tree env in
      super#on_ET_Splice env e
  end

class virtual endo =
  object (self)
    inherit [_] Aast.endo as super

    method on_'ex _ ex = ex

    method on_'en _ en = en

    (* Entry point *)
    method go ctx program = self#on_list (fun () -> self#go_def ctx) () program

    method go_def ctx x = self#on_def (Env.def_env ctx x) x

    method! on_expr_ env expr_ =
      match expr_ with
      | Aast.Invalid _ -> expr_
      | _ -> super#on_expr_ env expr_

    method! on_fun_ env x = super#on_fun_ (Env.restore_fun_env env x) x

    method! on_method_ env x =
      let env =
        if
          String.equal
            (snd x.Aast.m_name)
            Naming_special_names.Members.__construct
        then
          Env.set_inside_constructor env
        else if x.Aast.m_static then
          Env.set_static env
        else
          env
      in
      super#on_method_ (Env.restore_method_env env x) x

    method! on_class_var env cv =
      let env =
        if cv.Aast.cv_is_static then
          Env.set_static env
        else
          env
      in
      super#on_class_var env cv

    method! on_Assign env _ e1 bop e2 =
      let op' = Option.map ~f:(self#on_bop env) bop in
      let e1' = self#on_expr (Env.set_val_kind env Typing_defs.Lval) e1 in
      let e2' = self#on_expr env e2 in
      Aast.Assign (e1', op', e2')

    method! on_Is env e h =
      let env = Env.set_allow_wildcards env in
      super#on_Is env e h

    method! on_As env e =
      let env = Env.set_allow_wildcards env in
      super#on_As env e

    method! on_expression_tree
        env Aast.{ et_class; et_runtime_expr; et_free_vars } =
      let et_class = self#on_id env et_class in
      let et_runtime_expr =
        let env = Env.inside_expr_tree env et_class in
        self#on_expr env et_runtime_expr
      in
      Aast.{ et_class; et_runtime_expr; et_free_vars }

    method! on_ET_Splice env e =
      let env = Env.outside_expr_tree env in
      super#on_ET_Splice env e
  end

(** A {!handler} is an {!iter} visitor which is not in control of the iteration
    (and thus cannot change the order of the iteration or choose not to visit
    some subtrees).

    Intended to be used with {!iter_with} to aggregate many checks into a
    single pass over a TAST. *)
class type handler =
  object
    method at_class_ : Env.t -> Tast.class_ -> unit

    method at_typedef : Env.t -> Tast.typedef -> unit

    method at_gconst : Env.t -> Tast.gconst -> unit

    method at_fun_def : Env.t -> Tast.fun_def -> unit

    method at_module_def : Env.t -> Tast.module_def -> unit

    method at_method_ : Env.t -> Tast.method_ -> unit

    method at_expr : Env.t -> Tast.expr -> unit

    method at_stmt : Env.t -> Tast.stmt -> unit

    method at_fun_ : Env.t -> Tast.fun_ -> unit

    method at_Call : Env.t -> Tast.call_expr -> unit

    method at_hint : Env.t -> Aast.hint -> unit

    method at_tparam : Env.t -> Tast.tparam -> unit

    method at_user_attribute : Env.t -> Tast.user_attribute -> unit

    method at_class_const : Env.t -> Tast.class_const -> unit

    method at_class_typeconst_def : Env.t -> Tast.class_typeconst_def -> unit

    method at_xhp_child : Env.t -> Tast.xhp_child -> unit
  end

(** A {!handler} which does not need to make use of every visitation method can
    inherit from this no-op base class. *)
class virtual handler_base : handler =
  object
    method at_class_ _ _ = ()

    method at_typedef _ _ = ()

    method at_gconst _ _ = ()

    method at_fun_def _ _ = ()

    method at_module_def _ _ = ()

    method at_method_ _ _ = ()

    method at_expr _ _ = ()

    method at_stmt _ _ = ()

    method at_fun_ _ _ = ()

    method at_Call _ _ = ()

    method at_hint _ _ = ()

    method at_tparam _ _ = ()

    method at_user_attribute _ _ = ()

    method at_class_const _ _ = ()

    method at_class_typeconst_def _ _ = ()

    method at_xhp_child _ _ = ()
  end

(** Return an {!iter} visitor which invokes all of the given handlers upon
    visiting each node. *)
let iter_with (handlers : handler list) : iter =
  object
    inherit iter as super

    method! on_class_ env x =
      List.iter handlers ~f:(fun v -> v#at_class_ env x);
      super#on_class_ env x

    method! on_typedef env x =
      List.iter handlers ~f:(fun v -> v#at_typedef env x);
      super#on_typedef env x

    method! on_gconst env x =
      List.iter handlers ~f:(fun v -> v#at_gconst env x);
      super#on_gconst env x

    method! on_fun_def env x =
      List.iter handlers ~f:(fun v -> v#at_fun_def env x);
      super#on_fun_def env x

    method! on_module_def env x =
      List.iter handlers ~f:(fun v -> v#at_module_def env x);
      super#on_module_def env x

    method! on_method_ env x =
      List.iter handlers ~f:(fun v -> v#at_method_ env x);
      super#on_method_ env x

    method! on_expr env x =
      List.iter handlers ~f:(fun v -> v#at_expr env x);
      super#on_expr env x

    method! on_stmt env x =
      List.iter handlers ~f:(fun v -> v#at_stmt env x);
      super#on_stmt env x

    method! on_fun_ env x =
      List.iter handlers ~f:(fun v -> v#at_fun_ env x);
      super#on_fun_ env x

    method! on_Call env call_expr =
      List.iter handlers ~f:(fun v -> v#at_Call env call_expr);
      super#on_Call env call_expr

    method! on_hint env h =
      List.iter handlers ~f:(fun v -> v#at_hint env h);
      super#on_hint env h

    method! on_tparam env h =
      List.iter handlers ~f:(fun v -> v#at_tparam env h);
      super#on_tparam env h

    method! on_user_attribute env ua =
      List.iter handlers ~f:(fun v -> v#at_user_attribute env ua);
      super#on_user_attribute env ua

    method! on_class_const env tc =
      List.iter handlers ~f:(fun v -> v#at_class_const env tc);
      super#on_class_const env tc

    method! on_class_typeconst_def env tc =
      List.iter handlers ~f:(fun v -> v#at_class_typeconst_def env tc);
      super#on_class_typeconst_def env tc

    method! on_xhp_child env c =
      List.iter handlers ~f:(fun v -> v#at_xhp_child env c);
      super#on_xhp_child env c
  end
