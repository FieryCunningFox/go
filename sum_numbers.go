package main

import "fmt"

// посчитать сумму от a до b

func main() {
	var a, b int
	var res = 0
	fmt.Scan(&a, &b)
	for i := a; i <= b; i++ {
		res += i
	}
	fmt.Print(res)
}
