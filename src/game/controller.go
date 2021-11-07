package game

import (
	"fmt"
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
	for hasAppleNeighbor(snakeBoard, emptyPositions[index]) && tries < size {
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

func executeCommand(cmd command, arg int) (gameOver bool) {
	switch cmd {
	case up:
		snakey.direction = north
	case down:
		snakey.direction = south
	case left:
		snakey.direction = west
	case right:
		snakey.direction = east
	}
	if arg > 0 {
		return advanceSnake(arg)
	} else {
		return advanceSnake(1)
	}
}

func advanceSnake(steps int) (gameOver bool) {
	for i := 0; i < steps; i++ {
		if isGameOver(snakeBoard, snakey) {
			return true
		}
		willEatApple := snakeBoard.getObject(snakey.getNextHeadPos()) == apple
		if willEatApple {
			snakey.growTail()
			placeApple()
		}
		snakey.move()
		snakeBoard.placeSnake(snakey)
	}
	return false
}

func isGameOver(b board, s *snake) bool {
	nextPos := s.getNextHeadPos()
	if !b.isOnBoard(nextPos) {
		fmt.Println("The snake hit the wall!")
		return true
	}
	if b.getObject(nextPos) == snakeSegment {
		fmt.Println("The snake ate itself!")
		return true
	}
	return false
}
