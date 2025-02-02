package entities

import (
	h "go-gamedev-practice/internal/utils"
	"path/filepath"

	r "github.com/gen2brain/raylib-go/raylib"
)

type Player struct {
	Texture r.Texture2D
	Sprite
}

func (p Player) Draw() {
	r.DrawTextureV(p.Texture, p.Position, r.White)
}
func (p *Player) Update(dt float32) {
	// ship input
	p.Direction.X = float32(h.BoolToInt(r.IsKeyDown(r.KeyRight))) - float32(h.BoolToInt(r.IsKeyDown(r.KeyLeft)))
	p.Direction.Y = float32(h.BoolToInt(r.IsKeyDown(r.KeyDown))) - float32(h.BoolToInt(r.IsKeyDown(r.KeyUp)))
	p.Direction = r.Vector2Normalize(p.Direction)

	p.Position.X += p.Direction.X * p.Speed * dt
	p.Position.Y += p.Direction.Y * p.Speed * dt
}
func PlayerCreate(x float32, y float32, speed float32) Player {
	player := Player{
		Texture: r.LoadTexture(filepath.Join("..", "..", "internal", "assets", "spaceship.png")),
		Sprite: Sprite{
			Position: r.Vector2{X: x, Y: y},
			Speed:    speed,
		},
	}
	return player
}
