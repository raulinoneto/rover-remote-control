package main

import (
	"fmt"
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestHoverPosition_Move(t *testing.T) {
	rt := NewPosition("5 5")

	rp := New("1 2 N")
	rp.Move("LMLMLMLMM", rt)
	assert.Equal(t, rp.XPosition, 1)
	assert.Equal(t, rp.YPosition, 3)
	assert.Equal(t, rp.Facing, North)

	rp = New("3 3 E")
	rp.Move("MMRMMRMRRM", rt)
	assert.Equal(t, rp.XPosition, 5)
	assert.Equal(t, rp.YPosition, 1)
	assert.Equal(t, rp.Facing, East)

	rp = New("3 1 W")
	rp.Move("LMLMMLMM", rt)
	assert.Equal(t, rp.XPosition, 5)
	assert.Equal(t, rp.YPosition, 2)
	assert.Equal(t, rp.Facing, North)
}

func TestHoverPosition_Walk(t *testing.T) {
	rt := NewPosition("5 5")

	rp := &RoverPosition{
		Position: Position{
			XPosition: 1,
			YPosition: 2,
		},
		Facing: East,
	}
	rp.walk(rt)
	assert.Equal(t, rp.XPosition, 2)
	rp.Facing = North
	rp.walk(rt)
	assert.Equal(t, rp.YPosition, 3)
	rp.Facing = West
	rp.walk(rt)
	assert.Equal(t, rp.XPosition, 1)
	rp.Facing = South
	rp.walk(rt)
	assert.Equal(t, rp.YPosition, 2)
}

func TestHoverPosition_TurnLeft(t *testing.T) {
	rp := &RoverPosition{
		Facing:   East,
	}
	rp.turnLeft()
	assert.Equal(t, rp.Facing, North)
	rp.turnLeft()
	assert.Equal(t, rp.Facing, West)
	rp.turnLeft()
	assert.Equal(t, rp.Facing, South)
	rp.turnLeft()
	assert.Equal(t, rp.Facing, East)
}

func TestFacing_TurnRight(t *testing.T) {
	rp := &RoverPosition{
		Facing:   East,
	}
	rp.turnRight()
	assert.Equal(t, rp.Facing, South)
	rp.turnRight()
	assert.Equal(t, rp.Facing, West)
	rp.turnRight()
	assert.Equal(t, rp.Facing, North)
	rp.turnRight()
	assert.Equal(t, rp.Facing, East)
}

func ExampleHoverPosition_Move() {
	rt := NewPosition("5 5")

	rp := New("1 2 N")
	rp.Move("LMLMLMLMM",rt)
	fmt.Printf("%+v\n", *rp)

	rp = New("3 3 E")
	rp.Move("MMRMMRMRRM", rt)
	fmt.Printf("%+v\n", *rp)

	//Output: {Position:{XPosition:1 YPosition:3} Facing:78}
	//{Position:{XPosition:5 YPosition:1} Facing:69}
}

func BenchmarkHoverPosition_Move(b *testing.B) {
	rt := NewPosition("5 5")

	rp := New("1 2 N")
	rp.Move("LMLMLMLMM",rt)

	rp = New("3 3 E")
	rp.Move("MMRMMRMRRM", rt)
}