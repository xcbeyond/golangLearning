package main

import "fmt"

func main() {
	var a int = 30

	if a < 20 {
		fmt.Printf("a 小于 20\n")
	} else {
		fmt.Printf("a 大于 20\n")
	}

	var b int = 100

	if a == 30 {
		if b == 100 {
			fmt.Printf("a = %d, b = %d\n", a, b)
		}
	}
}
