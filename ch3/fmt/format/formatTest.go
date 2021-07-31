package main

import (
	"fmt"
	"math"
)

type Person struct {
	Name string
}

func main() {
	// 普通占位符
	fmt.Println("普通占位符：")
	var p = new(Person)
	p.Name = "xcbeyond"
	fmt.Printf("%v\n", p.Name)
	fmt.Printf("%+v\n", p)
	fmt.Printf("%#v\n", p)
	fmt.Printf("%T\n", p.Name)
	fmt.Printf("%%\n")

	// 布尔占位符
	fmt.Println("布尔占位符：")
	fmt.Printf("%t\n", 1 == 1)

	// 复数占位符
	fmt.Println("复数占位符：")
	fmt.Printf("%.3e\n", math.Pi)
	fmt.Printf("%E\n", math.Pi)
	fmt.Printf("%.3f\n", math.Pi)
	fmt.Printf("%.3g\n", math.Pi)
	fmt.Printf("%.3G\n", 10.20+5i)

	// 字符串占位符
	fmt.Println("字符串占位符：")
	fmt.Printf("%s\n", "xcbeyond")
	fmt.Printf("%10s\n", "xcbeyond")
	fmt.Printf("%-10s\n", "xcbeyond")
	fmt.Printf("%.5s\n", "xcbeyond")
	fmt.Printf("%5.10s\n", "xcbeyond")
	fmt.Printf("%-5.10s\n", "xcbeyond")
	fmt.Printf("%5.3s\n", "xcbeyond")
	fmt.Printf("%010s\n", "xcbeyond")
	fmt.Printf("%q\n", "xcbeyond")
	fmt.Printf("%x\n", "xcbeyond")
	fmt.Printf("%X\n", "xcbeyond")

	// 指针占位符
	fmt.Println("指针占位符：")
	fmt.Printf("%p\n", &p)
	fmt.Printf("%#p\n", &p)
}
