package main

import "fmt"

func main() {
	// intArr := [...]int16{1, 2, 3}
	// fmt.Println(intArr[0:2])
	// fmt.Println(&intArr[0])

	var intSlice []int = []int{4, 5, 6}
	fmt.Println(len(intSlice))
	intSlice = append(intSlice, 465)
	fmt.Println(len(intSlice))

	var intSlice2 []int = []int{8, 9}
	intSlice = append(intSlice, intSlice2...)
	fmt.Println(len(intSlice))

	var intSlice3 []int = make([]int, 3, 5)
	intSlice3 = append(intSlice3, intSlice2...)
	fmt.Println(intSlice3)

	// myMap := make(map[string]int)

	var myMap2 = map[string]uint8{"Amir": 11, "Sina": 42}
	// delete(myMap2, "Amir")

	favNumber, ok := myMap2["Amir"]

	if ok {
		fmt.Printf("The Favorite Number is %v \n", favNumber)
	} else {
		fmt.Printf("Invalid Name \n")
	}

	for Name := range myMap2 {
		fmt.Printf(" name: %v\n", Name)
	}
}
