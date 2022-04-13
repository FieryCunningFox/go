package main

import "fmt"

func main() {
	var x, y, cur, p float64
	var i int
	fmt.Scanln(&x)
	fmt.Scanln(&p)
	fmt.Scanln(&y)
	cur = x
	for ; cur < y;  {
		cur *= (1 + p / 100)
		i++;
	}
	fmt.Print(i)
}