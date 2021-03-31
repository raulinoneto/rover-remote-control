package rover_remote_control

import (
	"fmt"
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestHoverPosition_Move(t *testing.T) {
	rt := NewPosition("5 5")

	hp := New("1 2 N")
	hp.Move("LMLMLMLMM", rt)
	assert.Equal(t, hp.XPosition, 1)
	assert.Equal(t, hp.YPosition, 3)
	assert.Equal(t, hp.Facing, North)

	hp = New("3 3 E")
	hp.Move("MMRMMRMRRM", rt)
	assert.Equal(t, hp.XPosition, 5)
	assert.Equal(t, hp.YPosition, 1)
	assert.Equal(t, hp.Facing, East)

	hp = New("3 1 W")
	hp.Move("LMLMMLMM", rt)
	assert.Equal(t, hp.XPosition, 5)
	assert.Equal(t, hp.YPosition, 2)
	assert.Equal(t, hp.Facing, North)
}

func TestHoverPosition_Walk(t *testing.T) {
	rt := NewPosition("5 5")

	hp := &HoverPosition{
		Position: Position{
			XPosition: 1,
			YPosition: 2,
		},
		Facing: East,
	}
	hp.walk(rt)
	assert.Equal(t, hp.XPosition, 2)
	hp.Facing = North
	hp.walk(rt)
	assert.Equal(t, hp.YPosition, 3)
	hp.Facing = West
	hp.walk(rt)
	assert.Equal(t, hp.XPosition, 1)
	hp.Facing = South
	hp.walk(rt)
	assert.Equal(t, hp.YPosition, 2)
}

func TestHoverPosition_TurnLeft(t *testing.T) {
	hp := &HoverPosition{
		Facing:   East,
	}
	hp.turnLeft()
	assert.Equal(t, hp.Facing, North)
	hp.turnLeft()
	assert.Equal(t, hp.Facing, West)
	hp.turnLeft()
	assert.Equal(t, hp.Facing, South)
	hp.turnLeft()
	assert.Equal(t, hp.Facing, East)
}

func TestFacing_TurnRight(t *testing.T) {
	hp := &HoverPosition{
		Facing:   East,
	}
	hp.turnRight()
	assert.Equal(t, hp.Facing, South)
	hp.turnRight()
	assert.Equal(t, hp.Facing, West)
	hp.turnRight()
	assert.Equal(t, hp.Facing, North)
	hp.turnRight()
	assert.Equal(t, hp.Facing, East)
}

func ExampleHoverPosition_Move() {
	rt := NewPosition("5 5")

	hp := New("1 2 N")
	hp.Move("LMLMLMLMM",rt)
	fmt.Printf("%+v\n", *hp)

	hp = New("3 3 E")
	hp.Move("MMRMMRMRRM", rt)
	fmt.Printf("%+v\n", *hp)

	//Output: {Position:{XPosition:1 YPosition:3} Facing:78}
	//{Position:{XPosition:5 YPosition:1} Facing:69}
}

func BenchmarkHoverPosition_Move(b *testing.B) {
	rt := NewPosition("5 5")

	hp := New("1 2 N")
	hp.Move("LMLMLMLMM",rt)

	hp = New("3 3 E")
	hp.Move("MMRMMRMRRM", rt)
}