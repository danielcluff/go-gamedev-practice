package program

import (
	"path/filepath"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func rld() {
	rl.InitWindow(800, 450, "raylib [core] example - basic window")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	spaceship := rl.LoadTexture(filepath.Join("assets", "spaceship.png"))
	spaceship_image := rl.LoadImage(filepath.Join("assets", "spaceship.png"))
	rl.ImageColorGrayscale(spaceship_image)
	newShip := rl.LoadTextureFromImage(spaceship_image)

	cowbowImage := rl.LoadImage("assets/animation/0.png")
	rl.ImageColorInvert(cowbowImage)
	cowbow := rl.LoadTextureFromImage(cowbowImage)

	darkRed := rl.NewColor(80, 11, 12, 255)

	dafont := rl.LoadFont(filepath.Join("assets", "Zero Hour.otf"))

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.DarkGray)

		rl.DrawText("Wow this is so cool!", 190, 200, 20, rl.White)

		rl.DrawLineEx(rl.Vector2{X: 0, Y: 0}, rl.Vector2{X: 500, Y: 200}, 10, darkRed)
		rl.DrawTexture(spaceship, 120, 120, rl.White)
		rl.DrawTexture(newShip, 240, 240, rl.White)
		rl.DrawTexture(cowbow, 700, 300, rl.White)
		rl.DrawTextEx(dafont, "Some more text", rl.Vector2{X: 200, Y: 240}, 20, 0, rl.Lime)

		rl.EndDrawing()
	}
}
