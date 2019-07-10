package main

import "fmt"

const s string = "constant"

func main() {
	fmt.Println(s)
	const n = 500000000
	fmt.Println("constant n = ", n)
	//n = 1
}
