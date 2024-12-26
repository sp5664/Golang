package main

import "fmt"

func main() {

	slice := []int{1, 2, 3, 4}
	printslice(slice)
	println(slice[0])

	m := make(map[string]int)
	m["k1"] = 1

	printmap(m)

	fmt.Println("map:", m)

	s1 := 1

	fmt.Println(s1)
}

func printmap(m map[string]int) {
	m["k2"] = 2
}

func printslice(slice []int) {
	//fmt.Println(slice[0])
	slice[0] = 100
}
