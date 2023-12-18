from dataclasses import dataclass


@dataclass
class Point:
    x: int
    y: int

    def __hash__(self):
        return hash((self.x, self.y))


def bool_to_int(b):
    return int(b)


def point_in_polygon_fast(target, polygon):
    if len(polygon) < 3:
        return False

    q_patt = [[0, 1], [3, 2]]
    end = len(polygon) - 1
    prev_point = [polygon[end].x - target.x, polygon[end].y - target.y]
    prev_q = q_patt[bool_to_int(prev_point[1] < 0)
                    ][bool_to_int(prev_point[0] < 0)]

    w = 0

    for i in range(end + 1):
        cur_point = [polygon[i].x - target.x, polygon[i].y - target.y]
        q = q_patt[bool_to_int(cur_point[1] < 0)
                   ][bool_to_int(cur_point[0] < 0)]

        if q - prev_q == -3:
            w += 1
        elif q - prev_q == 3:
            w -= 1
        elif q - prev_q == -2:
            if prev_point[0] * cur_point[1] >= prev_point[1] * cur_point[0]:
                w += 1
        elif q - prev_q == 2:
            if not prev_point[0] * cur_point[1] >= prev_point[1] * cur_point[0]:
                w -= 1

        prev_point = cur_point
        prev_q = q

    return w != 0


def applyDirection(p: Point, d: str, step: int) -> Point:
    if d == 'R':
        return Point(p.x+1, p.y)
    if d == 'L':
        return Point(p.x-1, p.y)
    if d == 'D':
        return Point(p.x, p.y+1)
    if d == 'U':
        return Point(p.x, p.y-1)

# 76387


def task1():
    with open("day18/input") as f:
        parts = [(x[0], int(x[1]), x[2].strip("()"))
                 for x in [l.rstrip().split(" ") for l in f.readlines()]]
        perimetr = [Point(0, 0)]
        perimIndex = {Point(0, 0)}

        min_x = float('inf')
        min_y = float('inf')
        max_x = float('-inf')
        max_y = float('-inf')
        for p in parts:
            for i in range(p[1]):
                last_point = perimetr[-1]
                nextPoint = applyDirection(last_point, p[0], p[1])
                perimIndex.add(nextPoint)
                if nextPoint.x > max_x:
                    max_x = nextPoint.x
                if nextPoint.y > max_y:
                    max_y = nextPoint.y
                if nextPoint.x < min_x:
                    min_x = nextPoint.x
                if nextPoint.y < min_y:
                    min_y = nextPoint.y
                perimetr.append(nextPoint)
        print(min_x, max_x, min_y, max_y, len(perimIndex), len(perimetr))
        print(perimetr)

        cubics = 0
        total_cells = (max_x-min_x+1)*(max_y-min_y+1)
        i = 0
        for y in range(min_y, max_y+1):
            for x in range(min_x, max_x+1):
                if Point(x, y) in perimIndex:
                    cubics += 1
                    continue
                if point_in_polygon_fast(Point(x, y), perimetr):
                    cubics += 1
                i += 1
                if i % 1000 == 0:
                    print(i, total_cells, i*100/total_cells)

        print(cubics)

    return cubics


out1 = task1()
print(out1 == 76387, out1)
