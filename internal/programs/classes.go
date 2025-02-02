package program

import (
	e "go-gamedev-practice/internal/entities"
	"go-gamedev-practice/internal/utils"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func Classes() {
	window := utils.Coords{
		X: 800,
		Y: 450,
	}

	rl.InitWindow(int32(window.X), int32(window.Y), "raylib [core] example - basic window")
	defer rl.CloseWindow()
	rl.SetTargetFPS(60)

	player := e.PlayerCreate(500, 200, 400)
	block := e.BlockCreate(500, 200, 400)
	sprites := []e.Entity{
		&player,
		&block,
	}
	for !rl.WindowShouldClose() {

		rl.BeginDrawing()
		rl.ClearBackground(rl.DarkGray)
		dt := rl.GetFrameTime()
		e.Update(sprites, dt)
		e.Render(sprites)
		rl.DrawFPS(0, 0)
		rl.EndDrawing()
	}
}
