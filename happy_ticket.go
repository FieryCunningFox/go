package main

import "fmt"

// проверить является ли билет счастливым

func sum(a int) int {
	var res = (a / 100) + (a % 100) / 10 + a % 10
	return res
}

func main() {
	var ticket int
	fmt.Scan(&ticket)
	if sum(ticket/1000) == sum(ticket%1000) {
		fmt.Print("YES")
	} else {
		fmt.Print("NO")
	}
}