package program

import (
	"go-gamedev-practice/internal/utils"

	r "github.com/gen2brain/raylib-go/raylib"
)

func Collisions() {
	window := utils.Coords{
		X: 1920,
		Y: 1080,
	}
	r.InitWindow(int32(window.X), int32(window.Y), "raylib [core] example - basic window")
	defer r.CloseWindow()

	levelMap := []string{
		"1111111111111111111",
		"1010000000000000001",
		"1010000000001111111",
		"1000000000000000111",
		"1000000200000000011",
		"1000000000000100001",
		"1000000000000100001",
		"1001100000000100001",
		"1001100000000100001",
		"1001100000000100001",
		"1111111111111111111",
	}
	player := r.Rectangle{X: 400, Y: 300, Width: 60, Height: 60}
	speed := float32(300)
	direction := r.Vector2{}
	blocks := []r.Rectangle{}
	blockSize := 100
	for rowIndex, row := range levelMap {
		for colIndex, cell := range row {
			if cell == '1' {
				x := colIndex * blockSize
				y := rowIndex * blockSize
				block := r.Rectangle{X: float32(x), Y: float32(y), Width: float32(blockSize), Height: float32(blockSize)}
				blocks = append(blocks, block)
			}
		}
	}

	collision := func(axis rune) {
		for _, block := range blocks {
			if r.CheckCollisionRecs(block, player) {
				if axis == 'x' {
					if direction.X > 0 { // moving right
						player.X = block.X - player.Width
					} else if direction.X < 0 { // moving left
						player.X = block.X + block.Width
					}
				} else {
					if direction.Y > 0 { // moving down
						player.Y = block.Y - player.Height
					} else if direction.Y < 0 { // moving up
						player.Y = block.Y + block.Height
					}
				}
			}
		}
	}

	r.SetTargetFPS(60)
	for !r.WindowShouldClose() {
		direction.X = float32(utils.BoolToInt(r.IsKeyDown(r.KeyRight))) - float32(utils.BoolToInt(r.IsKeyDown(r.KeyLeft)))
		direction.Y = float32(utils.BoolToInt(r.IsKeyDown(r.KeyDown))) - float32(utils.BoolToInt(r.IsKeyDown(r.KeyUp)))

		// movement
		dt := r.GetFrameTime()
		player.X += direction.X * speed * dt
		collision('x')
		player.Y += direction.Y * speed * dt
		collision('y')

		// updates
		r.BeginDrawing()
		r.ClearBackground(r.DarkGray)
		for _, block := range blocks {
			r.DrawRectangleRec(block, r.Gray)
		}
		r.DrawRectangleRec(player, r.Red)
		r.DrawFPS(0, 0)
		r.EndDrawing()
	}
}
