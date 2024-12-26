package main

import (
	"fmt"
	"strconv"
)

var a int = 10

func main() {
	fmt.Println(a)
	var x int = 42
	var y float64 = float64(x) // Converting int to float64
	fmt.Println(y)

	var x1 string = strconv.Itoa(x)

	fmt.Println(x1)
}
