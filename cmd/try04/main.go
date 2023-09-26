package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// COMPLETE THIS FUNCTION
func processInput(inputInt int, inputArray []string) {
	var positive []int
	var negative []int
	var resultArray []string

	for _, e := range inputArray {
		v, _ := strconv.Atoi(e)
		if v < 0 {
			negative = append(negative, v)
		} else {
			positive = append(positive, v)
		}
	}

	for _, e := range positive {
		resultArray = append(resultArray, strconv.Itoa(e))
	}

	for _, e := range negative {
		resultArray = append(resultArray, strconv.Itoa(e))
	}

	fmt.Println(resultArray)

	for _, e := range resultArray {
		fmt.Printf("%s ", e)
	}
}

// 6
// 20 10 80 -2 -4 1
// expected: 20 10 80 1 -2 -4

// Do not change the function name. You do not need to edit any code below this line.
func main() {
	reader := bufio.NewReader(os.Stdin)
	input1, _ := reader.ReadString('\n')
	inputInt, _ := strconv.Atoi(strings.TrimSpace(input1))
	input2, _ := reader.ReadString('\n')
	inputArray := strings.Split(input2, " ")
	processInput(inputInt, inputArray)
}
