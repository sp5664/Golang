package main

import (
	"fmt"
	"unicode/utf8"
)

var i12 int

func main() {
	var i1 = 2
	fmt.Println(i1)

	i := 1
	fmt.Println("initial:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	zeroptr(&i)

	fmt.Println("zeroptr:", i)

	fmt.Println(&i)

	const s = "Hello world"

	fmt.Println("length", len(s))

	for i := 0; i < len(s); i++ {
		fmt.Printf("%x", s[i])
	}

	fmt.Println("Run count:", utf8.RuneCountInString(s))

	for idx, runValue := range s {
		fmt.Printf("%#U starts at %d\n", runValue, idx)
	}
}

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}
