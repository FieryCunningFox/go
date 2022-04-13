package main

import "fmt"

// найти сумму двухзачных чисел, кратных 8

func main() {
	var n, cur int
	var sum = 0
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&cur)
		if cur%8 == 0 && cur/100 == 0 && cur/10 > 0 {
			sum += cur
		}
	}
	fmt.Print(sum)
}
