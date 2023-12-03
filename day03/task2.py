from collections import namedtuple

# value - number adjacent to a star
# row, col - row and col coordinated of the star sign
StarNeighbour = namedtuple('DataStruct', ['value', 'row', 'col'])

def is_adjacent_to_star(mtx, r, c):
    rows = len(mtx)
    line_len = len(mtx[0])

    for i in [-1, 0, 1]:
        for j in [-1, 0, 1]:
            if (i == 0 and j == 0) or (r+i < 0 or c+j < 0 or r+i >= rows or c+j >= line_len):
                continue
            if mtx[r+i][c+j] == "*":
                return True, r+i, c+j

    return False, 0, 0

def filer_by_star(star, coll):
    return [e for e in coll if e.row == star.row and e.col == star.col]

def is_in_collection(item, coll):
    return len(filer_by_star(item, coll)) > 0

def filter_multiply(numbers_with_coords):
    numbers = []
    checked = []

    for item in numbers_with_coords:
        if is_in_collection(item, checked):
            continue

        filtered = filer_by_star(item, numbers_with_coords)
        # print(filtered)

        if len(filtered) == 2:
            checked.append(item)
            mux = 1
            for item in filtered:
                mux = mux * item.value
            numbers.append(mux)

        # print("------")

    return numbers

def numbers_with_stars(matrix):
    numbers_adj_to_stars = []

    for r in range(len(matrix)):
        number_digits = []
        star_coords = []
        # print("row", r)

        for c in range(len(matrix[0])):
            char = matrix[r][c]

            # build number from digits and fluh collections
            if not char.isdigit():
                if len(number_digits) > 0 and len(star_coords) > 0:
                    # print("".join(number_digits),"is adjacent to", matrix[star_coords[0]][star_coords[1]])
                    new_entry = StarNeighbour(int("".join(number_digits)), star_coords[0], star_coords[1])
                    numbers_adj_to_stars.append(new_entry)

                # if len(number_digits) > 0 and len(star_coords) == 0:
                    # print("".join(number_digits),"is not adjacent")

                star_coords = []
                number_digits = []

                continue

            # collect number digits and star coords
            if char.isdigit():
                number_digits.append(char)
                if len(star_coords) > 0:
                    continue
                ok, rs, cs = is_adjacent_to_star(matrix, r, c)
                if ok:
                    star_coords.append(rs)
                    star_coords.append(cs)

        # process possible last number in the row
        if len(number_digits) > 0 and len(star_coords) > 0:
            new_entry = StarNeighbour(int("".join(number_digits)), star_coords[0], star_coords[1])
            numbers_adj_to_stars.append(new_entry)

    return numbers_adj_to_stars

def task2(matrix):
    # precollect all
    numbers_adj_to_stars = numbers_with_stars(matrix)
    # print(numbers_adj_to_stars)

    # find only pairs adjusted to a single star and multiply them
    numbers = filter_multiply(numbers_adj_to_stars)
    # print(numbers)

    return sum(numbers)
