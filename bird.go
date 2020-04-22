package main

import rl "github.com/gen2brain/raylib-go/raylib"

// Bird is a bird
type Bird struct {
	rl.Rectangle
	Velocity float32
}

const (
	birdSize = 30
	gravity  = 0.35
	jump     = 6.9
)

func NewBird(x, y int) *Bird {
	return &Bird{
		Rectangle: rl.NewRectangle(float32(x), float32(y), birdSize, birdSize),
		Velocity:  1,
	}
}

func (bird *Bird) Draw() {
	b := bird.ToInt32()
	rl.DrawRectangle(b.X, b.Y, b.Width, b.Height, rl.Yellow)
}

func (b *Bird) Update() {
	if b.Velocity < 9 {
		b.Velocity += gravity
	}

	b.Y += b.Velocity
}

func (b *Bird) Jump() {
	b.Velocity = -jump
}
