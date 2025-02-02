package program

import (
	"fmt"
	"go-gamedev-practice/internal/utils"
	"path/filepath"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type AnimatedSprite struct {
	frameIdx float32
	frames   []rl.Texture2D
	speed    float32
	name     string
}

func (a *AnimatedSprite) Update(dt float32) {
	a.frameIdx += 5 * dt
}
func (a AnimatedSprite) Render() {
	rl.DrawTexture(a.frames[int(a.frameIdx)%len(a.frames)], 0, 0, rl.White)
}
func AnimatedSpriteCreate(name string, frames int, speed float32) AnimatedSprite {
	return AnimatedSprite{
		name:   name,
		speed:  speed,
		frames: GenerateFrames(name, frames),
	}
}
func GenerateFrames(name string, frames int) []rl.Texture2D {
	var animationFrames []rl.Texture2D
	for i := 0; i < frames; i++ {
		frame := rl.LoadTexture(filepath.Join("..", "..", "internal", "assets", "animation", fmt.Sprintf("%v%d.png", name, i)))
		animationFrames = append(animationFrames, frame)
	}
	return animationFrames
}

func Animations() {
	window := utils.Coords{
		X: 1000,
		Y: 640,
	}

	rl.InitWindow(int32(window.X), int32(window.Y), "raylib animations")
	defer rl.CloseWindow()

	ani := AnimatedSpriteCreate("", 8, 5)

	rl.SetTargetFPS(60)
	for !rl.WindowShouldClose() {
		rl.SetExitKey(rl.KeyEscape)
		// updates
		dt := rl.GetFrameTime()
		ani.Update(dt)

		rl.BeginDrawing()
		rl.ClearBackground(rl.DarkGray)
		ani.Render()
		rl.DrawFPS(0, 0)
		rl.EndDrawing()
	}
}
