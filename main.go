package main

import (
	"fmt"
	"io"
)

func main() {
	actions := map[string]func(){
		"help":     help,
		"add":      add,
		"subtract": subtract,
		"multiply": multiply,
	}
	fmt.Println("Welcome!")
	help()
	for {
		var command string
		_, err := fmt.Scan(&command)
		if err != nil {
			if err != io.EOF {
				panic(err)
			}
			break
		}
		action, ok := actions[command]
		if !ok {
			help()
			continue
		}
		action()
		fmt.Println("Next command?")
	}
}

func help() {
	fmt.Println(`You can type:
add
    For adding things!
subtract
    For subtracting things!
multiply
    I bet you can guess!
help
    To display this dialog

`)
}

func add() {
	first, second := getTwoFloatsFromUser()
	fmt.Println(first + second)
}

func subtract() {
	first, second := getTwoFloatsFromUser()
	fmt.Println(first - second)
}

func multiply() {
	first, second := getTwoFloatsFromUser()
	fmt.Println(first * second)
}

func getTwoFloatsFromUser() (firstFloat float64, secondFloat float64) {
	for {
		fmt.Println("Enter the first number you would like to add:")
		_, err := fmt.Scan(&firstFloat)
		if err != nil {
			fmt.Println("You must enter a number.")
			continue
		}
		break
	}
	for {
		fmt.Println("Enter the second number you would like to add:")
		_, err := fmt.Scan(&secondFloat)
		if err != nil {
			fmt.Println("You must enter a number.")
			continue
		}
		break
	}
	return
}
