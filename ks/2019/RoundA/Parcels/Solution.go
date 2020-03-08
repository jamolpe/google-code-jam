package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type (
	squareInfo struct {
		bestMahattanD int
		x             int
		y             int
	}
)

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func findRecursive(x int, y int, mapDimension [][]int) (int, int) {
	allNodesChecked := true
	postFound := false
	maxX := x
	maxY := y
	minX := x
	minY := y
	if len(mapDimension) > x+1 {
		allNodesChecked = false
		maxX = x + 1
	}
	if len(mapDimension[x]) > y+1 {
		allNodesChecked = false
		maxY = y + 1
	}
	if x-1 >= 0 {
		allNodesChecked = false
		minX = x - 1
	}
	if y-1 >= 0 {
		allNodesChecked = false
		minY = y - 1
	}
	if allNodesChecked == false {
		for i := minX; i <= maxX; i++ {
			for j := minY; j <= maxY; j++ {
				if mapDimension[i][j] == 1 {
					postFound = true
					x = i
					y = j
				}
			}
		}
	} else {
		return x, y
	}

	if postFound == true {
		return x, y
	}
	return findRecursive(minX, minY, mapDimension)

}
func findEsierWay(x int, y int, mapDimension [][]int) squareInfo {
	var square squareInfo
	square.x = x
	square.y = y
	postInfo := mapDimension[x][y]

	if postInfo == 1 {
		square.bestMahattanD = 0
		return square
	}
	closerX, closerY := findRecursive(x, y, mapDimension)
	manHattan := Abs((x - closerX)) + Abs((y - closerY))
	square.bestMahattanD = manHattan
	return square
}

func findWorstManhattan(R int, C int, mapDimension [][]int) []squareInfo {
	squares := make([]squareInfo, 0)
	for x := 0; x < R; x++ {
		for y := 0; y < C; y++ {
			square := findEsierWay(x, y, mapDimension)
			squares = append(squares, square)
		}
	}
	sort.Slice(squares, func(i, j int) bool {
		return squares[i].bestMahattanD > squares[j].bestMahattanD
	})
	return squares
}

func updateWorstManhattans(R int, C int, mapDimension [][]int, squares []squareInfo) int {
	squaresToCheck := make([]squareInfo, 0)
	bestManhattan := squares[0]
	for _, square := range squares {
		if square.bestMahattanD >= 1 {
			squaresToCheck = append(squaresToCheck, square)
		}
	}
	for _, square := range squaresToCheck {
		modifiedMap := make([][]int, len(mapDimension))
		for i := range mapDimension {
			modifiedMap[i] = make([]int, len(mapDimension[i]))
			copy(modifiedMap[i], mapDimension[i])
		}
		modifiedMap[square.x][square.y] = 1
		worstManhattan := findWorstManhattan(R, C, modifiedMap)
		if bestManhattan.bestMahattanD > worstManhattan[0].bestMahattanD {
			bestManhattan = worstManhattan[0]
		}
	}
	return bestManhattan.bestMahattanD
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		testCases, _ := strconv.Atoi(scanner.Text())
		for test := 0; test < testCases; test++ {
			scanner.Scan()
			firstInput := strings.Split(scanner.Text(), " ")
			R, _ := strconv.Atoi(firstInput[0])
			C, _ := strconv.Atoi(firstInput[1])
			mapDimension := make([][]int, R)
			for i := 0; i < R; i++ {
				mapDimension[i] = make([]int, C)
				scanner.Scan()
				colums := strings.Split(scanner.Text(), "")
				for j := 0; j < C; j++ {
					mapDimension[i][j], _ = strconv.Atoi(colums[j])
				}
			}
			worstManhattan := findWorstManhattan(R, C, mapDimension)
			result := updateWorstManhattans(R, C, mapDimension, worstManhattan)
			fmt.Printf("Case #%v: %v\n", test+1, result)
		}
	}

}
