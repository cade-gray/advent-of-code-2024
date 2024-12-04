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
	// Read the input file
	left, right, err := readInputFile("input.txt")
	if err != nil {
		fmt.Println("Error reading the file: ", err)
		return
	}
	// Sort the left and right arrays
	sort.Ints(left)
	sort.Ints(right)

	// Get difference of the left and right arrays, ensuring the difference is positive
	diffArr := []int{}
	for i := 0; i < len(left); i++ {
		diffArr = append(diffArr, abs(left[i]-right[i]))
	}

	// Sum up the diffArr for the final answer for part 1
	sum := 0
	for i := 0; i < len(diffArr); i++ {
		sum += diffArr[i]
	}
	fmt.Println("Sum of the diffArr and final answer for part 1: ", sum)

	// Collect similarity scores of left and right arrays.  Similarity score is the count of the same elements in the left and right arrays times the element itself.
	simScoresArr := []int{}
	for i := 0; i < len(left); i++ {
		count := 0
		for j := 0; j < len(right); j++ {
			if left[i] == right[j] {
				count++
			}
		}
		simScoresArr = append(simScoresArr, left[i]*count)
	}
	// Sum up similarity scores for the final answer for part 2
	simScoresArrSum := 0
	for i := 0; i < len(simScoresArr); i++ {
		simScoresArrSum += simScoresArr[i]
	}
	fmt.Println("Similarity score of left and right arrays and final answer for part 2: ", simScoresArrSum)
}

func readInputFile(filename string) ([]int, []int, error) {
	// Open the file using os package
	file, err := os.Open(filename)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()
	left := []int{}
	right := []int{}
	// Declare a scanner to read the file.
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Read the line
		line := scanner.Text()
		// Split the line by space
		values := strings.Fields(line)
		// Convert the strings to integers. ASCII to integer function part of strconv package
		num1, err := strconv.Atoi(values[0])
		if err != nil {
			return nil, nil, err
		}
		num2, err := strconv.Atoi(values[1])
		if err != nil {
			return nil, nil, err
		}
		// Append the integers to the left and right arrays
		left = append(left, num1)
		right = append(right, num2)

	}

	return left, right, nil
}
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
