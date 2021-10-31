package ui

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

var menuOptions = map[command]string {
	move: "move <n>",
	up: "up",
	down: "down",
	left: "left",
	right: "right",
	exit: "exit",
	invalid: "",
}


