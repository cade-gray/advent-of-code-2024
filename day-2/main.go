package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Open the file using os package
	file, err := os.Open("input.txt")
	if err != nil {
		return
	}
	defer file.Close()
	// Declare a scanner to read the file.
	scanner := bufio.NewScanner(file)
	safeCount := 0
	for scanner.Scan() {
		line := scanner.Text()
		values := strings.Fields(line)
		intValues := make([]int, len(values))
		for i := 0; i < len(values); i++ {
			intValues[i], _ = strconv.Atoi(values[i])
		}
		safe := safetyCheck(intValues)
		safeDampener := safetyCheckDampener(intValues)
		if safe || safeDampener {
			safeCount++
		} else {
			fmt.Println(values, "Are Unsafe")
		}
	}
	fmt.Println("Total Safe Lines are", safeCount)
}

func safetyCheck(values []int) bool {
	dec := isDecreasing(values)
	inc := isIncreasing(values)
	max := MaxDiff(values, 3)
	min := minDiff(values, 1)
	return (dec || inc) && max && min
}

func safetyCheckDampener(values []int) bool {
	for i := 0; i < len(values); i++ {
		valuesCpy := append([]int{}, values[:i]...)
		valuesCpy = append(valuesCpy, values[i+1:]...)
		if safetyCheck(valuesCpy) {
			return true
		}
	}
	return false
}

func isIncreasing(values []int) bool {
	increasing := true
	for i := 0; i < len(values); i++ {
		if i != 0 {
			if values[i] < values[i-1] {
				increasing = false
			}
		}
	}
	return increasing
}

func isDecreasing(values []int) bool {
	decreasing := true
	for i := 0; i < len(values); i++ {
		if i != 0 {
			if values[i] > values[i-1] {
				decreasing = false
			}
		}
	}
	return decreasing
}

func minDiff(input []int, min int) bool {
	for i := 1; i < len(input); i++ {
		if absValue(input[i]-input[i-1]) < min {
			return false
		}
	}
	return true
}

func MaxDiff(input []int, max int) bool {
	for i := 1; i < len(input); i++ {
		if absValue(input[i]-input[i-1]) > max {
			return false
		}
	}
	return true
}

func absValue(value int) int {
	if value < 0 {
		value = value * -1
	}
	return value
}

func safeChange(curr, prev int) bool {
	safelyChanged := true
	change := prev - curr
	if change < 0 {
		change = change * -1
	}
	if change >= 1 && change <= 3 {
		safelyChanged = true
	} else {
		safelyChanged = false
	}
	return safelyChanged
}
