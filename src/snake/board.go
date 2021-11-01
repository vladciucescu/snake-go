package snake

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
