package main

import (
	"fmt"
)

func main() {
	siteViews := map[string]int{
		"x1": 1,
		"x2": 2,
		"x3": 3,
		"x4": 4,
		"x5": 5,
		"x6": 6,
	}

	for k, i := range siteViews {
		fmt.Println(k, i)
	}
}
