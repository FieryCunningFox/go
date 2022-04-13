package main

import "fmt"

// дано отклонение градусов часовой стрелки, определить время

func main() {
	var d, h, m int
	fmt.Scan(&d)
	h = d / 30
	m = (d % 30) * 2
	fmt.Print("It is ", h, " hours ", m, " minutes.")
}
