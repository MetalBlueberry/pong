package pong

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2"
)

type Ball struct {
	image *ebiten.Image
	// X        int
	// Y        int
	Radius   int
	Position ebiten.GeoM
	Speed    ebiten.GeoM
}

func NewBall(image *ebiten.Image) *Ball {
	b := &Ball{
		image:  image,
		Radius: 12,
	}
	b.Position.Translate(100, 300)
	b.Speed.Translate(1, 3)
	return b
}

func (b *Ball) ImageRadius() float64 {
	return float64(b.image.Bounds().Dx() / 2)
}

func (b *Ball) Draw(screen *ebiten.Image) {
	opt := &ebiten.DrawImageOptions{}
	opt.GeoM.Translate(-float64(b.ImageRadius()), -float64(b.ImageRadius()))
	opt.GeoM.Concat(b.Position)

	screen.DrawImage(b.image, opt)
}

func (b *Ball) Update() {
	s := b.Speed
	b.Position.Translate(s.Apply(0, 0))
}

func (b *Ball) ReflectX() {
	vx, vy := b.Speed.Apply(0, 0)
	b.Speed.Reset()
	b.Speed.Translate(-vx, vy)
}

func (b *Ball) ReflectY() {
	vx, vy := b.Speed.Apply(0, 0)
	b.Speed.Reset()
	b.Speed.Translate(vx, -vy)
}

func (b *Ball) ReflectUpWithAngle(angle float64) {
	vx, vy := b.Speed.Apply(0, 0)
	b.Speed.Reset()
	b.Speed.Translate(math.Sqrt(vx*vx+vy*vy), 0)
	// b.Speed.Translate(0, -4)
	b.Speed.Rotate(-angle)
}

func (b *Ball) SpeedUp(multiplier float64) {
	b.Speed.Scale(multiplier, multiplier)
}

type Direction int

const (
	None Direction = iota
	Up
	Down
)

func (b *Ball) Direction() Direction {
	_, vy := b.Speed.Apply(0, 0)
	if vy > 0 {
		return Down
	}
	return Up
}
