package program

import (
	"go-gamedev-practice/internal/utils"
	"math/rand"

	r "github.com/gen2brain/raylib-go/raylib"
)

type Timer struct {
	duration  float32
	startTime float64
	active    bool
	repeat    bool
	callback  func()
}

func (t *Timer) Activate() {
	t.active = true
	t.startTime = r.GetTime()
}
func (t *Timer) Deactivate() {
	t.active = false
	t.startTime = 0
	if t.repeat {
		t.Activate()
	}
}
func (t *Timer) Update() {
	if t.active {
		if r.GetTime()-t.startTime >= float64(t.duration) {
			if t.callback != nil && t.startTime > 0 {
				t.callback()
			}
			t.Deactivate()
		}
	}
}
func TimerCreate(duration float32, repeat bool, autostart bool, callback func()) *Timer {
	timer := &Timer{
		duration: duration,
		repeat:   repeat,
		callback: callback,
	}

	if autostart {
		timer.Activate()
	}

	return timer
}

type timers struct {
	color    *Timer
	position *Timer
}
type Sprite struct {
	Rec    r.Rectangle
	Color  r.Color
	Timers timers
}

func (s *Sprite) RandomizeColor() {
	s.Color = RandomChoice([]r.Color{r.Red, r.Green, r.Blue, r.Yellow, r.Orange, r.Magenta, r.Purple})
}
func (s *Sprite) RandomizePosition() {
	s.Rec.X = float32(rand.Intn(1200))
	s.Rec.Y = float32(rand.Intn(900))
}
func (s Sprite) Draw() {
	r.DrawRectangleRec(s.Rec, s.Color)
}
func (s *Sprite) Update() {
	s.Timers.color.Update()
	s.Timers.position.Update()
}
func SpriteCreate(pos, size r.Vector2) *Sprite {
	sprite := &Sprite{
		Rec:   r.Rectangle{X: pos.X, Y: pos.Y, Width: size.X, Height: size.Y},
		Color: r.White,
	}
	sprite.Timers = timers{
		color:    TimerCreate(1.5, true, true, sprite.RandomizeColor),
		position: TimerCreate(4, true, true, sprite.RandomizePosition),
	}
	return sprite
}

func Timers() {
	window := utils.Coords{
		X: 1200,
		Y: 900,
	}
	r.InitWindow(int32(window.X), int32(window.Y), "raylib [core] example - basic window")
	defer r.CloseWindow()

	spboi := SpriteCreate(r.Vector2{X: 400, Y: 250}, r.Vector2{X: 50, Y: 50})

	// colorTimer := TimerCreate(1.5, true, true, spboi.RandomizeColor)
	// positionTimer := TimerCreate(4, true, true, spboi.RandomizePosition)
	r.SetTargetFPS(60)
	for !r.WindowShouldClose() {
		// colorTimer.Update()
		// positionTimer.Update()
		spboi.Update()

		r.BeginDrawing()
		r.DrawRectangleRec(spboi.Rec, spboi.Color)
		r.ClearBackground(r.DarkGray)
		r.EndDrawing()
	}
}
