type x =
  | A of int option
  | B of bool
  | C of float
  | D of int ref
  | E of string
  | F of string * string

type all_nullary =
  | Cons1
  | Cons2
