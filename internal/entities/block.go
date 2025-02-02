package entities

import (
	r "github.com/gen2brain/raylib-go/raylib"
)

type Block struct {
	Size r.Vector2
	Sprite
}

func (e Block) Draw() {
	r.DrawRectangleV(e.Position, e.Size, r.White)
}
func (e *Block) Update(dt float32) {
	e.Position.X += e.Direction.X * e.Speed * dt
	e.Position.Y += e.Direction.Y * e.Speed * dt
}
func BlockCreate(x float32, y float32, speed float32) Block {
	block := Block{
		Size: r.Vector2{X: 50, Y: 50},
		Sprite: Sprite{
			Position:  r.Vector2{X: x, Y: y},
			Direction: r.Vector2{X: 1, Y: 0},
			Speed:     speed,
		},
	}
	return block
}
