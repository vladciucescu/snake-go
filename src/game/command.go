package game

type command int

const (
	up command = iota
	down
	left
	right
	exit
	invalid
)

var menuOptions = map[command]string{
	up:    "up <n>",
	down:  "down <n>",
	left:  "left <n>",
	right: "right <n>",
	exit:  "exit",
}

func getCommand(text string) command {
	switch text {
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
