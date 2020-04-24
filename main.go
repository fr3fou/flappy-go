package main

import rl "github.com/gen2brain/raylib-go/raylib"

const (
	width  = 800
	height = 900
)

func main() {
	rl.InitWindow(width, height, "Flappy Bird!")
	g := NewGame()

	rl.SetTargetFPS(60)

	hasStarted := false
	for !rl.WindowShouldClose() {
		if hasStarted {
			g.Update()
		}
		g.Draw()

		if !hasStarted && rl.IsKeyPressed(rl.KeySpace) {
			hasStarted = true
			continue
		}
	}
}
