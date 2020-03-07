package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		testCases, _ := strconv.Atoi(scanner.Text())
		for test := 0; test < testCases; test++ {
			scanner.Scan()
			firstInput := strings.Split(scanner.Text(), " ")
			numStudents, _ := strconv.Atoi(firstInput[0])
			sPerTeam, _ := strconv.Atoi(firstInput[1])
			scanner.Scan()
			secondInput := strings.Split(scanner.Text(), " ")
			skills := []int{}
			for i := 0; i < numStudents; i++ {
				j, _ := strconv.Atoi(secondInput[i])
				skills = append(skills, j)
			}
			sort.Sort(sort.Reverse(sort.IntSlice(skills)))
			optimalCase := skills[0] * skills[0]

			for i := 0; i < numStudents-(sPerTeam-1); i++ {
				skill := skills[i]
				skilDiference := 0
				for j := 0; j < (sPerTeam - 1); j++ {
					skilDiference += skill - skills[(i+1)+j]
				}
				if skilDiference < optimalCase {
					optimalCase = skilDiference
				}
			}
			fmt.Printf("Case #%v: %v\n", test+1, optimalCase)
		}
	}
}
