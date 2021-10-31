package ui

import (
	"bufio"
	"fmt"
	"os"
	controller "snake/src/controller"
)

func Run() {
	showMenu()
	drawObjects()
	processCommands()
}

func showMenu() {
	for _, menuOption := range menuOptions {
		fmt.Println(menuOption)
	}
}

func drawObjects() {
	controller.PlaceSnake()
	controller.PlaceApple()
	printBoard()
}

func printBoard() {

}

func processCommands() {
	const goodbyeMessage = "We are sorry to see you go.."
	const invalidCommand = "Invalid command, please try again."
	reader := bufio.NewReader(os.Stdin)

loop:
	for {
		fmt.Print("Command: ")
		text, _ := reader.ReadString('\n')
		command, _ := parseCommand(text)

		switch command {
		case exit:
			fmt.Println(goodbyeMessage)
			break loop
		case invalid:
			fmt.Println(invalidCommand)
			continue
		default: panic("Unknown command")
		}
	}
}

// parseCommand parses the given text and returns a command and an argument to use with that command.
// If no valid argument was found, 0 is returned
func parseCommand(text string) (cmd command, arg int) {
	return exit, 0
}

