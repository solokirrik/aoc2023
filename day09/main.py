from copy import deepcopy

from task1 import task1
from task2 import task2

if __name__ == '__main__':
    with open('./day09/input') as f:
        numbers = [[int(i) for i in x]
                   for x in [line.strip().split() for line in f.readlines()]]

        out1 = task1(deepcopy(numbers))
        print(out1 == 1641934234, out1)

        out2 = task2(deepcopy(numbers))
        print(out2 == 975, out2)
