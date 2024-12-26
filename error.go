package main

//
//import (
//	"errors"
//	"fmt"
//)
//
//var ErrOutOfTea = fmt.Errorf("No more tea available")
//var ErrPower = fmt.Errorf("can't boil water")
//
//func makeTea(arg int) error {
//	if arg == 2 {
//		return ErrOutOfTea
//	} else if arg == 4 {
//		return fmt.Errorf("making tea : %w", ErrPower)
//	}
//	return nil
//}
//
//func main() {
//	for i := 1; i < 6; i++ {
//		if err := makeTea(i); err != nil {
//
//			if errors.Is(err, ErrOutOfTea) {
//				fmt.Println("We should buy new tea!")
//			} else if errors.Is(err, ErrPower) {
//				fmt.Println("Now it is dark")
//			} else {
//				fmt.Printf("Unknown error: %s\n", err)
//			}
//			continue
//		}
//		fmt.Println("Tea is ready!")
//	}
//}
