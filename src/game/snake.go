package game

import "fmt"

type position struct {
	row, column int
}

func (p position) String() string {
	return fmt.Sprintf("{%d %d}", p.row+1, p.column+1)
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

func (s *snake) getHead() position {
	return s.segments[0]
}

func (s *snake) getBody() []position {
	return s.segments[1:]
}

func (s *snake) getNextHeadPos() position {
	head := s.getHead()
	x, y := s.direction.getMoveIndexes()
	return position{head.row + x, head.column + y}
}

func (s *snake) growTail() {
	tail := s.segments[len(s.segments)-1]
	tailDir := s.getTailDirection()
	x, y := tailDir.getMoveIndexes()
	newTail := position{tail.row + x, tail.column + y}
	s.segments = append(s.segments, newTail)
}

func (s *snake) getTailDirection() direction {
	currTail := s.segments[len(s.segments)-1]
	prevTail := s.segments[len(s.segments)-2]
	if currTail.row == prevTail.row && currTail.column > prevTail.column {
		return east
	}
	if currTail.row == prevTail.row && currTail.column < prevTail.column {
		return west
	}
	if currTail.column == prevTail.column && currTail.row > prevTail.row {
		return south
	}
	if currTail.column == prevTail.column && currTail.row < prevTail.row {
		return north
	}
	panic("Invalid tail")
}

func (s *snake) move() {
	s.segments = append([]position{s.getNextHeadPos()}, s.segments...)
	s.segments = s.segments[:len(s.segments)-1]
}
