package main

import "fmt"

func main() {
	// 声明一个没有初始化的int类型变量，默认为零值0
	var a int
	fmt.Println(a)

	// 声明一个没有初始化的bool类型变量，默认为零值false
	var b bool
	fmt.Println(b)

	var c = true
	fmt.Println(c)

	// 短变量声明
	d := "xcbeyond"
	fmt.Println(d)

	// 批量声明
	var (
		e int
		f string
		g float64
	)
	fmt.Println(e,f,g)


	// 匿名变量
	firstName, _, _ := getName()
	_, lastName, _ := getName()
	_, _, nickName := getName()
	fmt.Println(firstName, lastName, nickName)

	sum := sum(1,3)
	fmt.Println(sum)
}

func getName() (firstName, lastName, nickName string) { 
    return "X", "C", "xcbeyond" 
}

// 两数求和
func sum(a, b int) int{
	num := a + b
	return num
}