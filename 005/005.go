// Smallest Multiple

package main

import (
	"fmt"

	"github.com/K-Saikrishnan/project-euler/helper"
)

func main() {
	const num = 20

	numRange := make([]int, num)
	for i := range numRange {
		numRange[i] = i + 1
	}

	smallestMultiple := helper.Lcm(numRange...)

	fmt.Println(smallestMultiple)
}
