package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type (
	Result struct {
		num1 string
		num2 string
	}
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		testCases, _ := strconv.Atoi(scanner.Text())
		for test := 0; test < testCases; test++ {
			scanner.Scan()
			input := scanner.Text()
			numbers := strings.Split(input, "")
			leftZero := true
			result := Result{}
			for i := 0; i < len(numbers); i++ {
				number := numbers[i]
				if number == "4" {
					result.num1 += "3"
					result.num2 += "1"
					leftZero = false
					continue
				}
				if !leftZero {
					result.num2 += "0"
				}
				result.num1 += number
			}
			fmt.Printf("Case #%v: %v %v\n", (test + 1), result.num1, result.num2)
		}
	}
}
