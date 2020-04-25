package flappy

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

// Pipe is a pipe
type Pipe struct {
	rl.Rectangle
	HasPassedThrough bool
}

const (
	PipeWidth     = 70
	PipeHeight    = Height / 2
	PipeBorder    = 3
	VerticalGap   = 120
	HorizontalGap = 230
	Speed         = 2.3
)

func NewPipe(x, y int) *Pipe {
	return &Pipe{
		Rectangle: rl.NewRectangle(
			float32(x), float32(y),
			PipeWidth, float32(rl.GetRandomValue(VerticalGap, PipeHeight)),
		),
	}
}

func (pipe *Pipe) Draw() {
	// TODO: use sprites instead of gradients!
	p := pipe.ToInt32()

	// Top Pipe Border
	bTop := rl.NewRectangle(
		pipe.X-PipeBorder, pipe.Y,
		pipe.Width+PipeBorder*2, pipe.Height+PipeBorder,
	)
	rl.DrawRectangleLinesEx(bTop, PipeBorder, rl.DarkGray)

	// Top Pipe
	gTop := pipe.Rectangle
	rl.DrawRectangleGradientEx(gTop, rl.Lime, rl.Lime, rl.Green, rl.Green)

	// Bottom Pipe Border
	bBottom := rl.NewRectangle(
		pipe.X-PipeBorder, pipe.Y+VerticalGap+pipe.Height-PipeBorder,
		pipe.Width+PipeBorder*2, Height-pipe.Height+PipeBorder,
	)
	rl.DrawRectangleLinesEx(bBottom, PipeBorder, rl.DarkGray)

	// Bottom Pipe
	gBottom := rl.NewRectangle(
		pipe.X, float32(pipe.Y+VerticalGap+pipe.Height),
		pipe.Width, float32(Height-p.Height),
	)
	rl.DrawRectangleGradientEx(gBottom, rl.Lime, rl.Lime, rl.Green, rl.Green)
}

func (p *Pipe) CollidesWith(other rl.Rectangle) bool {
	top := rl.NewRectangle(
		p.X-PipeBorder, p.Y,
		p.Width+PipeBorder*2, p.Height+PipeBorder,
	)

	bottom := rl.NewRectangle(
		p.X-PipeBorder, p.Y+VerticalGap+p.Height-PipeBorder,
		p.Width+PipeBorder*2, Height-p.Height+PipeBorder,
	)

	return rl.CheckCollisionRecs(top, other) || rl.CheckCollisionRecs(bottom, other)
}

func (p *Pipe) IsAround(other rl.Rectangle) bool {
	if p.HasPassedThrough {
		return false
	}

	middle := rl.NewRectangle(
		p.X-PipeBorder, p.Y+p.Height+PipeBorder,
		p.Width,
		VerticalGap,
	)

	if rl.CheckCollisionRecs(middle, other) {
		p.HasPassedThrough = true
	}

	return p.HasPassedThrough
}

func (p *Pipe) IsOffscreen() bool {
	return p.X+p.Width+PipeBorder <= 0
}
