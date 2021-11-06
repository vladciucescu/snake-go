package game

type position struct {
	row, column int
}

type snake struct {
	direction
	segments []position
}

func newSnake(startingRow, startingColumn int) *snake {
	dir := north
	var segments = make([]position, 0)
	head := position{startingRow, startingColumn}
	xIndex, yIndex := dir.getMoveIndexes()
	firstSegment := position{startingRow - xIndex, startingColumn - yIndex}
	secondSegment := position{startingRow - xIndex*2, startingColumn - yIndex*2}
	segments = append(segments, head)
	segments = append(segments, firstSegment)
	segments = append(segments, secondSegment)
	return &snake{dir, segments}
}

func (s snake) getHead() position {
	return s.segments[0]
}

func (s snake) getBody() []position {
	return s.segments[1:]
}
