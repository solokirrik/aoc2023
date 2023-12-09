from util import build_till_zeros

def task1(numbers):
    new_vals = []

    for line in numbers:
        new_vals.append(extract_history_val1(line))

    return sum(new_vals)

def extract_history_val1(line):
    all_zeros = build_till_zeros(line)

    for line in all_zeros:
        line.append(0)

    return history_val1(all_zeros)


def history_val1(all_zeros):
    for i in range(len(all_zeros)-2, -1, -1):
        line = all_zeros[i]
        line[-1] = all_zeros[i+1][-1] + line[-2]

    return all_zeros[0][-1]
