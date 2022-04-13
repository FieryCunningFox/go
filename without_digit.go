package main

import "fmt"

func main() {
	var n, digit, res, k int
	k = 1
	fmt.Scan(&n)
	fmt.Scan(&digit)
	for i := n % 10; n > 0; i = n % 10 {
		if i != digit {
			res += i * k
			k *= 10
		}
		n /= 10
	}
	fmt.Print(res)
}