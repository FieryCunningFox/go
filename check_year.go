package main

import "fmt"

// определить является ли год високосным

func main() {
	var year int
	fmt.Scan(&year)
	if (year%400 == 0) || (year%4 == 0 && year%100 != 0) {
		fmt.Print("YES")
	} else {
		fmt.Print("NO")
	}
}