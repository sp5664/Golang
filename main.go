package main

import "fmt"

//func main() {
//	//m := make(map[string]int)
//	//m["k1"] = 18
//	//m["k2"] = 19
//	//m["k3"] = 20
//	//fmt.Println(m)
//	//delete(m, "k1")
//	//_, prs := m["k3"]
//	//fmt.Println(m)
//	//fmt.Println("prs", prs)
//	//
//	//res := plus(10, 1)
//	//fmt.Println("plus", res)
//	//
//	//res1 := plusplus(1, 2, 3)
//	//fmt.Println("plusplus", res1)
//	//
//	//a, b := vals()
//	//fmt.Println("a:", a)
//	//fmt.Println("b", b)
//	//sum(1, 2)
//	//sum(1, 2, 3)
//
//	nextInt := intSeq()
//	fmt.Println(nextInt())
//	fmt.Println(nextInt())
//	fmt.Println(nextInt())
//
//	newInts := intSeq()
//	fmt.Println(newInts())
//
//}

func plus(a int, b int) int {
	return a + b
}

func plusplus(a, b, c int) int {
	return a + b + c
}

func vals() (int, int) {
	return 3, 7
}
func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0

	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
