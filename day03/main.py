from task1 import task1
from task2 import task2

if __name__ == '__main__':
    with open('./input') as f:
        schema = [[c for c in l.rstrip('\n')] for l in f.readlines()]

        out1 = task1(schema)
        print(528799 == out1, out1)

        out2 = task2(schema)
        print(84907174 == out2, out2)
