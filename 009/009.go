// Special Pythagorean Triplet

package main

import "fmt"

func main() {
	const sum = 1000

	for a := 1; a < sum/3; a++ {
		for b := a + 1; b < sum/2; b++ {
			c := sum - a - b
			if a*a+b*b == c*c {
				fmt.Println(a * b * c)
				return
			}
		}
	}

}
