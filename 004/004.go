// Largest Palindrome Product

package main

import (
	"fmt"
	"math"

	"github.com/K-Saikrishnan/project-euler/helper"
)

func main() {
	const digits = 3

	minNum := int(math.Pow10(digits - 1))
	maxNum := int(math.Pow10(digits))
	maxProduct := 0

	var j, dj int

	for i := maxNum; i >= minNum; i-- {
		if i%11 == 0 {
			j = maxNum
			dj = 1
		} else {
			j = maxNum - (maxNum % 11)
			dj = 11
		}

		for j = maxNum; j >= i; j -= dj {
			product := i * j
			if product <= maxProduct {
				break
			}
			if helper.IsPalindrome(product) {
				maxProduct = max(maxProduct, product)
			}
		}
	}

	fmt.Println(maxProduct)
}
