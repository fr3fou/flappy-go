package main

import rl "github.com/gen2brain/raylib-go/raylib"

// Bird is a bird
type Bird struct {
	rl.Rectangle
	Pos rl.Vector2
}

const birdSize = 50

func NewBird(x, y float32) *Bird {
	return &Bird{
		Pos:       rl.NewVector2(x, y),
		Rectangle: rl.NewRectangle(x, y, 50, 50),
	}
}
