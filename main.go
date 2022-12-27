package main

import (
	"fmt"
)

func main() {
	const answer = 27
	var number int
	fmt.Printf("please enter 0-100 value:")
	fmt.Scanln(&number)
	if answer == number {
		fmt.Println("number=", number, "Congratulations!")
	} else if number < answer {
		fmt.Println("to small!")
	} else if number > answer {
		fmt.Println("to big!")
	}
}
