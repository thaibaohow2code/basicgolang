package main

import (
	"fmt"
	"time"
)

// for if/esle statements
func main() {
	n := time.Now().Unix()
	// if - else if - else
	if n%2 == 0 {
		fmt.Println("n % 2 == 0", n)
	} else if n%4 == 0 {
		fmt.Println("n % 4 == 0", n)
	} else if n%6 == 0 {
		fmt.Println("n % 6 == 0", n)
	} else {
		fmt.Println("n == ", n)
	}
}
