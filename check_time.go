package main

import "fmt"

func main() {
	var k, h, m int
	fmt.Scan(&k)
	h = k / (60 * 60)
	m = (k - 3600 * h) / 60
	fmt.Print("It is ", h, " hours ", m, " minutes.")
}