from util import build_till_zeros


def task2(numbers):
    new_vals = []

    for line in numbers:
        new_vals.append(extract_history_val2(line))

    return sum(new_vals)


def extract_history_val2(line):
    all_zeros = build_till_zeros(line)

    for line in all_zeros:
        line.insert(0, 0)

    return history_val2(all_zeros)


def history_val2(all_zeros):
    for i in range(len(all_zeros)-2, -1, -1):
        line = all_zeros[i]
        line[0] = line[1] - all_zeros[i+1][0]

    return all_zeros[0][0]
