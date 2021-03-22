package main

import "fmt"

// 定义int类型的变量为NewInt
type NewInt int

// 定义int类型的别名IntAlias
type IntAlias = int

func main() {
	var a NewInt
	fmt.Printf("a type:%T\n", a)

	var a2 IntAlias
	fmt.Printf("a2 type:%T\n", a2)
}
