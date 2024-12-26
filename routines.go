package main

//
//import (
//	"fmt"
//	"runtime"
//	"time"
//)
//
//func f(from string) {
//	for i := 0; i < 3; i++ {
//		fmt.Println(from, ":", i)
//	}
//}
//
//func main() {
//	f("direct")
//
//	go f("goroutine")
//
//	go func(msg string) {
//		fmt.Println(msg)
//	}("going")
//
//	time.Sleep(time.Second)
//	fmt.Println("done")
//
//	//	messages := make(chan string)
//
//	//go func() { messages <- "ping" }()
//
//	//msg := <-messages
//
//	//fmt.Println(msg)
//
//	msg1 := make(chan string, 2)
//	msg1 <- "Hello"
//	msg1 <- "Hello1"
//	//msg1 <- "Hello2"
//
//	fmt.Println(<-msg1)
//	fmt.Println(<-msg1)
//
//	buf := make([]byte, 1024)
//	n := runtime.Stack(buf, false)
//	fmt.Printf("Goroutine Stack Trace: %s\n", string(buf[:n]))
//}
//
////func printGoroutineID() {
////	fmt.Printf("Current Goroutine ID: %d\n", runtime.go)
////}
