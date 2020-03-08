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
			optimalCase := uint64(skills[0]) * uint64(skills[0])

			for i := 0; i <= numStudents-sPerTeam; i++ {
				actualSkill := uint64(skills[i])
				teamSkill := 0
				for j := i; j < (i + sPerTeam); j++ {
					teamSkill += skills[j]
				}
				totalSkill := uint64(sPerTeam)*actualSkill - uint64(teamSkill)
				if totalSkill < optimalCase {
					optimalCase = totalSkill
				}
			}
			fmt.Printf("Case #%v: %v\n", test+1, optimalCase)
		}
	}
}
