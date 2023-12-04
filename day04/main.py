from day04.task import task1, task2

if __name__ == '__main__':
    with open('./day04/input') as f:
        lines = f.readlines()

        out1 = task1(lines)
        print(out1)

        out2 = task2(lines)
        print(out2)
