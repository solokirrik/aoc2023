from collections import namedtuple

def line_to_state(line):
    if line.startswith("seed-to-soil map:"):
        return "seed-to-soil"
    if line.startswith("soil-to-fertilizer map:"):
        return "soil-to-fertilizer"
    if line.startswith("fertilizer-to-water map:"):
        return "fertilizer-to-water"
    if line.startswith("water-to-light map:"):
        return "water-to-light"
    if line.startswith("light-to-temperature map:"):
        return "light-to-temperature"
    if line.startswith("temperature-to-humidity map:"):
        return "temperature-to-humidity"
    if line.startswith("humidity-to-location map:"):
        return "humidity-to-location"
    return "numbers"

Boundaries = namedtuple('Boundaries', ['source_min', 'destination_min', 'elements'])

def build_transitions(lines):
    transitions = {}

    while len(lines) > 0:
        line = lines.pop(0)
        state = line_to_state(line)
        if state != "numbers":
            prev_state = state
            continue

        line = [int(x) for x in line.split(" ")]
        if transitions.get(prev_state, None) is None:
            transitions[prev_state] = [Boundaries(line[1], line[0], line[2])]
        else:
            transitions[prev_state].append(Boundaries(line[1], line[0], line[2]))

    return transitions

def fins_locations(seeds, transitions):
    maps = [
        "seed-to-soil",
        "soil-to-fertilizer",
        "fertilizer-to-water",
        "water-to-light",
        "light-to-temperature",
        "temperature-to-humidity",
        "humidity-to-location"
    ]
    locations = []

    for seed in seeds:
        val = seed
        new_val = None
        for tr in maps:
            for bound in transitions[tr]:
                if bound.source_min <= val <= bound.source_min + bound.elements - 1:
                    new_val = bound.destination_min + (val - bound.source_min)
                    break
                else:
                    new_val = None
            if tr == "humidity-to-location":
                if new_val is not None:
                    locations.append(new_val)
                else:
                    locations.append(val)

            if new_val is None:
                val = val
                # print("No transition", seed, val)
                continue
            if new_val is not None:
                # print("Transition", seed, val, new_val)
                val = new_val
                continue
    return locations

def task1(lines):
    seeds = {int(x):{} for x in lines.pop(0)[len("seeds: "):].split(" ")}
    transitions = build_transitions(lines)
    locations = fins_locations(seeds, transitions)

    print(min(locations), locations)

    return min(locations)
