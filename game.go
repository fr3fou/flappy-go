package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

// Game is a game
type Game struct {
	Bird  *Bird
	Pipes []*Pipe
}

func NewGame() *Game {
	g := &Game{}
	g.Init()
	return g
}

func (g *Game) Init() {
	g.Bird = NewBird(Width/2-birdSize*2, Height/2)
	g.Pipes = make([]*Pipe, 50)

	for i := range g.Pipes {
		g.Pipes[i] = NewPipe((i)*pipeWidth+100*(i+1), 0)
	}
}

func (g *Game) Draw() {
	rl.BeginDrawing()
	rl.ClearBackground(rl.SkyBlue)

	b := g.Bird.ToInt32()
	rl.DrawRectangle(b.X, b.Y, b.Width, b.Height, rl.Yellow)

	for _, pipe := range g.Pipes {
		p := pipe.ToInt32()
		rl.DrawRectangle(p.X, p.Y, p.Width, p.Height, rl.Lime)
		rl.DrawRectangle(p.X, p.Y+pipeGap+p.Height, p.Width, Height-p.Height, rl.Lime)
	}

	rl.EndDrawing()
}
