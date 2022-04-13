package main

import "fmt"

func main() {
	var a, cur, x, flag int
	x = 1
	cur = 1
	fmt.Scan(&a)
	for i := 3; cur <= a; i++ {
		b := x
		x = cur
		cur += b
		if a == cur {
			fmt.Print(i)
			flag = 1
		}
	}
	if flag == 0 {
		fmt.Print(-1)
	}
}
