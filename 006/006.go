// Sum Square Difference

package main

import "fmt"

func main() {
	const num = 100

	sumOfSquares := num * (num + 1) * (2*num + 1) / 6
	squareOfSum := num * num * (num + 1) * (num + 1) / 4

	fmt.Println(squareOfSum - sumOfSquares)
}
