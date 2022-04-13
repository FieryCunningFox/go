package main

import "fmt"

func main() {
	var a, k, n int
    var m = 10000000000
	fmt.Scan(&a)
    for i := 0; i < a; i++ {
        fmt.Scan(&k)
        if k < m {
            m = k
            n = 1
        } else {
            if k == m {
                n++
            }
        }
    }
    fmt.Print(n)
}

