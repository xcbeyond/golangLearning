package main

import "fmt"

func main() {
	fmt.Println(multipleForTowNum(1, 1))
	fmt.Println(multipleForNums(1, 2, 3, 4, 5))
	fmt.Println(multipleForTowNums(1, 1, 1, 2, 3, 4, 5))

	// 内置函数
	println("")

	x, _ := swap(1, 5)
	fmt.Println(x)
}

func multipleForTowNum(num1 int, num2 int) int {
	result := num1 * num2
	return result
}

func multipleForNums(nums ...int) int {
	var result int = 1
	for _, num := range nums {
		result = result * num
	}
	return result
}

func multipleForTowNums(num1 int, num2 int, nums ...int) int {
	result := num1 + num2
	for _, num := range nums {
		result = result * num
	}
	return result
}

func swap(x, y int) (int, int) {
	return y, x
}
