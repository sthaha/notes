#!/usr/bin/env python3

# add (11+111)BBBBB to get (11111)BBBBB

ADDER = {
# input
    ("B", "s1"): ("(", "R", "s2"),
    ("B", "s2"): ("1", "R", "s3"),
    ("B", "s3"): ("1", "R", "s4"),
    ("B", "s4"): ("+", "R", "s5"),
    ("B", "s5"): ("1", "R", "s6"),
    ("B", "s6"): ("1", "R", "s7"),
    ("B", "s7"): ("1", "R", "s8"),
    ("B", "s8"): (")", "R", "s9"),

# logic: go back until we find +

    ("B", "s9"): ("B", "L", "s10"),
    (")", "s10"): (")", "L", "s11"),
    ("1", "s11"): ("1", "L", "s11"),   # loop back

    ("+", "s11"): ("1", "R", "s12"),
    ("1", "s12"): ("1", "R", "s12"),
    (")", "s12"): ("B", "L", "s13"),
    ("1", "s13"): (")", "R", "s14"),
    ("B", "s14"): ("B", "R", "s14"),
}

# The adder adds numbers encoded as 1s e.g. 3 is encoded as 111

def print_state(tape, head, state):
    print(state.rjust(4), ":", "".join(tape))
    print("     ", " " * head, "^")

def execute_terse(instructions):
    head, state, tape = 0, "s1", ['B'] * 32

    for _ in range(22):
        print_state(tape, head, state)
        tape[head], move_dir, state = instructions[(tape[head], state)]
        head += 1 if move_dir == "R" else -1

execute_terse(ADDER)
