package main

import (
	"fmt"
	"os"
	"sync"
)

var (
	// 一个演示用的映射
	valueByKey = make(map[string]int)
	// 保证使用映射时的并发安全的互斥锁
	valueByKeyGuard sync.Mutex
)

func main() {
	fmt.Println("defer begin")

	defer fmt.Println(1)

	defer fmt.Println(2)

	defer fmt.Println(3)

	fmt.Println("defer end")

	// fmt.Println(fileSize("file"))
}

// 根据键读取值
func readValue(key string) int {
	// 对共享资源加锁
	valueByKeyGuard.Lock()
	// 取值
	v := valueByKey[key]
	// 对共享资源解锁
	valueByKeyGuard.Unlock()
	// 返回值
	return v
}

// 使用延迟并发解锁
func readValue2(key string) int {
	valueByKeyGuard.Lock()
	// defer后面的语句不会马上调用, 而是延迟到函数结束时调用
	defer valueByKeyGuard.Unlock()
	return valueByKey[key]
}

// 根据文件名查询其大小
func fileSize(filename string) int64 {
	// 根据文件名打开文件, 返回文件句柄和错误
	f, err := os.Open(filename)
	// 如果打开时发生错误, 返回文件大小为0
	if err != nil {
		return 0
	}
	// 取文件状态信息
	info, err := f.Stat()

	// 如果获取信息时发生错误, 关闭文件并返回文件大小为0
	if err != nil {
		f.Close()
		return 0
	}
	// 取文件大小
	size := info.Size()
	// 关闭文件
	f.Close()

	// 返回文件大小
	return size
}

func fileSize2(filename string) int64 {
	f, err := os.Open(filename)
	// 延迟调用Close, 此时Close不会被调用
	defer f.Close()
	if err != nil {
		return 0
	}

	info, err := f.Stat()
	if err != nil {
		// defer机制触发, 调用Close关闭文件
		return 0
	}
	size := info.Size()
	// defer机制触发, 调用Close关闭文件
	return size
}
