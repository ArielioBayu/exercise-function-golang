package main

import (
	"fmt"
)

func FindMin(nums ...int) int {
	min := nums[0]
	for _, result := range nums {
		if result < min {
			min = result
		}
	}
	return min // TODO: replace this
}

func FindMax(nums ...int) int {
	max := nums[0]
	for _, result := range nums {
		if result > max {
			max = result
		}
	}
	return max // TODO: replace this
}

func SumMinMax(nums ...int) int {
	min := FindMin(nums...)
	max := FindMax(nums...)

	return min + max // TODO: replace this
}

// gunakan untuk melakukan debug
func main() {
	fmt.Println(SumMinMax(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
	fmt.Println(SumMinMax(333, 456, 654, 123, 111, 1000, 1500, 2000, 3000, 1250, 1111))
}
