from task1 import task1

test_cases = [
    ["./input-example0", 4361],
    ["./input-example1", 413],
    ["input", 528799],
]

def test_task1():
    for case in test_cases:
        with open(case[0]) as f:
            schema = [[c for c in l.rstrip('\n')] for l in f.readlines()]
            result = task1(schema)
            if result != case[1]:
                print("[TEST FAILED] task1", case[0] ,"got=", result, "wanted=", case[1])
            else:
                print("[TEST PASSED] task1", case[0])

test_task1()