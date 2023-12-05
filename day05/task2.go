package main

import (
	"bytes"
	_ "embed"
	"log/slog"
	"math"
	"strconv"
	"sync"
)

type Map struct {
	destFrom   uint64
	sourceFrom uint64
	items      uint64
}

type Mappers [][]Map

func (m Mappers) Build(inputMaps [][]byte) Mappers {
	for index, split := range inputMaps[1:] {
		contentTransitions := bytes.Split(split, []byte("\n"))[1:]
		for i := range contentTransitions {
			parts := bytes.Split(contentTransitions[i], []byte(" "))

			destFrom, err := strconv.Atoi(string(parts[0]))
			panicOnError(err)
			sourceFrom, err := strconv.Atoi(string(parts[1]))
			panicOnError(err)
			items, err := strconv.Atoi(string(parts[2]))
			panicOnError(err)

			m[index] = append(m[index], Map{uint64(destFrom), uint64(sourceFrom), uint64(items)})
		}
	}

	return m
}

type SeedsRange struct {
	from  uint64
	items uint64
}

//go:embed input
var lines []byte

func main() {
	inputMaps := bytes.Split(lines, []byte("\n\n"))
	seedsRanges := parseSeedsRanges(inputMaps[0])

	wg := new(sync.WaitGroup)
	wg.Add(len(seedsRanges))

	sc := SeedsChecker{
		mux:         &sync.Mutex{},
		mappers:     make(Mappers, 7).Build(inputMaps),
		minLocation: math.MaxUint64,
	}

	for i := range seedsRanges {
		slog.Info(
			"Running seeds batch",
			"from", seedsRanges[i].from,
			"items", seedsRanges[i].items,
		)

		go func(i int) {
			defer wg.Done()

			from := seedsRanges[i].from
			to := seedsRanges[i].from + seedsRanges[i].items - 1

			for seedVal := from; seedVal <= to; seedVal++ {
				sc.checkSeed(seedVal, 0)
			}
		}(i)
	}

	wg.Wait()

	slog.Info("task2", "location", sc.minLocation)
}

type SeedsChecker struct {
	mux         *sync.Mutex
	mappers     Mappers
	minLocation uint64
}

func (sc *SeedsChecker) checkSeed(sourceVal uint64, index uint64) {
	nextSourceVal := sourceVal

	for _, mapper := range sc.mappers[index] {
		from := mapper.sourceFrom
		to := mapper.sourceFrom + mapper.items - 1

		if sourceVal >= from && sourceVal <= to {
			nextSourceVal = sourceVal + (mapper.destFrom - mapper.sourceFrom)
		}
	}

	if index < 6 {
		sc.checkSeed(nextSourceVal, index+1)
	} else {
		sc.mux.Lock()
		if nextSourceVal < sc.minLocation {
			sc.minLocation = nextSourceVal
		}
		sc.mux.Unlock()
	}
}

func parseSeedsRanges(inp []byte) []SeedsRange {
	seedsNumbersStr := bytes.Split(bytes.Split(inp, []byte(": "))[1], []byte(" "))
	seedInt := make([]uint64, 0, len(seedsNumbersStr))

	for _, seedNumberStr := range seedsNumbersStr {
		seedNumber, err := strconv.Atoi(string(seedNumberStr))
		panicOnError(err)
		seedInt = append(seedInt, uint64(seedNumber))
	}

	seedsBatches := make([]SeedsRange, 0, len(seedInt)/2)
	for i := 0; i < len(seedInt); i += 2 {
		seedsBatches = append(seedsBatches, SeedsRange{from: seedInt[i], items: seedInt[i+1]})
	}

	return seedsBatches
}

func panicOnError(err error) {
	if err != nil {
		panic(err)
	}
}
