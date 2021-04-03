package main

import "fmt"

func main() {
	name := "xcbeyond"

	fmt.Print(name) // xcbeyond
	fmt.Println()

	fmt.Printf("%.5s", name)
	fmt.Println()

	fmt.Println(name)

	retStr := fmt.Sprint(name)
	fmt.Print(retStr) // xcbeyond
	fmt.Println()

	retStr2 := fmt.Sprintf("%.5s", name)
	fmt.Print(retStr2) // xcbey
	fmt.Println()

	retStr3 := fmt.Sprintln(name)
	fmt.Print(retStr3) // xcbeyond  换行

	// fmt.Fprint("xx", name)
	// fmt.Fprintf()
	// fmt.Fprintln()

}
