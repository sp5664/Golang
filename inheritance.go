package main

import "fmt"

type Animal1 struct {
	name string
}

type Dog1 struct {
	Animal1
}

type Speak1 interface {
	speak1()
}

func (a Animal1) speak1() {
	fmt.Println(a.name, "make sound")
}

func (d Dog1) speak1() {
	fmt.Println(d.name, "barks")
}

func main() {
	d := Dog1{Animal1{name: "Rex"}}
	d.speak1()
}
