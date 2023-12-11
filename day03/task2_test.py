from task2 import task2

test_cases = [
    ["./input-example0", 467835],
    ["./input-example1", 6756],
    ["input", 84907174],
]


def test_task2():
    for case in test_cases:
        with open(case[0]) as f:
            schema = [[c for c in l.rstrip('\n')] for l in f.readlines()]
            result = task2(schema)
            if result != case[1]:
                print("[TEST FAILED] task2", case[0],
                      "got=", result, "wanted=", case[1])
            else:
                print("[TEST PASSED] task2", case[0])


test_task2()
