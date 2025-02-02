package program

import (
	"go-gamedev-practice/internal/utils"
	"path/filepath"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func bouncer() {
	window := utils.Coords{
		X: 800,
		Y: 450,
	}
	offset := utils.Coords{
		X: 112,
		Y: 75,
	}
	rl.InitWindow(800, 450, "raylib [core] example - basic window")
	defer rl.CloseWindow()

	ship := rl.LoadTexture(filepath.Join("assets", "spaceship.png"))
	shipPos := rl.Vector2{X: 0, Y: 0}
	shipDirection := rl.Vector2{X: 1, Y: 1}
	shipSpeed := 200

	rl.SetTargetFPS(60)
	for !rl.WindowShouldClose() {
		// updates
		dt := rl.GetFrameTime()
		shipPos.X += shipDirection.X * float32(shipSpeed) * dt
		shipPos.Y += shipDirection.Y * float32(shipSpeed) * dt
		if shipPos.X+offset.X > window.X {
			shipDirection.X = -1
		}
		if shipPos.Y+offset.Y > window.Y {
			shipDirection.Y = -1
		}
		if shipPos.X < 0 {
			shipDirection.X = 1
		}
		if shipPos.Y < 0 {
			shipDirection.Y = 1
		}

		rl.BeginDrawing()
		rl.ClearBackground(rl.DarkGray)
		rl.DrawTextureV(ship, shipPos, rl.White)
		rl.DrawFPS(0, 0)
		rl.EndDrawing()
	}
}
