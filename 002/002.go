// Even Fibonacci numbers

package main

import "fmt"

func main() {
	f0, f1 := 1, 1
	sum := 0

	for f1 < 4_000_000 {
		sum += f0 + f1
		f0, f1 = f0+2*f1, 2*f0+3*f1
	}

	fmt.Println(sum)
}
