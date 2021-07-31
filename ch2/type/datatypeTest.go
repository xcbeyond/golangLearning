package main

import (
	"fmt"
	"math"
)

func main() {

	var a = 10
	fmt.Println(a == 10) // true
	fmt.Println(a == 5)  // false
	fmt.Println(a != 10) // false
	fmt.Println(a != 5)  // true

	// fmt.Println(a == true)	// cannot use true (type untyped bool) as type int

	// var n bool
	// fmt.Println(int(n) * 2)		// cannot convert n (type bool) to type int

	fmt.Println(math.MaxFloat32) // float32类型的最大值

	// 复数
	var cmp complex128 = complex(1, 2)
	fmt.Println(real(cmp)) // 输出实部
	fmt.Println(imag(cmp)) // 输出虚部

	// 字符串
	var str = "兄弟们！\n赶紧开始学习Go语言啦。"
	fmt.Println(str)
	var str2 = `第一行
	第二行
	第三行
	`
	fmt.Println(str2)

	// 指针
	var i int = 20                   // 声明实际变量
	var ip *int                      // 声明指针变量
	ip = &i                          // 指针变量的存储地址
	fmt.Println("i变量的地址是: ", &i)     // 变量的存储地址
	fmt.Println("ip变量储存的指针地址: ", ip) // 指针变量的存储地址
	fmt.Println("*ip变量的值: ", *ip)    // 使用指针访问值

	// 数组
	var intArr [10]int // 定义一个int类型数组
	var idx int
	// 遍历数组赋值
	for idx = 0; idx < 10; idx++ {
		intArr[idx] = idx * 10
	}
	for idx = 0; idx < 10; idx++ {
		fmt.Printf("intArr[%d] = %d\n", idx, intArr[idx])
	}

	// 通道
	var c chan int
	if c == nil {
		fmt.Println("通道c是nil的, 不能使用，需要先创建通道。")
		c = make(chan int)
		fmt.Printf("c的数据类型是：%T\n", c)
	}

	fmt.Println(max(10, 7))

	x, y := swap("xcbeyond", "Niki")
	fmt.Println(x, y)

	// Map
	personMap := make(map[string]string)
	personMap["zhangsan"] = "张三"
	personMap["lisi"] = "李四"
	personMap["wanger"] = "王二"
	personMap["zhaowu"] = "赵五"
	for p := range personMap {
		fmt.Println(p, "的中文名是", personMap[p])
	}

	// 类型转换
	var i16 int16 = 10000
	i32 := int32(i16)
	fmt.Print(i32)

}

// 获取两数的最大值
func max(num1, num2 int) int {
	if num1 > num2 {
		return num1
	} else {
		return num2
	}
}

// 交换两个字符串
func swap(x, y string) (string, string) {
	return y, x
}
