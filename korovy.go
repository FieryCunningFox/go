package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	if n%10 == 1 {
		if n%100 != 11 {
			fmt.Print(n, " korova")
		} else {
			fmt.Print(n, " korov")
		}
	} else {
		if (n%10) > 1 && (n%10) < 5 {
			if (n%100) > 11 && (n%100) < 15 {
				fmt.Print(n, " korov")
			} else {
				fmt.Print(n, " korovy")
			}
		} else {
			fmt.Print(n, " korov")
		}
	}
}
