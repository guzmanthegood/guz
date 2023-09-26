package main

import (
	"fmt"
)

func badCall() {
	panic("bad end")
}

func test() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Printf("panic! %s\\n", e)
		}
	}()
	fmt.Println("Start")
	badCall()
	fmt.Println("End")
}

func main() {
	fmt.Println("Calling test")
	test()
	fmt.Println("Test completed")
}
