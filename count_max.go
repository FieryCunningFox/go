package main

import "fmt"

// найти количество максимальных элементов последовательности

func main() {
	var n int
	var m, count = 0, 0
	for fmt.Scanln(&n); n != 0; fmt.Scanln(&n) {
		if n > m {
			count = 1
			m = n
		} else {
			if n == m {
				count++
			}
		}
	}
	fmt.Print(count)
}