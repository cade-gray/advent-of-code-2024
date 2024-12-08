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
		safe, changeCd := safetyCheck(intValues)
		if safe {
			fmt.Println(values, "Are Safe and the change is", changeCd)
			safeCount++

		} else {
			fmt.Println(values, "Are Unsafe")
		}
	}
	fmt.Println("Total Safe Lines are", safeCount)
}

func safetyCheck(values []int) (bool, string) {
	unsafe := false
	prev := 0
	changeCd := ""
	for i := 0; i < len(values); i++ {
		curr := values[i]
		if i != 0 {
			if i == 1 {
				if curr < prev && safeChange(curr, prev) {
					changeCd = "dec"
				} else if curr > prev && safeChange(curr, prev) {
					changeCd = "inc"
				} else {
					unsafe = true
				}
			} else {
				if prev == curr {
					unsafe = true
				} else if changeCd == "dec" {
					if curr > prev || !safeChange(curr, prev) {
						unsafe = true
					}
				} else if changeCd == "inc" {
					if curr < prev || !safeChange(curr, prev) {
						unsafe = true
					}
				}
			}
		}
		prev = curr
	}
	return !unsafe, changeCd

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
