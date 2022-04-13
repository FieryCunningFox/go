package main

import "fmt"

func main() {
	var n, cur int
	fmt.Scan(&n)
	for ; n >= 10; {
		for ; n > 0; {
			cur += n % 10
			n = n / 10
		}
		n = cur
		cur = 0
	}
	fmt.Print(n)
}