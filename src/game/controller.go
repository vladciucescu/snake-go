package game

import (
	"log"
	"math/rand"
	"snake/src/config"
	"time"
)

var snakeBoard board
var snakey *snake
var random = rand.New(rand.NewSource(time.Now().UnixNano()))

func init() {
	snakeBoard = newBoard(config.GetBoardDimensions())
	snakey = newSnake(config.GetStartingPosition())
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
	pos, _ := getNewApplePosition(snakeBoard)
	snakeBoard.placeObject(pos, apple)
}

func getNewApplePosition(b board) (position, bool) {
	emptyPositions := b.getEmptyPositions()
	size := len(emptyPositions)
	var tries int
	index := random.Intn(size)
	for hasAppleNeighbor(snakeBoard, emptyPositions[index]) && tries<size {
		tries++
		index = random.Intn(size)
	}
	if tries == size {
		return position{}, false
	}
	return emptyPositions[index], true
}

func hasAppleNeighbor(b board, p position) bool {
	neighbors := b.getNeighbors(p)
	for _, pos := range neighbors {
		if b.getObject(pos) == apple {
			return true
		}
	}
	return false
}

func executeCommand(cmd command, arg int) {

}
