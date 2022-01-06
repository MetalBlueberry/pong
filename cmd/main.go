package main

import (
	"image"
	"log"

	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/metalblueberry/pong/pkg/pong"
)

func init() {
}

func main() {
	var err error
	img, _, err := ebitenutil.NewImageFromFile("sprites_4.png")
	if err != nil {
		log.Fatal(err)
	}

	ebiten.SetWindowSize(640, 480)
	// ebiten.SetWindowFloating(false)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(pong.NewGame(loadPlayer(img), loadBall(img))); err != nil {
		log.Fatal(err)
	}
}

func loadBall(img *ebiten.Image) *ebiten.Image {

	sx := 100
	frameWidth := 50
	sy := 0
	frameHeight := 50

	return img.SubImage(image.Rect(sx, sy, sx+frameWidth, sy+frameHeight)).(*ebiten.Image)
}
func loadPlayer(img *ebiten.Image) *ebiten.Image {

	sx := 100
	frameWidth := 150
	sy := 100
	frameHeight := 50

	return img.SubImage(image.Rect(sx, sy, sx+frameWidth, sy+frameHeight)).(*ebiten.Image)
}
