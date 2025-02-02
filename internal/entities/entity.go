package entities

type Entity interface {
	Draw()
	Update(dt float32)
}

func Render(sprites []Entity) {
	for i := 0; i < len(sprites); i++ {
		sprites[i].Draw()
	}
}

func Update(sprites []Entity, dt float32) {
	for i := 0; i < len(sprites); i++ {
		sprites[i].Update(dt)
	}
}
