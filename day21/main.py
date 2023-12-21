def task1():
    with open("day21/input") as f:
        mtx = [[c for c in line] for line in f.readlines()]
        start_y, start_x = find_start(mtx)
        step_points = [[(start_y, start_x)]]
        for i in range(0, 64):
            points = step_points[-1]
            print(len(points))
            step_points.append([])
            for point in points:
                possibilities = childs(mtx, point[0], point[1])
                for possibilitie in possibilities:
                    if possibilitie not in step_points[-1]:
                        step_points[-1].append(possibilitie)

        print(start_y, start_x, len(step_points[-1]))


def find_start(mtx) -> (int, int):
    for r, row in enumerate(mtx):
        for c, char in enumerate(row):
            if char == "S":
                return r, c


def childs(mtx, y, x):
    moves = [(y + 1, x), (y - 1, x), (y, x + 1), (y, x - 1)]
    out = []
    max_y = len(mtx)-1
    max_x = len(mtx[0])-1
    for move in moves:
        if (move[0] >= 0 and move[0] <= max_y) and \
            (move[1] >= 0 and move[1] <= max_x) and \
                mtx[move[0]][move[1]] != "#":
            out.append(move)
    return out


out1 = task1()
print(3615 == out1, out1)
