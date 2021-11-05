package game

import (
	"log"
	"snake/src/config"
)

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
	mustHaveValidInitialPosition(snakey)
	snakeBoard.placeSnake(snakey)
}

func mustHaveValidInitialPosition(snakey *snake) {
	if !snakeBoard.isOnBoard(snakey.getHead()) {
		panic("Snake head not on board!")
	}
	for i, pos := range snakey.getBody() {
		if !snakeBoard.isOnBoard(pos) {
			log.Panicf("Segment %d not on board!", i+1)
		}
	}
}

func placeApple() {

}

func executeCommand(cmd command, arg int) {

}
