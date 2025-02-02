package program

import (
	"fmt"
	"go-gamedev-practice/internal/utils"
	"path/filepath"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func input() {
	window := utils.Coords{
		X: 800,
		Y: 450,
	}
	// offset := Coords{
	// 	X: 112,
	// 	Y: 75,
	// }

	rl.InitWindow(int32(window.X), int32(window.Y), "raylib [core] example - basic window")
	defer rl.CloseWindow()

	ship := rl.LoadTexture(filepath.Join("assets", "spaceship.png"))
	shipPos := rl.Vector2{X: 0, Y: 0}
	shipDirection := rl.Vector2{X: 0, Y: 0}
	shipSpeed := 200

	rl.SetTargetFPS(60)
	for !rl.WindowShouldClose() {
		rl.SetExitKey(rl.KeyEscape)
		// updates
		dt := rl.GetFrameTime()
		shipPos.X += shipDirection.X * float32(shipSpeed) * dt
		shipPos.Y += shipDirection.Y * float32(shipSpeed) * dt

		// ship input
		shipDirection.X = float32(utils.BoolToInt(rl.IsKeyDown(rl.KeyRight))) - float32(utils.BoolToInt(rl.IsKeyDown(rl.KeyLeft)))
		shipDirection.Y = float32(utils.BoolToInt(rl.IsKeyDown(rl.KeyDown))) - float32(utils.BoolToInt(rl.IsKeyDown(rl.KeyUp)))
		shipDirection = rl.Vector2Normalize(shipDirection)

		lastKey := rl.GetKeyPressed()
		if lastKey != 0 {
			fmt.Println(lastKey)
		}

		// for rl.IsKeyPressed()

		rl.BeginDrawing()
		rl.ClearBackground(rl.DarkGray)
		rl.DrawTextureV(ship, shipPos, rl.White)
		rl.DrawFPS(0, 0)
		rl.EndDrawing()
	}
}
