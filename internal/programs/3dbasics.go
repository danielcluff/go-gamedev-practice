package program

import (
	r "github.com/gen2brain/raylib-go/raylib"
)

func Basics() {
	r.InitWindow(1200, 900, "3D base")
	defer r.CloseWindow()

	camera := r.Camera3D{}
	camera.Position = r.Vector3{X: 0, Y: 5, Z: 5}
	camera.Up.Y = 1
	camera.Fovy = 45
	camera.Projection = r.CameraPerspective

	// models
	mesh := r.GenMeshCube(1, 1, 1)
	model := r.LoadModelFromMesh(mesh)

	meshCyl := r.GenMeshCylinder(1, 2, 50)
	modelCyl := r.LoadModelFromMesh(meshCyl)

	image := r.GenImageGradientLinear(20, 20, 1, r.Lime, r.SkyBlue)
	texture := r.LoadTextureFromImage(image)
	r.SetMaterialTexture(modelCyl.Materials, 0, texture)

	// move
	pos := r.Vector3{}
	rotation := float32(0)

	for !r.WindowShouldClose() {
		dt := r.GetFrameTime()
		pos.X += 2 * dt
		rotation += 4 * dt
		modelCyl.Transform = r.MatrixRotateX(rotation)

		r.ClearBackground(r.DarkGray)
		r.BeginDrawing()
		r.BeginMode3D(camera)
		r.DrawGrid(10, 1)

		r.DrawModel(model, r.Vector3{}, 1, r.Red)
		r.DrawModel(modelCyl, pos, 1, r.White)
		r.DrawLine3D(r.Vector3{X: -4, Y: 0, Z: -2}, r.Vector3{X: 5, Y: 2, Z: 3}, r.Black)

		r.EndMode3D()
		r.EndDrawing()
	}
}
