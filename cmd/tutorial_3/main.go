package main

import (
	"errors"
	"fmt"
)

func main() {
	var printValue string = "hello world!"
	printMe(printValue)

	var firstNum int = 100
	var secondNum int = 0
	result, remainder, err := intDivison(firstNum, secondNum)

	switch {
	case err != nil:
		fmt.Println(err.Error())
	case remainder == 0:
		fmt.Printf("the result of the integer devison is %v", result)
	default:
		fmt.Printf("the results of the integer devision is %v with remainder %v", result, remainder)
	}

	// if err != nil {
	// 	fmt.Println(err.Error())
	// } else if remainder == 0 {
	// 	fmt.Printf("the result of the integer devison is %v", result)
	// } else {
	// 	fmt.Printf("the results of the integer devision is %v with remainder %v", result, remainder)
	// }

}

func printMe(value string) {
	fmt.Println(value)
}

func intDivison(a int, b int) (int, int, error) {
	var err error

	if b == 0 {
		err = errors.New("cannot devide by zero")
		return a, 0, err
	}

	var result int = a / b
	var remainder int = a % b
	return result, remainder, err
}
