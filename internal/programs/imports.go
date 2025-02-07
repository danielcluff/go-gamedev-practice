package program

import (
	"path/filepath"

	r "github.com/gen2brain/raylib-go/raylib"
)

func Imports() {
	r.InitWindow(1200, 900, "3D base")
	defer r.CloseWindow()

	camera := r.Camera3D{}
	camera.Position = r.Vector3{X: 0, Y: 5, Z: 5}
	camera.Up.Y = 1
	camera.Fovy = 45
	camera.Projection = r.CameraPerspective

	// models
	ship := r.LoadModel(filepath.Join("..", "..", "internal", "assets", "models", "ship.glb"))
	rupee := r.LoadModel(filepath.Join("..", "..", "internal", "assets", "models", "rupee.gltf"))

	for !r.WindowShouldClose() {
		ship.Transform = r.MatrixScale(1, 2, 1)
		r.BeginDrawing()
		r.ClearBackground(r.DarkGray)
		r.BeginMode3D(camera)
		r.DrawGrid(10, 1)

		r.DrawModel(ship, r.Vector3{}, 1, r.White)
		r.DrawModel(rupee, r.Vector3{X: 1}, 1, r.White)

		r.EndMode3D()
		r.EndDrawing()
	}
}
