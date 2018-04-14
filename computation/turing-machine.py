#!/usr/bin/env python3

X_B = {
    ("B", "s1"): ("X", "R", "s2"),
    ("B", "s2"): ("B", "L", "s3"),
    ("X", "s3"): ("B", "R", "s4"),
    ("B", "s4"): ("B", "L", "s1"),
}


def print_state(tape, head, state):
    print(state, ":", "".join(tape))
    print("   ", " "*head, "^")

def execute(instructions):
    # init the machine
    # forever
        # read the value at head
        # lookup instruction for value and state
        # set the value
        # move the head
    head = 0

    # these are HACKS as we won't set the state list this and
    # the tape should be infinite
    state = "s1"
    tape = ['B', 'B']

    for _ in range(9):
        print_state(tape, head, state)

        current_val = tape[head]
        lookup = (current_val, state)

        target_val, move_dir, target_state = instructions[lookup]

        tape[head] = target_val
        state = target_state
        head += 1 if move_dir == "R" else -1

def execute_terse(instructions):
    head, state, tape = 0, "s1", ['B', 'B']

    for _ in range(9):
        print_state(tape, head, state)
        tape[head], move_dir, state = instructions[(tape[head], state)]
        head += 1 if move_dir == "R" else -1

execute_terse(X_B)
print("..........................")
execute(X_B)
