// Multiples of 3 or 5

package main

import "fmt"

func sumDivisibleBy(divisor int, num int) int {
	num--
	quotient := num / divisor
	return divisor * quotient * (quotient + 1) / 2
}

func main() {
	const num = 1000
	sum := sumDivisibleBy(3, num) + sumDivisibleBy(5, num) - sumDivisibleBy(15, num)
	fmt.Println(sum)
}
