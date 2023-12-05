from task1 import task1

if __name__ == '__main__':
    with open('./day05/input') as f:
        lines = [line.strip() for line in f.readlines() if line != "\n"]

        out1 = task1(lines)
        print(out1, 462648396 == out1)
