package rover_remote_control

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	// Facing East Direction
	East  rune = 'E'
	// Facing West Direction
	West  rune = 'W'
	// Facing North Direction
	North rune = 'N'
	// Facing South Direction
	South rune = 'S'

	// Right Movement
	Right rune = 'R'
	// Left Movement
	Left  rune = 'L'
	// Move to next position in plateau
	Move  rune = 'M'
)

// HoverPosition defines where the hover is withe the facing
type HoverPosition struct {
	Position
	Facing rune
}

// New HoverPosition factory based in string with x position (int) y position (int) and facing (char)
// The facing available is N for North, S for South, W for West and E for East otherwise will panic
// Input Example: "1 2 N"
func New(st string) *HoverPosition {
	sl := strings.Split(st, " ")
	if len(sl) > 3 {
		panic(" Incorrect HoverPosition string format ")
	}

	x, err := strconv.Atoi(sl[0])
	if err != nil {
		panic(" Incorrect HoverPosition x position string format ")
	}

	y, err := strconv.Atoi(sl[1])
	if err != nil {
		panic(" Incorrect HoverPosition y position string format ")
	}

	facingSl := []rune(sl[2])
	if len(facingSl) > 1 {
		panic(" Incorrect HoverPosition facing string format ")
	}

	return &HoverPosition{
		Position: Position{
			XPosition: x,
			YPosition: y,
		},
		Facing: facingSl[0],
	}
}

// Moves Hover based in a string
// moves are a string with the commands to your hover route
// The possible commands are 'L', 'R' and 'M'
// Example: "LMLMLMLMM"
// rt is the right and top position in plateau
func (hp *HoverPosition) Move(moves string, rt Position) {
	for _, m := range moves {
		switch m {
		case Right:
			hp.turnRight()
		case Left:
			hp.turnLeft()
		case Move:
			hp.walk(rt)
		default:
			panic(fmt.Sprintf("%s%s", "Invalid Movement to ", string(m)))
		}
	}
}

// Moves the rover position to the facing direction
func (hp *HoverPosition) walk(rt Position) {
	switch hp.Facing {
	case East:
		if hp.XPosition >= rt.XPosition {
			panic(" Your rover is out of the plateau, cant move to the East")
		}
		hp.XPosition++
	case North:
		if hp.YPosition >= rt.YPosition {
			panic(" Your rover is out of the plateau, cant move to the North")
		}
		hp.YPosition++
	case West:
		if hp.XPosition <= 0 {
			panic(" Your rover is out of the plateau, cant move to the West")
		}
		hp.XPosition--
	case South:
		if hp.YPosition <= 0 {
			panic(" Your rover is out of the plateau, cant move to the South")
		}
		hp.YPosition--
	}
}

// changes the rover facing when it turns the right
func (hp *HoverPosition) turnRight() {
	switch hp.Facing {
	case East:
		//fNew := South
		hp.Facing = South
	case South:
		hp.Facing = West
	case West:
		hp.Facing = North
	case North:
		hp.Facing = East
	default:
		panic(" Incorrect Facing string format ")
	}
}

//changes the rover facing when it turns the left
func (hp *HoverPosition) turnLeft() {
	switch hp.Facing {
	case East:
		hp.Facing = North
	case North:
		hp.Facing = West
	case West:
		hp.Facing = South
	case South:
		hp.Facing = East
	default:
		panic(" Incorrect Facing string format ")
	}
}

// Position in plateau
type Position struct {
	XPosition int
	YPosition int
}

// NewPosition factory based in string with x position (int) y position (int)
// Input Example: "5 5"
func NewPosition(st string) Position {
	sl := strings.Split(st, " ")
	if len(sl) > 2 {
		panic(" Incorrect HoverPosition string format ")
	}

	x, err := strconv.Atoi(sl[0])
	if err != nil {
		panic(" Incorrect HoverPosition x position string format ")
	}

	y, err := strconv.Atoi(sl[1])
	if err != nil {
		panic(" Incorrect HoverPosition y position string format ")
	}

	return Position{
		XPosition: x,
		YPosition: y,
	}
}