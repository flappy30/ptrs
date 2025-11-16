package main

import (
	"fmt"
	"math/rand"
)

var number = 1
var ptr *int

func main() {
	pointerNumber := &number
	enter(pointerNumber)

	if ptr != nil {
		fmt.Println(*ptr)
	} else {
		fmt.Println("Паник")
	}

}

func enter(a *int) {
	
	for {
		randomNum := rand.Intn(5)
		if randomNum == 1 {
			break
		}

		*a++
	}
	fmt.Println(*a)
}