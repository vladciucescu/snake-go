package game

import "log"

type boardObject int

const (
	empty boardObject = iota
	apple
	snakeHead
	snakeSegment
)

func (bo boardObject) String() string {
	switch bo {
	case empty:
		return " "
	case apple:
		return "o"
	case snakeHead:
		return "*"
	case snakeSegment:
		return "+"
	default:
		return "?"
	}
}

type board struct {
	rows, columns int
	objects       [][]boardObject
}

func newBoard(rows, columns int) board {
	objects := make([][]boardObject, rows)
	for row := 0; row < rows; row++ {
		objects[row] = make([]boardObject, columns)
	}
	return board{rows, columns, objects}
}

func (b board) isOnBoard(pos position) bool {
	return 0 <= pos.row &&
		pos.row < b.rows &&
		0 <= pos.column &&
		pos.column < b.columns
}

func (b board) placeSnake(s *snake) {
	b.removeSnake()
	b.placeObject(s.getHead(), snakeHead)
	for _, segment := range s.getBody() {
		b.placeObject(segment, snakeSegment)
	}
}

func (b board) removeSnake() {
	for i := 0; i < b.rows; i++ {
		for j := 0; j < b.columns; j++ {
			if b.objects[i][j] == snakeHead || b.objects[i][j] == snakeSegment {
				b.objects[i][j] = empty
			}
		}
	}
}

func (b board) placeObject(pos position, object boardObject) {
	if !b.isOnBoard(pos) {
		log.Panicf("Position %v is not on board", pos)
	}
	b.objects[pos.row][pos.column] = object
}
