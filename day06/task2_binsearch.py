def task2_binsearch(lines):
    round_time = int("".join(lines[0][1:]))
    round_dist = int("".join(lines[1][1:]))

    options = round_to_options_binsearch(round_time, round_dist)

    return options


def round_to_options_binsearch(round_time, record):
    min_time = find_min_push_time(round_time, record)
    max_time = find_max_push_time(round_time, record)

    return max_time-min_time+1


def find_min_push_time(round_time, record):
    min_push_time = 1
    max_push_time = round_time
    options_agg = []

    while max_push_time - min_push_time > 1:
        push_time = (max_push_time + min_push_time) // 2
        distance = push_time * (round_time - push_time)
        if distance > record:
            max_push_time = push_time
            options_agg.append(push_time)
        elif distance < record:
            min_push_time = push_time
        else:
            break

    return min(options_agg)


def find_max_push_time(round_time, record):
    min_push_time = 1
    max_push_time = round_time
    options_agg = []

    while max_push_time - min_push_time > 1:
        push_time = (max_push_time + min_push_time) // 2
        distance = push_time * (round_time - push_time)
        if distance > record:
            min_push_time = push_time
            options_agg.append(push_time)
        elif distance < record:
            max_push_time = push_time
        else:
            break

    return max(options_agg)
