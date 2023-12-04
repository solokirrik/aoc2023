def task1(lines):
    result = 0

    for line in lines:
        common_items = count_common_cards(line)
        if len(common_items) == 0:
            continue
        bonus = 2**(len(common_items) - 1)
        result += bonus

    return result

def count_common_cards(line):
    cards = line.split(": ")[1].split(" | ")
    winning_cards = [x.strip() for x in cards[0].split(" ")if x != ""]
    my_cards = [x.strip() for x in cards[1].split(" ") if x != ""]
    common_items = [num for num in my_cards if num in winning_cards]

    return common_items

def task2(lines):
    result = 0

    cards = {}
    for i in range(len(lines)):
        cards[i+1] = 1

    for i, line in enumerate(lines):
        card_copies = cards[i+1]
        common_items = count_common_cards(line)
        for k in range(card_copies):
            if len(common_items) > 0:
                for j in range(1, len(common_items)+1):
                    if cards.get(i+1+j) is not None:
                        cards[i+1+j] += 1


    for v in cards.values():
        result += v

    return result