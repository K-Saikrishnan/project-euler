package helper

func IsPalindrome(n int) bool {
	rev := 0
	orig := n

	for n > 0 {
		rev = rev*10 + n%10
		n /= 10
	}

	return orig == rev
}

func Gcd(integers ...int) int {
	gcd := integers[0]
	for _, num := range integers[1:] {
		for num != 0 {
			gcd, num = num, gcd%num
		}
	}
	return gcd
}

func Lcm(integers ...int) int {
	lcm := integers[0]
	for _, num := range integers[1:] {
		lcm = (lcm * num) / Gcd(lcm, num)
	}
	return lcm
}
