(*
 * Copyright (c) 2015, Facebook, Inc.
 * All rights reserved.
 *
 * This source code is licensed under the MIT license found in the
 * LICENSE file in the "hack" directory of this source tree.
 *
 *)

open Hh_prelude

(****************************************************************************)
(* Moduling Making buckets.
 * When we parallelize, we need to create "buckets" of tasks for the
 * workers.
 * Given a list of files, we want to split it up into buckets such that
 * every worker is busy long enough. If the bucket is too big, it hurts
 * load balancing, if it is too small, the overhead in synchronization time
 * hurts *)
(****************************************************************************)

type 'a bucket =
  | Job of 'a
  | Wait
  | Done

let is_done = function
  | Done -> true
  | Wait
  | Job _ ->
    false

type 'a next = unit -> 'a bucket

let default_max_size = 500

let calculate_bucket_size ~num_jobs ~num_workers ?max_size () =
  let max_size = Option.value max_size ~default:default_max_size in
  if num_jobs < num_workers * max_size then
    max 1 (1 + (num_jobs / num_workers))
  else
    max_size

(** [make_ progress_fn bucket_size jobs] returns a function
  able to provide the next bucket of jobs from [jobs] *)
let make_ progress_fn bucket_size (jobs : 'job array) : unit -> 'job list =
  let i = ref 0 in
  fun () ->
    let bucket_size = min (Array.length jobs - !i) bucket_size in
    progress_fn ~start:!i ~length:bucket_size;
    let result = Array.sub jobs ~pos:!i ~len:bucket_size in
    i := bucket_size + !i;
    Array.to_list result

let make_list ~num_workers ?progress_fn ?max_size (jobs : 'job list) :
    unit -> 'job list =
  let progress_fn =
    Option.value ~default:(fun ~total:_ ~start:_ ~length:_ -> ()) progress_fn
  in
  let jobs = Array.of_list jobs in
  let bucket_size =
    calculate_bucket_size
      ~num_jobs:(Array.length jobs)
      ~num_workers
      ?max_size
      ()
  in
  make_ (progress_fn ~total:(Array.length jobs)) bucket_size jobs

let of_list = function
  | [] -> Done
  | wl -> Job wl

let make ~num_workers ?progress_fn ?max_size jobs =
  let maker = make_list ~num_workers ?progress_fn ?max_size jobs in
  (fun () -> of_list (maker ()))

type 'a of_n = {
  work: 'a;
  bucket: int;
  total: int;
}

let make_n_buckets ~buckets ~split =
  let next_bucket = ref 0 in
  fun () ->
    let current = !next_bucket in
    incr next_bucket;
    if current < buckets then
      Job { work = split ~bucket:current; bucket = current; total = buckets }
    else
      Done

let map t ~f =
  match t with
  | Job x -> Job (f x)
  | Done -> Done
  | Wait -> Wait
