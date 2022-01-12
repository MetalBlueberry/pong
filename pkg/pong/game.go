package pong

import (
	"fmt"
	"image/color"
	"log"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	player   *Player
	ball     *Ball
	Width    int
	Height   int
	Finished bool
}

func NewGame(player *ebiten.Image, ball *ebiten.Image) *Game {
	return &Game{
		player: NewPlayer(player),
		ball:   NewBall(ball),
		Width:  620,
		Height: 620,
	}
}

func (g *Game) Update() error {
	if g.Finished {
		return nil
	}
	playerPosition := 20
	mx, _ := ebiten.CursorPosition()
	g.player.X = mx
	g.player.Y = g.Height - playerPosition
	g.ball.Update()

	x, y := g.ball.Position.Apply(0, 0)
	switch {
	case y+float64(g.ball.Radius) > float64(g.Height-playerPosition-g.player.Height/2) && g.ball.Direction() == Down:
		log.Print(x, y)
		min := g.player.X - g.player.Width/2
		max := g.player.X + g.player.Width/2
		log.Print(min, max)
		if int(x) < min || int(x) > max {
			g.Finished = true
			return nil
		}
		angle := float64(int(x)-min) / float64(max-min)
		g.ball.ReflectUpWithAngle(float64(angle)*4*math.Pi/6 + math.Pi/6)
		g.ball.SpeedUp(1.2)
	case y < float64(g.ball.Radius):
		g.ball.ReflectY()
	}
	switch {
	case x+float64(g.ball.Radius) > float64(g.Width):
		g.ball.ReflectX()
	case x < float64(g.ball.Radius):
		g.ball.ReflectX()
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	ebitenutil.DrawRect(screen, 0, 0, float64(g.Width), float64(g.Height), color.RGBA{
		G: 255,
		A: 255,
	})
	ebitenutil.DrawRect(screen, 5, 5, float64(g.Width)-10, float64(g.Height)-10, color.RGBA{
		A: 255,
	})

	g.player.Draw(screen)
	g.ball.Draw(screen)
	// bx, by := g.ball.Position.Apply(0, 0)
	// ebitenutil.DrawRect(screen, bx-float64(g.ball.Radius), by-float64(g.ball.Radius), float64(g.ball.Radius*2), float64(g.ball.Radius*2), color.RGBA{
	// 	R: 255,
	// 	A: 100,
	// })

	mx, my := ebiten.CursorPosition()
	ebitenutil.DebugPrint(screen, fmt.Sprintf("x:%d y:%d", mx, my))
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.Width, g.Height
}
