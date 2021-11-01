package snake

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"text/tabwriter"
)

var tw = tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', tabwriter.Debug)

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
	placeSnake()
	placeApple()
	printBoard()
}

func processCommands() {
	const goodbyeMessage = "We are sorry to see you go.."
	const invalidCommand = "Invalid command, please try again."
	reader := bufio.NewReader(os.Stdin)
loop:
	for {
		fmt.Print("Command: ")
		text, _ := reader.ReadString('\n')
		command, arg := parseCommand(text)

		switch command {
		case exit:
			fmt.Println(goodbyeMessage)
			break loop
		case invalid:
			fmt.Println(invalidCommand)
			continue
		default:
			executeCommand(command, arg)
		}
		printBoard()
	}
}

// parseCommand parses the given text and returns a command and an argument to use with that command.
// If no valid argument was found, 0 is returned
func parseCommand(text string) (cmd command, arg int) {
	trimmed := strings.TrimSpace(text)
	split := strings.Split(trimmed, " ")
	if len(split) < 1 {
		return invalid, 0
	}
	cmd = getCommand(split[0])
	if len(split) > 1 {
		arg, _ = strconv.Atoi(split[1])
	}
	return cmd, arg
}

func printBoard() {
	board := getBoard()
	for column := 1; column <= board.columns; column++ {
		_, _ = fmt.Fprint(tw, column)
	}
	for rowNr, row := range board.objects {
		rowSymbols := Map(row, boardObject.String)
		_, _ = fmt.Fprintf(tw, "%d %v", rowNr, strings.Join(rowSymbols, " "))
	}
}

func Map(vs []boardObject, f func(object boardObject) string) []string {
	vsm := make([]string, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}
