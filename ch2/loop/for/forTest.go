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

	// 不使用标记
	fmt.Println("---- break ----")
	for i := 1; i <= 3; i++ {
		fmt.Printf("i: %d\n", i)
		for i2 := 11; i2 <= 13; i2++ {
			fmt.Printf("i2: %d\n", i2)
			break
		}
	}

	// 使用标记
	fmt.Println("---- break label ----")
re:
	for i := 1; i <= 3; i++ {
		fmt.Printf("i: %d\n", i)
		for i2 := 11; i2 <= 13; i2++ {
			fmt.Printf("i2: %d\n", i2)
			break re
		}
	}

print:
	fmt.Printf("****\n")

}
