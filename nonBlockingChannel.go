package main

//
//import (
//	"fmt"
//)
//
//func main() {
//	message := make(chan string)
//	signal := make(chan bool)
//
//	select {
//	case msg := <-message:
//		fmt.Println("received message", msg)
//	default:
//		fmt.Println("no message received")
//	}
//
//	msg := "Hi"
//
//	select {
//	case message <- msg:
//		fmt.Println("Send message:", msg)
//	default:
//		fmt.Println("no message sent")
//	}
//
//	select {
//	case msg := <-message:
//		fmt.Println("received message:", msg)
//	case sig := <-signal:
//		fmt.Println("received signal:", sig)
//	default:
//		fmt.Println("no activity")
//	}
//
//}
