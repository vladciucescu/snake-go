package game

type direction int

const (
	north direction = iota
	south
	west
	east
)

var directions = []direction{north, south, west, east}

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
