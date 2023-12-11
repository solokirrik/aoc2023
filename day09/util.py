def build_till_zeros(line):
    all_zeros = [line]
    diff_is_zero = False

    while not diff_is_zero:
        target_line = all_zeros[-1]
        new_line = []
        for i in range(0, len(target_line)-1):
            diff = target_line[i+1] - target_line[i]
            new_line.append(diff)
            if diff == 0:
                diff_is_zero = True
            else:
                diff_is_zero = False
        all_zeros.append(new_line)

    return all_zeros
