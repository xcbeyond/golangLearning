package main

import "fmt"

func main() {
	sum := 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println(sum)

	arr := []string{"A", "B", "C", "D"}
	for k, v := range arr {
		fmt.Printf("arr[%d]=%s\n", k, v)
	}

	var a int = 10
	for a < 20 {
		fmt.Printf("a=%d\n", a)
		a++

		if a > 15 {
			continue
		}
		fmt.Printf("****\n")
	}

	var b int = 10
	for b < 20 {
		fmt.Printf("b=%d\n", b)
		b++

		if b > 15 {
			goto print
		}

	}

print:
	fmt.Printf("****\n")

}
