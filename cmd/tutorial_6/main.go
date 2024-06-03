package main

import "fmt"

type gasEngine struct {
	mpg       uint8
	gallons   uint8
	ownerInfo owner
}

type owner struct {
	name string
	age  uint8
}

func (e gasEngine) milesLeft() uint8 {
	return e.gallons * e.mpg
}

func canMakeIt(e gasEngine, miles uint8) {
	if miles <= e.milesLeft() {
		fmt.Println("You can make it there!")
	} else {
		fmt.Println("Need to fuel up First!")
	}
}

func main() {
	var myEngine gasEngine = gasEngine{25, 15, owner{"alex", 35}}
	// fmt.Println(myEngine.mpg, myEngine.gallons)
	var milesLeft = myEngine.milesLeft()
	fmt.Printf("Total miles left in tank: %v", milesLeft)
}
