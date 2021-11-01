package game

type command int

const (
	move command = iota
	up
	down
	left
	right
	exit
	invalid
)

var menuOptions = map[command]string{
	move:    "move <n>",
	up:      "up",
	down:    "down",
	left:    "left",
	right:   "right",
	exit:    "exit",
}

func getCommand(text string) command {
	switch text {
	case "move":
		return move
	case "up":
		return up
	case "down":
		return down
	case "left":
		return left
	case "right":
		return right
	case "exit":
		return exit
	default:
		return invalid
	}
}
