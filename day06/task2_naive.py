from functools import reduce

from task1 import round_to_options

def task2_naive(lines):
    round_time = int("".join(lines[0][1:]))
    round_dist = int("".join(lines[1][1:]))

    options = round_to_options(round_time, round_dist)

    return options

def task2_reduce(lines):
    round_time = int("".join(lines[0][1:]))
    round_dist = int("".join(lines[1][1:]))

    return round_to_options_reduce(round_time, round_dist)

def round_to_options_reduce(round_time, record):
    return reduce(
        lambda count, push_time: count + (push_time * (round_time - push_time) > record),
        range(1, round_time),
        0
    )
