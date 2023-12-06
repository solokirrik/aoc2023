from task1 import task1
from task2_naive import task2_naive, task2_reduce
from task2_binsearch import task2_binsearch

if __name__ == '__main__':
    with open('./day06/input') as f:
        lines = [line.strip().split() for line in f.readlines()]

        out1 = task1(lines)
        print(1195150 == out1, out1)

        # avg execution time: 2.56s
        out2 = task2_naive(lines)
        print(42550411 == out2, out2)

        # avg execution time: 3.84s
        out2 = task2_reduce(lines)
        print(42550411 == out2, out2)

        # avg execution time: 0.08s
        out2 = task2_binsearch(lines)
        print(42550411 == out2, out2)
