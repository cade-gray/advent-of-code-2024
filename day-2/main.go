package main

import (
	"bufio"
	"day-2/utils"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	// Record the start time of the program
	start := time.Now()
	// Open the file using os package
	file, err := os.Open("input.txt")
	if err != nil {
		return
	}
	defer file.Close()
	// Declare a scanner to read the file.
	scanner := bufio.NewScanner(file)
	safeCount := 0
	SafeDampenerCount := 0
	for scanner.Scan() {
		line := scanner.Text()
		values := strings.Fields(line)
		intValues := make([]int, len(values))
		for i := 0; i < len(values); i++ {
			intValues[i], _ = strconv.Atoi(values[i])
		}
		safe := safetyCheck(intValues)
		safeDampener := safetyCheckDampener(intValues)
		if safe {
			safeCount++
		}
		if safe || safeDampener {
			SafeDampenerCount++
		}
	}
	fmt.Println("Total Safe Lines are", safeCount)
	fmt.Println("Total Safe Lines with Dampener are", SafeDampenerCount)

	// Calculate and print the elapsed time
	elapsed := time.Since(start)
	fmt.Printf("Execution time: %s\n", elapsed)
}

/*
Safety Check function checks if the slice meets the following criteria:
- The levels are either all increasing or all decreasing.
- Any two adjacent levels differ by at least one and at most three.

Parameters:
- values: A slice of integers representing the levels from the input file.

Returns:
- A boolean value indicating whether the slice meets the criteria.
*/
func safetyCheck(values []int) bool {
	dec := utils.IsDecreasing(values)
	inc := utils.IsIncreasing(values)
	max := utils.MaxDiff(values, 3)
	min := utils.MinDiff(values, 1)
	return (dec || inc) && max && min
}

/*
Safety Check Dampener function checks if the slice meets the following criteria:
- The levels are either all increasing or all decreasing.
- Any two adjacent levels differ by at least one and at most three.
- The slice is safe with one element removed.

Parameters:
- values: A slice of integers representing the levels from the input file.

Returns:
- A boolean value indicating whether the slice meets the criteria.
*/
func safetyCheckDampener(values []int) bool {
	// Loop through each element in the slice and remove it and check if the slice is safe with it removed.
	for i := 0; i < len(values); i++ {
		// Copy the slice
		valuesCpy := append([]int{}, values[:i]...)
		// Append the rest of the slice, excluding/cutting the current element
		valuesCpy = append(valuesCpy, values[i+1:]...)
		if safetyCheck(valuesCpy) {
			return true
		}
	}
	return false
}
