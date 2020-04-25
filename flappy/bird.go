package flappy

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
	rl.DrawRectangleRec(bird.Rectangle, rl.Yellow)
}

func (bird *Bird) Update() {
	if bird.Velocity < 9 {
		bird.Velocity += gravity
	}

	bird.Y += bird.Velocity
}

func (b *Bird) Jump() {
	b.Velocity = -jump
}
