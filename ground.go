package main

import rl "github.com/gen2brain/raylib-go/raylib"

// Ground is ground
type Ground struct {
	rl.Rectangle
}

const (
	groundHeight = 100
	grassHeight  = 5
)

// NewGround is a ctor for ground
func NewGround() *Ground {
	return &Ground{
		Rectangle: rl.NewRectangle(0, Height-groundHeight, Width, groundHeight),
	}
}

func (ground *Ground) Draw() {
	g := ground.Rectangle.ToInt32()

	// Dirt
	rl.DrawRectangle(g.X, g.Y, g.Width, g.Height, rl.Brown)

	// Grass
	rl.DrawRectangle(g.X, g.Y, Width, grassHeight, rl.DarkGreen)
}
