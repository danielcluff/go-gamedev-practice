package program

import (
	"go-gamedev-practice/internal/utils"
	"path/filepath"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func audio() {
	window := utils.Coords{
		X: 800,
		Y: 450,
	}

	rl.InitWindow(int32(window.X), int32(window.Y), "raylib [core] example - basic window")
	defer rl.CloseWindow()
	rl.InitAudioDevice()
	defer rl.CloseAudioDevice()

	laserSound := rl.LoadSound(filepath.Join("assets", "laser.wav"))
	rl.PlaySound(laserSound)

	music := rl.LoadMusicStream(filepath.Join("assets", "music.wav"))
	defer rl.UnloadMusicStream(music)
	rl.PlayMusicStream(music)

	rl.SetTargetFPS(60)
	for !rl.WindowShouldClose() {
		rl.UpdateMusicStream((music))
		rl.BeginDrawing()
		rl.ClearBackground(rl.DarkGray)
		rl.DrawFPS(0, 0)
		rl.EndDrawing()
	}
}
