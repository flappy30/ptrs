package main

import (
	"fmt"
	"math/rand"
)

var number = 1
var ptr *int
var ptrMessage *messageStruct

type messageStruct struct {
	firstWord string
	secondWord string
}

func main() {
	pointerNumber := &number
	enter(pointerNumber)

	if ptr != nil {
		fmt.Println(*ptr)
	} else {
		fmt.Println("Паник")
	}

	message := messageStruct{firstWord: "Hello,", secondWord: "World!"}
	ptrMessage = &message

	if ptrMessage != nil {
		fmt.Println(*ptrMessage)
	} else {
		fmt.Println("Паник")
	}

	editMessage(ptrMessage)
	fmt.Println(*ptrMessage)
}

func editMessage(message *messageStruct) {
	message.firstWord = "GoodBye,"
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