package main

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

// RoverPosition defines where the hover is withe the facing
type RoverPosition struct {
	Position
	Facing rune
}

// New RoverPosition factory based in string with x position (int) y position (int) and facing (char)
// The facing available is N for North, S for South, W for West and E for East otherwise will panic
// Input Example: "1 2 N"
func New(st string) *RoverPosition {
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

	return &RoverPosition{
		Position: Position{
			XPosition: x,
			YPosition: y,
		},
		Facing: facingSl[0],
	}
}

func (rp *RoverPosition) String() string{
	return fmt.Sprintf("%d %d %s", rp.XPosition, rp.YPosition, string(rp.Facing))
}


// Moves Hover based in a string
// moves are a string with the commands to your hover route
// The possible commands are 'L', 'R' and 'M'
// Example: "LMLMLMLMM"
// rt is the right and top position in plateau
func (rp *RoverPosition) Move(moves string, rt Position) {
	for _, m := range moves {
		switch m {
		case Right:
			rp.turnRight()
		case Left:
			rp.turnLeft()
		case Move:
			rp.walk(rt)
		default:
			panic(fmt.Sprintf("%s%s", "Invalid Movement to ", string(m)))
		}
	}
}

// Moves the rover position to the facing direction
func (rp *RoverPosition) walk(rt Position) {
	switch rp.Facing {
	case East:
		if rp.XPosition >= rt.XPosition {
			panic(" Your rover is out of the plateau, cant move to the East")
		}
		rp.XPosition++
	case North:
		if rp.YPosition >= rt.YPosition {
			panic(" Your rover is out of the plateau, cant move to the North")
		}
		rp.YPosition++
	case West:
		if rp.XPosition <= 0 {
			panic(" Your rover is out of the plateau, cant move to the West")
		}
		rp.XPosition--
	case South:
		if rp.YPosition <= 0 {
			panic(" Your rover is out of the plateau, cant move to the South")
		}
		rp.YPosition--
	}
}

// changes the rover facing when it turns the right
func (rp *RoverPosition) turnRight() {
	switch rp.Facing {
	case East:
		//fNew := South
		rp.Facing = South
	case South:
		rp.Facing = West
	case West:
		rp.Facing = North
	case North:
		rp.Facing = East
	default:
		panic(" Incorrect Facing string format ")
	}
}

//changes the rover facing when it turns the left
func (rp *RoverPosition) turnLeft() {
	switch rp.Facing {
	case East:
		rp.Facing = North
	case North:
		rp.Facing = West
	case West:
		rp.Facing = South
	case South:
		rp.Facing = East
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
		panic(" Incorrect Position string format ")
	}

	x, err := strconv.Atoi(sl[0])
	if err != nil {
		panic(" Incorrect x position string format ")
	}

	y, err := strconv.Atoi(sl[1])
	if err != nil {
		panic(" Incorrect y position string format ")
	}

	return Position{
		XPosition: x,
		YPosition: y,
	}
}