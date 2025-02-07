package program

import (
	"go-gamedev-practice/internal/utils"

	r "github.com/gen2brain/raylib-go/raylib"
)

func GetBoundingBox(model r.Model, pos r.Vector3) r.BoundingBox {
	boundingBox := r.GetMeshBoundingBox(model.GetMeshes()[0])
	minBoundary := r.Vector3Add(pos, boundingBox.Min)
	maxBoundary := r.Vector3Add(pos, boundingBox.Max)
	return r.BoundingBox{Min: minBoundary, Max: maxBoundary}
}

func CheckCollision(axis rune, player *Player, object *Object) {
	playerBbox := GetBoundingBox(player.model, player.pos)
	objectBbox := GetBoundingBox(object.model, object.pos)
	if r.CheckCollisionBoxes(playerBbox, objectBbox) {
		if axis == 'X' {
			if player.direction.X > 0 { // moving right
				player.pos.X = objectBbox.Min.X - float32(player.size[0])*0.50001
			} else if player.direction.X < 0 { // moving left
				player.pos.X = objectBbox.Max.X + float32(player.size[0])*0.50001
			}
		} else if axis == 'Y' {
			if player.direction.Y > 0 { // moving up
				player.pos.Y = objectBbox.Min.Y - float32(player.size[1])*0.50001
			} else if player.direction.Y < 0 { // moving down
				player.pos.Y = objectBbox.Max.Y + float32(player.size[1])*0.50001
			}
		} else if axis == 'Z' {
			if player.direction.Z > 0 { // moving forward
				player.pos.Z = objectBbox.Min.Z - float32(player.size[2])*0.50001
			} else if player.direction.Z < 0 { // moving backward
				player.pos.Z = objectBbox.Max.Z + float32(player.size[2])*0.50001
			}
		}
	}
}

type Object struct {
	model r.Model
	pos   r.Vector3
	size  [3]int
}
type Player struct {
	*Object
	direction r.Vector3
	speed     float32
}

func ObjectCreate(x, y, z int) *Object {
	object := &Object{
		model: r.LoadModelFromMesh(r.GenMeshCube(float32(x), float32(y), float32(z))),
		pos:   r.Vector3{},
		size:  [3]int{x, y, z},
	}
	return object
}
func PlayerCreate(x, y, z int) *Player {
	player := &Player{
		Object: &Object{
			model: r.LoadModelFromMesh(r.GenMeshCube(float32(x), float32(y), float32(z))),
			pos:   r.Vector3{},
			size:  [3]int{x, y, z},
		},
		direction: r.Vector3{},
		speed:     10,
	}
	return player
}

func ThreeDCollisions() {
	r.InitWindow(1200, 900, "3D base")
	defer r.CloseWindow()

	camera := r.Camera3D{}
	camera.Position = r.Vector3{X: 0, Y: 10, Z: 10}
	camera.Up.Y = 1
	camera.Fovy = 45
	camera.Projection = r.CameraPerspective

	// models

	player := PlayerCreate(1, 1, 1)

	obstacle := ObjectCreate(2, 6, 6)
	obstacle.pos.X = 3

	for !r.WindowShouldClose() {
		// input
		player.direction.X = float32(utils.BoolToInt(r.IsKeyDown(r.KeyRight))) - float32(utils.BoolToInt(r.IsKeyDown(r.KeyLeft)))
		player.direction.Z = float32(utils.BoolToInt(r.IsKeyDown(r.KeyDown))) - float32(utils.BoolToInt(r.IsKeyDown(r.KeyUp)))
		player.direction.Y = float32(utils.BoolToInt(r.IsKeyDown(r.KeyW))) - float32(utils.BoolToInt(r.IsKeyDown(r.KeyS)))

		// movement collisions
		dt := r.GetFrameTime()
		player.pos.X += player.direction.X * player.speed * dt
		CheckCollision('X', player, obstacle)
		player.pos.Z += player.direction.Z * player.speed * dt
		CheckCollision('Z', player, obstacle)
		player.pos.Y += player.direction.Y * player.speed * dt
		CheckCollision('Y', player, obstacle)

		r.ClearBackground(r.DarkGray)
		r.BeginDrawing()
		r.BeginMode3D(camera)
		r.DrawGrid(10, 1)

		r.DrawModel(player.model, player.pos, 1, r.Red)
		r.DrawModel(obstacle.model, obstacle.pos, 1, r.White)
		r.DrawBoundingBox(GetBoundingBox(player.model, player.pos), r.Green)
		r.DrawBoundingBox(GetBoundingBox(obstacle.model, obstacle.pos), r.Green)

		r.EndMode3D()
		r.EndDrawing()
	}
}
