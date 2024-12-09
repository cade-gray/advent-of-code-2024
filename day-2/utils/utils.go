package utils

func IsIncreasing(values []int) bool {
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

func IsDecreasing(values []int) bool {
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

func MinDiff(input []int, min int) bool {
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
