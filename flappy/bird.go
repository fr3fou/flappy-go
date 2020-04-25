package flappy

import rl "github.com/gen2brain/raylib-go/raylib"

// Bird is a bird
type Bird struct {
	rl.Rectangle
	Velocity float32
}

const (
	BirdSize = 30
	Gravity  = 0.35
	Jump     = 6.9
)

func NewBird(x, y int) *Bird {
	return &Bird{
		Rectangle: rl.NewRectangle(float32(x), float32(y), BirdSize, BirdSize),
		Velocity:  1,
	}
}

func (bird *Bird) Draw() {
	rl.DrawRectangleRec(bird.Rectangle, rl.Yellow)
}

func (bird *Bird) Update() {
	if bird.Velocity < 9 {
		bird.Velocity += Gravity
	}

	bird.Y += bird.Velocity
}

func (b *Bird) Jump() {
	b.Velocity = -Jump
}

func (b *Bird) AboveSky() bool {
	return b.Y < 0
}
