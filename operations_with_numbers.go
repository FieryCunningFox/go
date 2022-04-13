package main

import "fmt"

func main() {
	var n int
	for fmt.Scanln(&n); n <= 100; fmt.Scanln(&n) {
		if n >= 10 {
			fmt.Println(n)
		}
	}
}