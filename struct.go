package main

//
//import "fmt"
//
//type person struct {
//	name string
//	age  int
//}
//
//func newPerson(name string) *person {
//	p := person{name: name}
//	p.age = 42
//	return &p
//}
//
//type react struct {
//	width, height int
//}
//
//func (r *react) area() int {
//	return r.width * r.height
//}
//
//func (r *react) perim() int {
//	return 2*r.width + 2*r.height
//}
//
//func main() {
//	//fmt.Println(person{"Bob", 20})
//	//
//	//fmt.Println(person{name: "Alice", age: 30})
//	//
//	//fmt.Println(person{name: "fred"})
//	//
//	//fmt.Println(&person{name: "Ann", age: 40})
//	//
//	//fmt.Println(newPerson("Jon"))
//	//
//	//s := person{name: "sean", age: 50}
//	//fmt.Println(s.name)
//	//
//	//sp := s
//	//
//	//fmt.Println(sp.age)
//	r := react{height: 5, width: 10}
//
//	fmt.Println("area: ", r.area())
//	fmt.Println("perim: ", r.perim())
//
//	rp := &r
//	*rp = react{height: 1, width: 1}
//	fmt.Println("area: ", rp.area())
//	fmt.Println("perim; ", rp.perim())
//
//	fmt.Println("area: ", r.area())
//	fmt.Println("perim: ", r.perim())
//}
