package game

import "snake/src/config"

var snakeBoard board
var snakey *snake

func init() {
	snakeBoard = newBoard(config.GetBoardDimensions())
	snakey = newSnake(config.GetStartingPosition())
}

func getBoard() board {
	return snakeBoard
}

func placeSnake() {
	//validateInitialPosition(snake)
	//board.placeSnake(snake)
}

func placeApple() {

}

func executeCommand(cmd command, arg int) {

}
