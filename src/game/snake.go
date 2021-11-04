package game

type direction int
type position struct {
	row, column int
}

const (
	north direction = iota
	south
	west
	east
)

type snake struct {
	direction
	segments []position
}

func newSnake(startingRow, startingColumn int) *snake {
	dir := north
	var segments = make([]position, 0)
	head := position{startingColumn, startingRow}
	xIndex, yIndex := dir.getMoveIndexes()
	firstSegment := position{startingRow - xIndex, startingColumn - yIndex}
	secondSegment := position{startingRow - xIndex*2, startingColumn - yIndex*2}
	segments = append(segments, head)
	segments = append(segments, firstSegment)
	segments = append(segments, secondSegment)
	return &snake{dir, segments}
}

func (d direction) getMoveIndexes() (x, y int) {
	switch d {
	case north:
		return -1, 0
	case south:
		return 1, 0
	case west:
		return 0, -1
	case east:
		return 0, 1
	default:
		panic("Invalid direction")
	}
}
