package pong

import "github.com/hajimehoshi/ebiten/v2"

type Player struct {
	image  *ebiten.Image
	X      int
	Y      int
	Width  int
	Height int
}

func NewPlayer(image *ebiten.Image) *Player {
	return &Player{
		image:  image,
		Width:  150,
		Height: 20,
	}

}
func (p *Player) Draw(screen *ebiten.Image) {
	opt := &ebiten.DrawImageOptions{}
	opt.GeoM.Translate(-float64(150)/2, -float64(50)/2)
	opt.GeoM.Translate(float64(p.X), float64(p.Y))

	screen.DrawImage(p.image, opt)
}
