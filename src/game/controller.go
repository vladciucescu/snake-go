package game

import "snake/src/config"

var snakeBoard board

func init() {
	snakeBoard = newBoard(config.GetBoardDimensions())
}

func getBoard() board {
	return snakeBoard
}

func placeSnake() {

}

func placeApple() {

}

func executeCommand(cmd command, arg int) {

}
