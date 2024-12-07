package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Open the file using os package
	file, err := os.Open("sample.txt")
	if err != nil {
		return
	}
	defer file.Close()
	// Declare a scanner to read the file.
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		values := strings.Fields(line)
		for i := 0; i < len(values); i++ {
			fmt.Print(values[i])
		}
		fmt.Println()
	}
}

func safetyCheck(curr int, prev int) {

}
