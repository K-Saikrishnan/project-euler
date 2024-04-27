// Summation of Primes
package main

import (
	"fmt"

	"github.com/K-Saikrishnan/project-euler/helper"
)

func main() {
	const num = 2_000_000

	sum := 0
	for i := 2; i < num; i++ {
		if helper.IsPrime(i) {
			sum += i
		}
	}

	fmt.Println(sum)
}
