package main

import (
	"fmt"
	"strconv"
)

const (
	HEADER    = "\033[95m"
	OKBLUE    = "\033[94m"
	OKCYAN    = "\033[96m"
	OKGREEN   = "\033[92m"
	YELLOW    = "\033[93m"
	RED       = "\033[91m"
	BOLD      = "\033[1m"
	UNDERLINE = "\033[4m"
	ENDC      = "\033[0m"
)

func logGreen(in string, r, c, biass int, item, nextItem byte) {
	fmt.Println(OKGREEN+"\ttrue"+ENDC, "r=", green(r), "c=", green(c), "biass", biass, string(item), "->", string(nextItem))
}

func logRed(in string, r, c, biass int, item, nextItem byte) {
	fmt.Println(RED+"\tfalse"+ENDC, "r=", red(r), "c=", red(c), "biass", biass, string(item), "->", string(nextItem))
}

func green(in int) string {
	return OKGREEN + strconv.Itoa(in) + ENDC
}

func red(in int) string {
	return RED + strconv.Itoa(in) + ENDC
}
