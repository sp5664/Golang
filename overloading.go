package main

type MyStruct struct{}

func (m MyStruct) PrintValues(values ...interface{}) {

}

//type calc struct {
//}
//
//type Circle struct {
//	radius float64
//}
//
//func (c Circle) area() float64 {
//	return 3.14 * c.radius * c.radius
//}
//
//func (c calc) add(values ...int) int {
//	sum := 0
//	for _, v := range values {
//		sum += v
//	}
//	return sum
//}
//
//func main() {
//	c := Circle{radius: 12}
//	fmt.Println(c.area())
//	c1 := calc{}
//	fmt.Println(c1.add(1, 2))
//	fmt.Println(c1.add(1, 2, 3, 4))
//}
