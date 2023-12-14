package main

import (
	"fmt"
	"math"
	"sort"
	"strconv"
	"time"
)

var globalMtx [][]byte

func Task2v2(mtx [][]byte, cycles int) int {
	globalMtx = mtx
	squareRocks, roundRocks := rocks(mtx)
	rowsSquareRocks, columnsSquareRocks := buildLimits(squareRocks)
	rowsRoundRocks, columnsRoundRocks := initRoundRocksIndexes(roundRocks)

	start := time.Now()
	rows := len(mtx) - 1
	columns := len(mtx[0]) - 1
	cache := make(map[string]int)
	fmt.Println("ROWS:", rows, "COLUMNS:", columns)

	for i := 0; i < cycles; i++ {
		if i%100 == 0 {
			since := time.Since(start).String()
			fmt.Println("%", math.Round(float64(i)/float64(cycles)*100*10000)/10000,
				"| time", since[0:min(10, len(since))],
				"| iteration", i, "of", cycles)
		}

		for r := range roundRocks {
			rock := roundRocks[r]
			moveRockNorth(rock.X, rock, rowsSquareRocks, columnsSquareRocks, rowsRoundRocks, columnsRoundRocks)
		}
		resetRocks(roundRocks)

		for r := range roundRocks {
			rock := roundRocks[r]
			moveRockWest(rock.Y, rock, rowsSquareRocks, columnsSquareRocks, rowsRoundRocks, columnsRoundRocks)
		}
		resetRocks(roundRocks)

		for r := range roundRocks {
			rock := roundRocks[r]
			moveRockSouth(rock.X, rows, rock, rowsSquareRocks, columnsSquareRocks, rowsRoundRocks, columnsRoundRocks)
		}
		resetRocks(roundRocks)

		for r := range roundRocks {
			rock := roundRocks[r]
			moveRockEast(rock.Y, columns, rock, rowsSquareRocks, columnsSquareRocks, rowsRoundRocks, columnsRoundRocks)
		}
		resetRocks(roundRocks)

		if _, ok := cache[rocksHash(roundRocks)]; ok {
			i = cycles - (cycles-i)%(i-cache[rocksHash(roundRocks)])
		}

		cache[rocksHash(roundRocks)] = i
	}

	return loadOnNorthBeamV2(len(mtx), roundRocks)
}

func rocksHash(rocks []*Rock) string {
	out := ""
	for r := range rocks {
		out += strconv.Itoa(rocks[r].X) + "," + strconv.Itoa(rocks[r].Y) + "|"
	}

	return out
}

func resetRocks(roundRocks []*Rock) {
	for i := range roundRocks {
		roundRocks[i].alreadyMoved = false
	}
}

func moveRockNorth(c int,
	rock *Rock,
	rowsSquareRocks, columnsSquareRocks,
	rowsRoundRocks, columnsRoundRocks map[int][]*Rock,
) {
	if rock.Y == 0 {
		rock.alreadyMoved = true
		return
	}

	nextSquare := getNextVertical(c, rock, columnsSquareRocks, true)
	nextRound := getNextVertical(c, rock, columnsRoundRocks, true)

	if nextSquare == nil && nextRound == nil {
		migrateToNewRow(c, rock, 0, rowsRoundRocks, columnsRoundRocks)
		rock.alreadyMoved = true
		return
	}

	if nextSquare != nil && nextRound != nil &&
		nextSquare.Y > nextRound.Y {
		nextRound = nil
	}

	if nextSquare != nil && nextRound == nil {
		migrateToNewRow(c, rock, nextSquare.Y+1, rowsRoundRocks, columnsRoundRocks)
		rock.alreadyMoved = true
		return
	}

	if !nextRound.alreadyMoved {
		moveRockNorth(c, nextRound, rowsSquareRocks, columnsSquareRocks, rowsRoundRocks, columnsRoundRocks)
	}

	migrateToNewRow(c, rock, nextRound.Y+1, rowsRoundRocks, columnsRoundRocks)
	rock.alreadyMoved = true
}

func moveRockSouth(colIdx, maxRow int,
	rock *Rock,
	rowsSquareRocks, columnsSquareRocks,
	rowsRoundRocks, columnsRoundRocks map[int][]*Rock,
) {
	if rock.Y == maxRow {
		rock.alreadyMoved = true
		return
	}

	nextSquare := getNextVertical(colIdx, rock, columnsSquareRocks, false)
	nextRound := getNextVertical(colIdx, rock, columnsRoundRocks, false)

	if nextSquare == nil && nextRound == nil {
		migrateToNewRow(colIdx, rock, maxRow, rowsRoundRocks, columnsRoundRocks)
		rock.alreadyMoved = true
		return
	}

	if nextSquare != nil && nextRound != nil &&
		nextSquare.Y < nextRound.Y {
		nextRound = nil
	}

	if nextSquare != nil && nextRound == nil {
		migrateToNewRow(colIdx, rock, nextSquare.Y-1, rowsRoundRocks, columnsRoundRocks)
		rock.alreadyMoved = true
		return
	}

	if !nextRound.alreadyMoved {
		moveRockSouth(colIdx, maxRow, nextRound, rowsSquareRocks, columnsSquareRocks, rowsRoundRocks, columnsRoundRocks)
	}
	migrateToNewRow(colIdx, rock, nextRound.Y-1, rowsRoundRocks, columnsRoundRocks)
	rock.alreadyMoved = true
}

func moveRockEast(rowIdx, maxColumn int,
	rock *Rock,
	rowsSquareRocks, columnsSquareRocks,
	rowsRoundRocks, columnsRoundRocks map[int][]*Rock,
) {
	if rock.X == maxColumn {
		rock.alreadyMoved = true
		return
	}

	nextSquare := getNextHorizontal(rowIdx, rock, rowsSquareRocks, false)
	nextRound := getNextHorizontal(rowIdx, rock, rowsRoundRocks, false)

	if nextSquare == nil && nextRound == nil {
		migrateToNewColumn(rowIdx, rock, maxColumn, rowsRoundRocks, columnsRoundRocks)
		rock.alreadyMoved = true
		return
	}

	if nextSquare != nil && nextRound != nil &&
		nextSquare.X < nextRound.X {
		nextRound = nil
	}

	if nextSquare != nil && nextRound == nil {
		migrateToNewColumn(rowIdx, rock, nextSquare.X-1, rowsRoundRocks, columnsRoundRocks)
		rock.alreadyMoved = true
		return
	}

	if !nextRound.alreadyMoved {
		moveRockEast(rowIdx, maxColumn, nextRound, rowsSquareRocks, columnsSquareRocks, rowsRoundRocks, columnsRoundRocks)
	}
	migrateToNewColumn(rowIdx, rock, nextRound.X-1, rowsRoundRocks, columnsRoundRocks)
	rock.alreadyMoved = true
}

func moveRockWest(rowIdx int,
	rock *Rock,
	rowsSquareRocks, columnsSquareRocks,
	rowsRoundRocks, columnsRoundRocks map[int][]*Rock,
) {
	if rock.X == 0 {
		rock.alreadyMoved = true
		return
	}

	nextSquare := getNextHorizontal(rowIdx, rock, rowsSquareRocks, true)
	nextRound := getNextHorizontal(rowIdx, rock, rowsRoundRocks, true)

	// |..O <-
	if nextSquare == nil && nextRound == nil {
		migrateToNewColumn(rowIdx, rock, 0, rowsRoundRocks, columnsRoundRocks)
		rock.alreadyMoved = true
		return
	}

	// |..O..#..O <-
	if nextSquare != nil && nextRound != nil &&
		nextSquare.X > nextRound.X {
		nextRound = nil
	}

	// |...#...O <-
	if nextSquare != nil && nextRound == nil {
		migrateToNewColumn(rowIdx, rock, nextSquare.X+1, rowsRoundRocks, columnsRoundRocks)
		rock.alreadyMoved = true
		return
	}

	// |..O..O..O <-
	if !nextRound.alreadyMoved {
		moveRockWest(rowIdx, nextRound, rowsSquareRocks, columnsSquareRocks, rowsRoundRocks, columnsRoundRocks)
	}
	migrateToNewColumn(rowIdx, rock, nextRound.X+1, rowsRoundRocks, columnsRoundRocks)
	rock.alreadyMoved = true
}

func getNextVertical(idx int, rock *Rock, axisIndex map[int][]*Rock, isNextLower bool) *Rock {
	var next *Rock

	rocks := axisIndex[idx]
	for i := 0; i < len(rocks); i++ {
		if isNextLower {
			if rocks[i].Y > rock.Y {
				break
			}
			if rocks[i].Y != rock.Y {
				next = rocks[i]
			}
		} else {
			if rocks[i].Y != rock.Y {
				next = rocks[i]
			}
			if rocks[i].Y > rock.Y {
				break
			}
		}
	}

	if next != nil && (next.Y == rock.Y || (!isNextLower && next.Y < rock.Y)) {
		return nil
	}

	return next
}

func getNextHorizontal(idx int, rock *Rock, axisIndex map[int][]*Rock, isNextLower bool) *Rock {
	var next *Rock

	rocks := axisIndex[idx]
	for i := 0; i < len(rocks); i++ {
		if isNextLower {
			if rocks[i].X > rock.X {
				break
			}
			if rocks[i].X != rock.X {
				next = rocks[i]
			}
		} else {
			if rocks[i].X != rock.X {
				next = rocks[i]
			}
			if rocks[i].X > rock.X {
				break
			}
		}
	}

	if next != nil && (next.X == rock.X || (!isNextLower && next.X < rock.X)) {
		return nil
	}

	return next
}

func migrateToNewRow(c int, rock *Rock, newRow int, rowsRoundRocks, columnsRoundRocks map[int][]*Rock) {
	// remove rock from old row
	for i := 0; i < len(rowsRoundRocks[rock.Y]); i++ {
		// fmt.Println(i, len(rowsRoundRocks[rock.Y]))
		if rowsRoundRocks[rock.Y][i].X == rock.X && rowsRoundRocks[rock.Y][i].Y == rock.Y {
			rowsRoundRocks[rock.Y][i] = rowsRoundRocks[rock.Y][len(rowsRoundRocks[rock.Y])-1]
			rowsRoundRocks[rock.Y] = rowsRoundRocks[rock.Y][:len(rowsRoundRocks[rock.Y])-1]
			sort.Slice(rowsRoundRocks[rock.Y], func(i, j int) bool {
				return rowsRoundRocks[rock.Y][i].X < rowsRoundRocks[rock.Y][j].X
			})
		}
	}

	// update rock position in column index
	rock.Y = newRow
	sort.Slice(columnsRoundRocks[c], func(i, j int) bool {
		return columnsRoundRocks[c][i].Y < columnsRoundRocks[c][j].Y
	})

	// update rock position in row index
	rowsRoundRocks[rock.Y] = append(rowsRoundRocks[rock.Y], rock)
	sort.Slice(rowsRoundRocks[rock.Y], func(i, j int) bool {
		return rowsRoundRocks[rock.Y][i].X < rowsRoundRocks[rock.Y][j].X
	})
}

func migrateToNewColumn(r int, rock *Rock, newColumn int, rowsRoundRocks, columnsRoundRocks map[int][]*Rock) {
	// remove rock from old column
	for i := 0; i < len(columnsRoundRocks[rock.X]); i++ {
		if columnsRoundRocks[rock.X][i].X == rock.X && columnsRoundRocks[rock.X][i].Y == rock.Y {
			columnsRoundRocks[rock.X][i] = columnsRoundRocks[rock.X][len(columnsRoundRocks[rock.X])-1]
			columnsRoundRocks[rock.X] = columnsRoundRocks[rock.X][:len(columnsRoundRocks[rock.X])-1]
			sort.Slice(columnsRoundRocks[rock.X], func(i, j int) bool {
				return columnsRoundRocks[rock.X][i].Y < columnsRoundRocks[rock.X][j].Y
			})
		}
	}

	// update rock position in row index
	rock.X = newColumn
	sort.Slice(rowsRoundRocks[r], func(i, j int) bool {
		return rowsRoundRocks[r][i].X < rowsRoundRocks[r][j].X
	})

	// update rock position in column index
	columnsRoundRocks[rock.X] = append(columnsRoundRocks[rock.X], rock)
	sort.Slice(columnsRoundRocks[rock.X], func(i, j int) bool {
		return columnsRoundRocks[rock.X][i].Y < columnsRoundRocks[rock.X][j].Y
	})
}

func initRoundRocksIndexes(roundRocks []*Rock) (rows, cols map[int][]*Rock) {
	rows = make(map[int][]*Rock)
	cols = make(map[int][]*Rock)

	for r := range roundRocks {
		_, okRow := rows[roundRocks[r].Y]
		if !okRow {
			rows[roundRocks[r].Y] = []*Rock{roundRocks[r]}
		} else {
			rows[roundRocks[r].Y] = append(rows[roundRocks[r].Y], roundRocks[r])
		}

		_, okCol := cols[roundRocks[r].X]
		if !okCol {
			cols[roundRocks[r].X] = []*Rock{roundRocks[r]}
		} else {
			cols[roundRocks[r].X] = append(cols[roundRocks[r].X], roundRocks[r])
		}
	}

	return rows, cols
}

func rocks(mtx [][]byte) (squareRocks, roundRocks []*Rock) {
	squareRocks = []*Rock{}
	roundRocks = []*Rock{}

	for y := range mtx {
		for x := range mtx[y] {
			if mtx[y][x] == '#' {
				squareRocks = append(squareRocks, &Rock{Y: y, X: x})
			}
			if mtx[y][x] == 'O' {
				roundRocks = append(roundRocks, &Rock{Y: y, X: x})
			}
		}
	}

	return squareRocks, roundRocks
}

func buildLimits(squareRocks []*Rock) (rows, cols map[int][]*Rock) {
	rowsSquareRocks := make(map[int][]*Rock)
	columnsSquareRocks := make(map[int][]*Rock)

	for sqr := range squareRocks {
		_, okRow := rowsSquareRocks[squareRocks[sqr].Y]
		if !okRow {
			rowsSquareRocks[squareRocks[sqr].Y] = []*Rock{squareRocks[sqr]}
		} else {
			rowsSquareRocks[squareRocks[sqr].Y] = append(rowsSquareRocks[squareRocks[sqr].Y], squareRocks[sqr])
		}

		_, okColumn := columnsSquareRocks[squareRocks[sqr].X]
		if !okColumn {
			columnsSquareRocks[squareRocks[sqr].X] = []*Rock{squareRocks[sqr]}
		} else {
			columnsSquareRocks[squareRocks[sqr].X] = append(columnsSquareRocks[squareRocks[sqr].X], squareRocks[sqr])
		}
	}

	return rowsSquareRocks, columnsSquareRocks
}

func loadOnNorthBeamV2(lenMtx int, roundRocks []*Rock) int {
	out := 0

	for i := range roundRocks {
		out += lenMtx - roundRocks[i].Y
	}

	return out
}

const (
	RED  = "\033[91m"
	ENDC = "\033[0m"
)

func printFieldV2(mtx [][]byte, roundRocks []*Rock, rrr *Rock) {
	if rrr != nil {
		fmt.Println("ROCK:", rrr.Y, rrr.X)
	}

	ballsIndex := make(map[Rock]struct{})
	for b := range roundRocks {
		ballsIndex[*roundRocks[b]] = struct{}{}
	}

	for y := range mtx {
		row := make([]byte, 0, len(mtx[y]))
		for x := range mtx[y] {
			if mtx[y][x] == '#' {
				row = append(row, '#')
				continue
			}
			if _, ok := ballsIndex[Rock{Y: y, X: x}]; ok {
				if rrr != nil && rrr.Y == y && rrr.X == x {
					row = append(row, append(append([]byte(RED), 'O'), []byte(ENDC)...)...)
				} else {
					row = append(row, 'O')
				}

				continue
			}

			row = append(row, '.')
		}

		fmt.Println(string(row))
	}

	fmt.Println()
}
