(*
 * Copyright (c) 2015, Facebook, Inc.
 * All rights reserved.
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the "hack" directory of this source tree.
 *
 *)

(** Type-safe version of the in channel in Stdlib. *)
type 'a in_channel

(** Type-safe version of the out channel in Stdlib. *)
type 'a out_channel

type ('in_, 'out) channel_pair = 'in_ in_channel * 'out out_channel

val to_channel :
  'a out_channel ->
  ?flags:Marshal.extern_flags list ->
  ?flush:bool ->
  'a ->
  unit

val from_channel : 'a in_channel -> 'a

val flush : 'a out_channel -> unit

(* This breaks the type safety, but is necessary in order to allow select() *)
val descr_of_in_channel : 'a in_channel -> Unix.file_descr

val descr_of_out_channel : 'a out_channel -> Unix.file_descr

val cast_in : 'a in_channel -> Stdlib.in_channel

val cast_out : 'a out_channel -> Stdlib.out_channel

val close_out : 'a out_channel -> unit

val output_string : 'a out_channel -> string -> unit

val close_in : 'a in_channel -> unit

val input_char : 'a in_channel -> char

val input_value : 'a in_channel -> 'b

(* Spawning new process *)

(* In the absence of 'fork' on Windows, its usage must be restricted
   to Unix specifics parts.

   This module provides a mechanism to "spawn" new instance of the
   current program, but with a custom entry point (e.g. workers,
   DfindServer, ...). Then, alternate entry points should not depend
   on global references that may not have been (re)initialised in the
   new process.

   All required data must be passed through the typed channels.
   associated to the spawned process.
*)

(** Alternate entry points *)
type ('param, 'input, 'output) entry

(** Alternate entry points must be registered at toplevel, i.e.
     every call to `Daemon.register_entry_point` must have been
     evaluated when `Daemon.check_entry_point` is called at the
     beginning of `ServerMain.start`. *)
val register_entry_point :
  string ->
  ('param -> ('input, 'output) channel_pair -> unit) ->
  ('param, 'input, 'output) entry

val name_of_entry : ('param, 'input, 'output) entry -> string

(** Handler upon spawn and forked process. *)
type ('in_, 'out) handle = {
  channels: ('in_, 'out) channel_pair;
  pid: int;
}

(** for unit tests *)
val devnull : unit -> ('a, 'b) handle

val fd_of_path : string -> Unix.file_descr

val null_fd : unit -> Unix.file_descr

(** Fork and run a function that communicates via the typed channels. Doesn't work with start_memtracing. *)
val fork_FOR_TESTING_ON_UNIX_ONLY :
  ?channel_mode:[ `pipe | `socket ] ->
  (* Where the daemon's output should go *)
  Unix.file_descr * Unix.file_descr ->
  ('param -> ('input, 'output) channel_pair -> unit) ->
  'param ->
  ('output, 'input) handle

(** Spawn a new instance of the current process, and execute the
    alternate entry point. *)
val spawn :
  ?channel_mode:[ `pipe | `socket ] ->
  ?name:string ->
  (* Where the daemon's input and output should go *)
  Unix.file_descr * Unix.file_descr * Unix.file_descr ->
  ('param, 'input, 'output) entry ->
  'param ->
  ('output, 'input) handle

(** Close the typed channels associated to a 'spawned' child. *)
val close : ('a, 'b) handle -> unit

(** Force quit the spawned child process and close the associated typed channels. *)
val force_quit : ('a, 'b) handle -> unit

(** This will start the current process memtracing to the named file. It is suggested
that any process which is interesting to trace should call this function near the start
of its lifetime, so long as `--config memtrace_dir` is set. Callers are responsible
for constructing a filename that's in that directory and is unique with respect to
all other filenames used by the hh_server instance. The conventional suffix is ".ctf".
Recent versions of the memtrace library will suppress tracing after fork, but in
case we don't have the most recent version, then better not to fork!
However, it's okay to call Daemon.spawn while memtracing. *)
val start_memtracing : string -> unit

(** Main function, that executes an alternate entry point.
    It should be called only once, just before the main entry point.
    This function does not return when a custom entry point is selected. *)
val check_entry_point : unit -> unit
