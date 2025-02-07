package program

import (
	"path/filepath"

	r "github.com/gen2brain/raylib-go/raylib"
)

func Shaders() {
	r.InitWindow(1200, 900, "Shader")
	defer r.CloseWindow()

	camera := r.Camera3D{}
	camera.Position = r.Vector3{X: 0, Y: 5, Z: 5}
	camera.Up.Y = 1
	camera.Fovy = 45
	camera.Projection = r.CameraPerspective

	// models
	model := r.LoadModelFromMesh(r.GenMeshCylinder(1, 2, 5))
	texture := r.LoadTextureFromImage(r.GenImageGradientLinear(100, 100, 1, r.Red, r.White))
	r.SetMaterialTexture(&model.GetMaterials()[0], r.MapAlbedo, texture)

	// shader := r.LoadShader("", filepath.Join("..", "..", "internal", "assets", "shaders", "grayscale.fs"))
	shader := r.LoadShader("", filepath.Join("..", "..", "internal", "assets", "shaders", "flash.fs"))
	model.Materials.Shader = shader
	flashLoc := r.GetShaderLocation(shader, "flash")
	flash := []float32{1, 0}
	flashNot := []float32{0, 0}

	// move
	// pos := r.Vector3{}
	// rotation := float32(0)

	for !r.WindowShouldClose() {
		// dt := r.GetFrameTime()
		if r.IsKeyPressed(r.KeyA) {
			r.SetShaderValue(shader, flashLoc, flash, r.ShaderUniformVec2)
		}
		if r.IsKeyReleased(r.KeyA) {
			r.SetShaderValue(shader, flashLoc, flashNot, r.ShaderUniformVec2)
		}

		r.ClearBackground(r.DarkGray)
		r.BeginDrawing()
		r.BeginMode3D(camera)

		r.DrawModel(model, r.Vector3{}, 1, r.White)

		r.EndMode3D()
		r.EndDrawing()
	}
}
