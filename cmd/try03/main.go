package main

import (
	"fmt"
	"sync"
)

func main() {
	var mywg sync.WaitGroup

	mywg.Add(1)
	go func() {
		fmt.Print("A ")
		mywg.Done()
	}()

	mywg.Add(1)
	go func() {
		fmt.Print("B ")
		mywg.Done()
	}()

	mywg.Wait()
	fmt.Print("C ")
}
