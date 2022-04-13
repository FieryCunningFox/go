package main

import "fmt"

func main() {
	var n, k, res int
	k = 1
	fmt.Scan(&n)
	for i := n % 2; n > 0; i = n % 2 {
		res += k * i
		k *= 10
		n /= 2
	}
	fmt.Print(res)
}