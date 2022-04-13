package main

import "fmt"

// проверить все ли цифры числа различны

func main() {
	var n, a, b, c int
	fmt.Scan(&n)
	a = n % 10
	b = (n % 100) / 10
	c = n / 100
	if a == b || b == c || a == c {
		fmt.Print("NO")
	} else {
		fmt.Print("YES")
	}
}
