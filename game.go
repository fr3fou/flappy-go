package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

// Game is a game
type Game struct {
	Bird   *Bird
	Ground *Ground
	Over   bool
	Pipes  []*Pipe
}

func NewGame() *Game {
	g := &Game{}
	g.Init()
	return g
}

func (g *Game) Init() {
	g.Ground = NewGround()
	g.Bird = NewBird(birdSize*2, Height/2)

	g.Pipes = make([]*Pipe, 100)
	initialOffset := horizontalGap + birdSize*2*2
	offset := initialOffset
	for i := range g.Pipes {
		g.Pipes[i] = NewPipe((i)*pipeWidth+offset, 0)
		offset = horizontalGap*(i+1) + initialOffset
	}
}

func (g *Game) Update() {
	if g.Over {
		return
	}
	if rl.IsKeyReleased(rl.KeySpace) {
		g.Bird.Jump()
	}

	for i := range g.Pipes {
		if rl.CheckCollisionRecs(g.Bird.Rectangle, g.Pipes[i].Rectangle) {
			g.Over = true
			break
		}
		g.Pipes[i].X -= speed
	}

	g.Bird.Update()
}

func (g *Game) Draw() {
	rl.BeginDrawing()
	rl.ClearBackground(rl.SkyBlue)

	g.Bird.Draw()

	for _, pipe := range g.Pipes {
		pipe.Draw()
	}

	g.Ground.Draw()

	if g.Over {
		rl.DrawText("GAME OVER!", Width/2-55*2, Height/2-55, 55, rl.White)
	}

	rl.EndDrawing()
}
