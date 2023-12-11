import dataclasses


@dataclasses.dataclass
class Point:
    idx: int
    x: int
    y: int


def distance(p1, p2):
    return abs(p1.x-p2.x) + abs(p1.y-p2.y)


def task1(file="./day11/input-example"):
    scale = 2
    empty_rows = scale
    empty_cols = scale

    with open(file) as f:
        grid = [list(line.strip()) for line in f.readlines()]
        galaxies = extract_galaxies(grid)
        galaxies = sparse_galaxies(
            grid,
            extract_galaxies(grid),
            empty_rows,
            empty_cols,
        )

    return sum_distances(galaxies)


def task2(file="./day11/input-example", scale=1000000):
    empty_rows = scale
    empty_cols = scale

    with open(file) as f:
        grid = [list(line.strip()) for line in f.readlines()]
        galaxies = sparse_galaxies(
            grid,
            extract_galaxies(grid),
            empty_rows,
            empty_cols,
        )

    return sum_distances(galaxies)


def sparse_galaxies(grid, galaxies, empty_rows, empty_cols):
    rows, cols = extendable(grid)

    for g in galaxies:
        g.x += (empty_cols-1)*len([c for c in cols if c < g.x])
        g.y += (empty_rows-1)*len([r for r in rows if r < g.y])

    return galaxies


def sum_distances(galaxies):
    dist_sum = 0
    for r in range(len(galaxies)):
        for c in range(len(galaxies)):
            if c <= r:
                continue
            dist_sum += distance(galaxies[r], galaxies[c])
    return dist_sum


def extendable(grid):
    to_extend_lines = []
    to_extend_cols = []
    for i, row in enumerate(grid):
        if '#' not in row:
            to_extend_lines.append(i)

    for j in range(len(grid[0])):
        col_contains_galaxy = False
        for i in range(len(grid)):
            if grid[i][j] == '#':
                col_contains_galaxy = True
                break

        if not col_contains_galaxy:
            to_extend_cols.append(j)

    return to_extend_lines, to_extend_cols


def extract_galaxies(grid):
    galaxies = []
    for i in range(len(grid)):
        for j in range(len(grid[0])):
            if grid[i][j] == '#':
                galaxies.append(Point(idx=len(galaxies), x=j, y=i))
    return galaxies


if __name__ == '__main__':
    out1v1 = task1(file="./day11/input-example")
    print(374 == out1v1, out1v1)

    out1 = task1(file="./day11/input")
    print(9233514 == out1, out1)

    out2v1 = task2(file="./day11/input-example", scale=2)
    print(374 == out2v1, out2v1)

    out2v2 = task2(file="./day11/input-example", scale=10)
    print(1030 == out2v2, out2v2)

    out2v3 = task2(file="./day11/input-example", scale=100)
    print(8410 == out2v3, out2v3)

    out2 = task2(file="./day11/input", scale=1000000)
    print(363293506944 == out2, out2)
