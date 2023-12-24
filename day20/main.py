from dataclasses import dataclass
from typing import List, Dict

sep = " -> "

# Multiline comment
"""
Button module - push the button -> single LOW pulse sent to Broadcaster module

% - Flip-flop modules - on or off
    Init: off
    If a flip-flop module receives a high pulse, it is ignored and nothing happens.
    If a flip-flop module receives a low pulse, it flips between on and off.
        If it was off -> turns on -> sends high pulse
        If it was on -> turns off -> sends low pulse

& - Conjunction modules remember the type of the most recent pulse received from each of their connected input modules;
    Init: default to remembering a low pulse for each input.
    When a pulse is received, first updates its memory for that input
        - if high pulses for all inputs -> sends a low pulse
        - otherwise, it sends high pulse

FINISH:
    - flip-flop modules all end up off, so pushing the button again repeats the same sequence
"""

LOW = -1
NOTHING = 0
HIGH = 1


@dataclass
class BroadcasterModule:
    name: str
    outputs: List[str]

    def __init__(self, name: str, outputs: List[str]):
        self.name = name
        self.outputs = outputs

    def consume(self, input: str, pulse: int):
        return LOW, self.outputs

    def __str__(self) -> str:
        return f"{self.name} -> {self.outputs}"


@dataclass
# Flip-flop modules - %
class FlipFlopModule:
    name: str
    output_names: List[str]
    is_on: bool

    def __init__(self, name: str, outputs: str):
        self.name = name
        self.output_names = outputs
        self.is_on = False

    def consume(self, input_name: str, pulse: int) -> (int, List[str]):
        if pulse == HIGH:
            return NOTHING, []
        if self.is_on:
            self.is_on = False
            return LOW, self.output_names
        else:
            self.is_on = True
            return HIGH, self.output_names

    def __str__(self) -> str:
        return f"{self.name} -> {self.output_names}: {self.is_on}"


@dataclass
# Conjunction modules - &
class ConjunctionModule:
    name: str
    inputs: Dict[str, bool]
    outputs: List[str]

    def __init__(self, name: str, inputs: List[str], outputs: List[str]):
        self.name = name
        self.inputs = {i: False for i in inputs}
        self.outputs = outputs

    def add_input(self, input_name):
        self.inputs[input_name] = False

    def consume(self, input_name: str, pulse: int) -> (int, List[str]):
        self.inputs[input_name] = pulse == HIGH
        if all(self.inputs.values()):
            return LOW, self.outputs

        return HIGH, self.outputs

    def __str__(self) -> str:
        return f"{self.inputs} -> {self.name} -> {self.outputs}"


@dataclass
class Op:
    from_module_name: str
    target_module_name: str
    pulse: int

    def __init__(self, from_module_name: str, target_module_name: str, pulse: int):
        self.from_module_name = from_module_name
        self.target_module_name = target_module_name
        self.pulse = pulse

    def __str__(self) -> str:
        return f"{self.from_module_name} -> {self.target_module_name} : {self.pulse}"


@dataclass
class Output:
    name: str

    def __init__(self, name: str):
        self.name = name

    def consume(self, input_name: str, pulse: int) -> (int, List[str]):
        print(f"OUTPUT: {input_name} -> {pulse}")
        return NOTHING, []


def task1():
    clicks = 1000
    low_count, high_count = 0, 0

    with open("./day20/input") as f:
        modules = build_modules(f.readlines())
        modules["output"] = Output("output")

        print("MODULES:")
        for module in modules.values():
            print(module)
        print("----\n")

        print("Clicking button...")
        for i in range(clicks):
            diffLow, diffHi = click_buttonv1(modules)
            low_count += diffLow
            high_count += diffHi

    print(low_count, high_count)

    return low_count * high_count


def task2():
    # NOT SOLVED
    clicks = 1_000_000
    intervals = []

    with open("./day20/input") as f:
        modules = build_modules(f.readlines())
        modules["output"] = Output("output")

        print("MODULES:")
        for module in modules.values():
            print(module)
        print("----\n")

        print("Clicking button...")
        for i in range(clicks):
            toGfFrom, isGf = click_buttonv2(modules)
            if isGf:
                intervals.append((i, toGfFrom))
                continue
            if i % 10000 == 0:
                print(i*100/clicks, "%", i)

        print("NOT FOUND", i)

    return 0


def click_buttonv1(modules) -> (int, int):
    low_count, high_count = 0, 0
    ops_queue = [Op("button", "broadcaster", LOW)]

    while len(ops_queue):
        op = ops_queue.pop(0)
        if op.pulse == LOW:
            low_count += 1
        elif op.pulse == HIGH:
            high_count += 1

        module = modules.get(op.target_module_name)
        if module is None:
            # print("NO MODULE", op.from_module_name,
            #       op.pulse, op.target_module_name)
            continue

        new_pulse, outputs = module.consume(op.from_module_name, op.pulse)

        if new_pulse == NOTHING:
            continue

        for output in outputs:
            ops_queue.append(Op(module.name, output, new_pulse))

    return low_count, high_count


def click_buttonv2(modules) -> (str, bool):
    ops_queue = [Op("button", "broadcaster", LOW)]

    while len(ops_queue):
        op = ops_queue.pop(0)
        module = modules.get(op.target_module_name)
        if op.target_module_name == "gf":
            if op.pulse == HIGH:
                return op.from_module_name, True

        if module is None:
            continue

        new_pulse, outputs = module.consume(op.from_module_name, op.pulse)

        if new_pulse == NOTHING:
            continue

        for output in outputs:
            ops_queue.append(Op(module.name, output, new_pulse))

    return "", False


def build_modules(lines: List[str]):
    modules = {}
    conjs = []
    link_to_from = {}

    for line in lines:
        typed_module_name, outputs = line.strip().split(sep)
        module_name = typed_module_name.lstrip("&%")
        outputs = outputs.strip().split(", ")

        for output in outputs:
            if output not in link_to_from:
                link_to_from[output] = set([module_name])
                continue
            else:
                link_to_from[output].add(module_name)

    for line in lines:
        typed_module_name, outputs = line.strip().split(sep)
        module_name = typed_module_name.lstrip("&%")
        outputs = outputs.strip().split(", ")

        if typed_module_name == "broadcaster":
            module = BroadcasterModule(typed_module_name, outputs)

        elif typed_module_name == "output":
            module = Output(typed_module_name)

        elif typed_module_name.startswith("%"):
            module = FlipFlopModule(module_name, outputs)
            for output in outputs:
                if output in conjs:
                    conj_m = modules[output]
                    conj_m.add_input(module.name)
                    modules[output] = conj_m

        elif typed_module_name.startswith("&"):
            module = ConjunctionModule(
                module_name,
                list(link_to_from[module_name]),
                outputs
            )
            conjs.append(module.name)

        modules[module.name] = module

    return modules


out1 = task1()
print("TASK1", out1 == 886347020, out1)
# task2()
