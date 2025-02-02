package entities

import (
	r "github.com/gen2brain/raylib-go/raylib"
)

type Sprite struct {
	Position  r.Vector2
	Direction r.Vector2
	Speed     float32
}

func (e *Sprite) Move(dt float32) {
	e.Position.X = e.Direction.X * e.Speed * dt
	e.Position.Y = e.Direction.X * e.Speed * dt
}
