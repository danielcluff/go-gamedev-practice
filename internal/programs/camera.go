package program

import (
	"go-gamedev-practice/internal/utils"
	"math"
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func RandomChoice(slice []rl.Color) rl.Color {
	return slice[rand.Intn(len(slice))]
}

type Circle struct {
	Position rl.Vector2
	Radius   float32
	Color    rl.Color
}

func Camera() {
	window := utils.Coords{
		X: 800,
		Y: 450,
	}
	pos := rl.Vector2{}
	radius := 50
	direction := rl.Vector2{}
	speed := 400

	// circles
	circles := make([]Circle, 0, 100)
	for i := 0; i < 100; i++ {
		circle := Circle{
			Position: rl.Vector2{
				X: float32(rand.Intn(4001) - 2000),
				Y: float32(rand.Intn(2001) - 1000),
			},
			Radius: float32(50 + rand.Intn(151)),
			Color:  RandomChoice([]rl.Color{rl.Red, rl.Green, rl.Blue, rl.Yellow, rl.Orange}),
		}
		circles = append(circles, circle)
	}

	rl.InitWindow(int32(window.X), int32(window.Y), "raylib [core] example - basic window")
	defer rl.CloseWindow()

	// Camera
	camera := rl.Camera2D{}
	camera.Zoom = 1
	camera.Offset = rl.Vector2{X: window.X / 2, Y: window.Y / 2}
	camera.Rotation = 0

	rl.SetTargetFPS(60)
	for !rl.WindowShouldClose() {

		// input
		direction.X = float32(utils.BoolToInt(rl.IsKeyDown(rl.KeyRight))) - float32(utils.BoolToInt(rl.IsKeyDown(rl.KeyLeft)))
		direction.Y = float32(utils.BoolToInt(rl.IsKeyDown(rl.KeyDown))) - float32(utils.BoolToInt(rl.IsKeyDown(rl.KeyUp)))
		direction = rl.Vector2Normalize(direction)

		// Movement
		dt := rl.GetFrameTime()
		pos.X += direction.X * float32(speed) * dt
		pos.Y += direction.Y * float32(speed) * dt

		// Camera update
		rotateDirection := float32(utils.BoolToInt(rl.IsKeyDown(rl.KeyQ))) - float32(utils.BoolToInt(rl.IsKeyDown(rl.KeyE)))
		camera.Rotation += rotateDirection

		zoomDirection := float32(utils.BoolToInt(rl.IsKeyDown(rl.KeyW))) - float32(utils.BoolToInt(rl.IsKeyDown(rl.KeyS)))
		camera.Zoom += zoomDirection * dt * 2
		camera.Zoom = float32(math.Max(0.24, math.Min(2, float64(camera.Zoom))))
		camera.Target = pos

		// Drawing
		rl.BeginDrawing()
		rl.BeginMode2D(camera)
		rl.ClearBackground(rl.DarkGray)
		for i := 0; i < len(circles); i++ {
			rl.DrawCircleV(circles[i].Position, float32(circles[i].Radius), circles[i].Color)
		}
		rl.DrawCircleV(pos, float32(radius), rl.Black)
		rl.EndMode2D()
		rl.EndDrawing()
	}
}
