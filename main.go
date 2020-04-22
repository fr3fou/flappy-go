package main

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	Width  = 750
	Height = 750
)

func main() {
	rl.InitWindow(Width, Height, "Flappy Bird!")
	g := NewGame()

	for !rl.WindowShouldClose() {
		g.Draw()
	}
}
