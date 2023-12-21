from dataclasses import dataclass
from typing import List, Dict

sep = " -> "

# Multiline comment
"""
button module - push the button -> single low pulse sent to broadcaster module

% - Flip-flop modules - on or off
- initially off
If a flip-flop module receives a high pulse, it is ignored and nothing happens.
If a flip-flop module receives a low pulse, it flips between on and off.
    If it was off -> turns on -> sends high pulse
    If it was on -> turns off -> sends low pulse

& - Conjunction modules remember the type of the most recent pulse received from each of their connected input modules;
- initially default to remembering a low pulse for each input.
When a pulse is received, first updates its memory for that input
    - if high pulses for all inputs -> sends a low pulse
    - otherwise, it sends high pulse
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
class FlipFlopModule:
    name: str
    output_name: str
    is_on: bool

    def __init__(self, name: str, output: str):
        self.name = name
        self.output_name = output
        self.is_on = False

    def consume(self, input_name: str, pulse: int) -> (int, List[str]):
        if pulse == HIGH:
            return 0, []
        if self.is_on:
            self.is_on = False
            return LOW, [self.output_name]
        else:
            self.is_on = True
            return HIGH, [self.output_name]

    def __str__(self) -> str:
        return f"{self.name} -> {self.output_name}: {self.is_on}"


@dataclass
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
            return False, self.outputs

        return True, self.outputs

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


def task1():
    with open("./day20/input-example-1") as f:
        modules = build_modules(f.readlines())
        ops_queue = [Op("button", "broadcaster", False)]
        low_count, high_count = 0, 0

        for module in modules.values():
            print(module)

        while len(ops_queue):
            op = ops_queue.pop(0)
            module = modules[op.target_module_name]

            new_pulse, outputs = module.consume(op.from_module_name, op.pulse)
            for output in outputs:
                ops_queue.append(Op(module.name, output, new_pulse))

            if len(outputs) == 0:
                if new_pulse == LOW:
                    low_count += 1
                elif new_pulse == HIGH:
                    high_count += 1
        print(low_count, high_count)


def build_modules(lines: List[str]):
    modules = {}
    conjs = []
    links = {}

    for line in lines:
        line = line.strip()
        if sep not in line:
            continue
        module_name, outputs = line.split(sep)
        module_name = module_name.strip()
        outputs = outputs.strip().split(", ")
        for output in outputs:
            if output not in links:
                links[output] = set([module_name.lstrip("&%")])
                continue
            else:
                links[output].add(module_name.lstrip("&%"))

        if module_name == "broadcaster":
            module = BroadcasterModule(module_name, outputs)
        elif module_name.startswith("%"):
            module = FlipFlopModule(module_name[1:], outputs[0])
            if outputs[0] in conjs:
                conj_m = modules[outputs[0]]
                conj_m.add_input(module.name)
                modules[outputs[0]] = conj_m
        elif module_name.startswith("&"):
            module = ConjunctionModule(module_name[1:], [], outputs)
            conjs.append(module.name)
        modules[module.name] = module

    print(conjs)
    print(links)
    return modules


task1()
