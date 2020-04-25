package flappy

import rl "github.com/gen2brain/raylib-go/raylib"

// Ground is ground
type Ground struct {
	rl.Rectangle
}

const (
	GroundHeight = 100
	GrassHeight  = 5
)

// NewGround is a ctor for ground
func NewGround() *Ground {
	return &Ground{
		Rectangle: rl.NewRectangle(0, Height-GroundHeight, Width, GroundHeight),
	}
}

func (ground *Ground) Draw() {
	g := ground.Rectangle.ToInt32()

	// Dirt
	rl.DrawRectangle(g.X, g.Y, g.Width, g.Height, rl.Brown)

	// Grass
	rl.DrawRectangle(g.X, g.Y, Width, GrassHeight, rl.DarkGreen)
}

func (ground *Ground) CollidesWith(other rl.Rectangle) bool {
	return rl.CheckCollisionRecs(ground.Rectangle, other)
}
