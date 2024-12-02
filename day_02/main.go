package main

import (
	"fmt"
	"log"
	"strconv"

	_ "embed"
	"strings"
)

//go:embed input.txt
var input string

func part01(lines []string) int {

	count := 0
	for _, line := range lines {

		levels := strings.Split(line, " ")

		if isSafe(levels) {
			count++
		}

	}

	return count

}

func part02(lines []string) int {

	count := 0
	for _, line := range lines {

		levels := strings.Split(line, " ")

		if isSafe(levels) {
			fmt.Println("debug")
			count++
		} else {
			//remove one item

			for i := 0; i < len(levels); i++ {
				newArr := []string{}
				for j := 0; j < len(levels); j++ {
					if i != j {
						newArr = append(newArr, levels[j])
					}
				}

				if isSafe(newArr) {
					count++
					break
				}

			}
		}

	}

	return count
}

func isSafe(levels []string) bool {
	var safe bool
	ok := decreasing(levels)
	if !ok {
		safe = increasing(levels)
	} else {
		safe = true
	}

	return safe

}

func decreasing(levels []string) bool {

	for i := 1; i < len(levels); i++ {
		num, err := strconv.Atoi(levels[i-1])
		if err != nil {
			log.Fatal(err)
		}

		num2, err := strconv.Atoi(levels[i])
		if err != nil {
			log.Fatal(err)
		}
		sub := num - num2
		if sub < 1 {
			return false
		}

		if sub > 0 && sub < 4 { //gradually decreasing
			continue
		} else {
			return false
		}
	}

	return true
}

// these increasing and decreasing could handled in an one function
// wanted to solve it faster and avoided complexity
func increasing(levels []string) bool {
	for i := 1; i < len(levels); i++ {
		num, err := strconv.Atoi(levels[i-1])
		if err != nil {
			log.Fatal(err)
		}

		num2, err := strconv.Atoi(levels[i])
		if err != nil {
			log.Fatal(err)
		}
		sub := num2 - num
		if sub < 1 {
			return false
		}

		if sub > 0 && sub < 4 { //gradually decreasing
			continue
		} else {
			return false
		}
	}
	return true
}

func main() {

	fmt.Println(input)

	var lines []string
	lines = append(lines, strings.Split(input, "\n")...)
	fmt.Println("lines", lines)
	// result := part01(lines)
	result := part02(lines)
	fmt.Println(result)
}
