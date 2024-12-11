package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var num int = rand.Intn(100)
	var guess int = -1

	for guess != num {
		fmt.Println("Enter a number: ")
		fmt.Scanln(&guess)
		if guess < num {
			fmt.Println("Too low!")
		} else if guess > num {
			fmt.Println("Too high!")
		} else {
			fmt.Println("You got it!")
		}
	}
}