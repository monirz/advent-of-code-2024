package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Utility function to read input file
func readInput(fileName string) ([][]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines [][]string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineSplitted := strings.Split(scanner.Text(), " ")
		lines = append(lines, lineSplitted)

	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	var leftArr []int
	var rightArr []int
	for _, v := range lines {
		leftNum, _ := strconv.Atoi(v[0])
		leftArr = append(leftArr, leftNum)
		rightNum, _ := strconv.Atoi(v[3])
		rightArr = append(rightArr, rightNum)
	}

	//==== part two
	leftNumsCount := 0
	leftNums := []int{}
	for _, v := range leftArr {
		for _, j := range rightArr {
			if v == j {
				leftNumsCount++
			}
		}

		leftNum := v * leftNumsCount
		leftNums = append(leftNums, leftNum)
		leftNumsCount = 0
	}

	sum := 0
	for _, v := range leftNums {
		sum += v
	}

	fmt.Println("====", sum)

	distance := []int{}
	dist := 0
	for i := 0; i < len(lines); i++ {

		leftSmallest := 1
		rightSmallest := 1
		pos := 0
		rightPos := 0
		for k, v := range leftArr {

			if k == 0 {
				leftSmallest = v
			} else if v < leftSmallest {
				leftSmallest = v
				pos = k
			}

		}

		for k, v := range rightArr {
			if k == 0 {
				rightSmallest = v
			} else if v < rightSmallest {
				rightSmallest = v
				rightPos = k
			}

		}

		if leftArr[pos] > rightArr[rightPos] {
			dist := leftArr[pos] - rightArr[rightPos]
			distance = append(distance, dist)

		} else {
			dist := rightArr[rightPos] - leftArr[pos]

			distance = append(distance, dist)
		}

		leftArr = append(leftArr[:pos], leftArr[pos+1:]...)

		rightArr = append(rightArr[:rightPos], rightArr[rightPos+1:]...)
		// fmt.Println(n, "debug left after remove --")

	}

	fmt.Println("dist ", distance)

	for _, v := range distance {
		dist += v
	}
	fmt.Println("==================================================", dist)

	return lines, nil

}

// Converts a slice of strings to a slice of integers
func toIntSlice(input []string) ([]int, error) {
	var result []int

	for _, line := range input {
		splitted := strings.Split(line, " ")
		value, err := strconv.Atoi(splitted[0])
		if err != nil {
			return nil, err
		}

		result = append(result, value)

	}
	return result, nil
}

// Main function for Part 1
func solvePart1(input [][]string) int {
	// Implement the logic for part 1 here

	fmt.Println(input)
	return 0
}

// Main function for Part 2
func solvePart2(input []string) int {
	// Implement the logic for part 2 here
	return 0
}

func main() {
	// Adjust the file name as needed for the day's challenge
	inputFile := "input.txt"

	// Read input
	input, err := readInput(inputFile)
	if err != nil {
		fmt.Printf("Error reading input file: %v\n", err)
		return
	}

	// Solve Part 1
	solvePart1(input)
	// fmt.Printf("Part 1: %d\n", result1)

	// // Solve Part 2
	// result2 := solvePart2(input)
	// fmt.Printf("Part 2: %d\n", result2)
}
