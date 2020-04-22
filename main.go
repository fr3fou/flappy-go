package main

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	Width  = 800
	Height = 900
)

func main() {
	rl.InitWindow(Width, Height, "Flappy Bird!")
	g := NewGame()

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		g.Update()
		g.Draw()
	}
}
