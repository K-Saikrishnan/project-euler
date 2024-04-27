// 10001st Prime

package main

import (
	"fmt"

	"github.com/K-Saikrishnan/project-euler/helper"
)

func main() {
	var primeCount, num int

	for primeCount < 10001 {
		num++
		if helper.IsPrime(num) {
			primeCount++
		}
	}

	fmt.Println(num)
}
