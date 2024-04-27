// Largest Prime Factor

package main

import "fmt"

func main() {
	num := 600_851_475_143

	factor := 3

	for factor*factor <= num {
		if num%factor == 0 {
			num /= factor
		} else {
			factor += 2
		}
	}

	fmt.Println(num)
}
