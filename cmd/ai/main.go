package main

import (
	"github.com/fr3fou/flappy-go/flappy"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(flappy.Width, flappy.Height, "Flappy Bird!")
	g := flappy.NewGame()

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
