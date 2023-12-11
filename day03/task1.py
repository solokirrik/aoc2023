
def is_symbol(char):
    return char in "!@#$%^&*()_-+={}[]|\:;'<>,?/"


def is_adjacent_to_symbol(mtx, r, c):
    rows = len(mtx)
    line_len = len(mtx[0])

    for i in [-1, 0, 1]:
        for j in [-1, 0, 1]:
            if (i == 0 and j == 0) or (r+i < 0 or c+j < 0 or r+i >= rows or c+j >= line_len):
                continue
            if is_symbol(mtx[r+i][c+j]):
                return True
    return False


def task1(matrix):
    numbers = []

    for r in range(len(matrix)):
        number_digits = []
        is_adjacent = False
        # print("row", r)

        for c in range(len(matrix[0])):
            char = matrix[r][c]

            if not char.isdigit():
                if len(number_digits) > 0 and is_adjacent:
                    # print("".join(number_digits),"is adjacent")
                    numbers.append(int("".join(number_digits)))
                # if len(number_digits) > 0 and not is_adjacent:
                    # print("".join(number_digits),"is not adjacent")

                is_adjacent = False
                number_digits = []
                continue

            if char.isdigit():
                number_digits.append(char)
                # print("digit", char, number_digits)
                if is_adjacent_to_symbol(matrix, r, c):
                    # print("adjacent", char)
                    is_adjacent = True
            # else:
            #     print("symbol", char, number_digits)

        if len(number_digits) > 0 and is_adjacent:
            # print("".join(number_digits),"is adjacent")
            numbers.append(int("".join(number_digits)))

    return sum(numbers)
