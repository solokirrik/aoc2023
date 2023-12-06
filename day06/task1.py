def task1(lines):
    time_line = [int(x) for x in lines[0][1:]]
    dist_line = [int(x) for x in lines[1][1:]]

    out = 1
    for idx, round in enumerate(time_line):
        options = round_to_options(round, dist_line[idx])
        out *= options

    return out

def round_to_options(round_time, record):
    options = 0
    for push_time in range(1, round_time):
        distance = push_time * (round_time - push_time)
        if distance > record:
            options += 1

    return options
