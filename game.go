package anyway

import (
	"errors"
	"image/color"

	"github.com/hajimehoshi/bitmapfont/v4"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

const margin = 10.0

// Config controls the window. Zero values fall back to defaults.
type Config struct {
	Title  string
	Width  int
	Height int
}

func (c Config) withDefaults() Config {
	if c.Title == "" {
		c.Title = "anyway"
	}
	if c.Width == 0 {
		c.Width = 640
	}
	if c.Height == 0 {
		c.Height = 480
	}
	return c
}

// Run opens a window and plays the scene until it ends or the window closes.
func Run(scene *Scene, cfg Config) error {
	if scene == nil {
		return errors.New("anyway: nil scene")
	}
	cfg = cfg.withDefaults()

	face := text.NewGoXFace(bitmapfont.Face)
	m := face.Metrics()
	g := &game{
		face:       face,
		lineHeight: m.HAscent + m.HDescent + m.HLineGap,
		width:      cfg.Width,
		height:     cfg.Height,
		p: player{
			scene: scene,
			width: float64(cfg.Width) - 2*margin,
			measure: func(s string) float64 {
				w, _ := text.Measure(s, face, 0)
				return w
			},
		},
	}
	g.p.advance()

	ebiten.SetWindowSize(cfg.Width, cfg.Height)
	ebiten.SetWindowTitle(cfg.Title)
	if err := ebiten.RunGame(g); err != nil && !errors.Is(err, ebiten.Termination) {
		return err
	}
	return nil
}

type game struct {
	p          player
	face       text.Face
	lineHeight float64
	width      int
	height     int
}

func (g *game) Update() error {
	g.p.box.tick()

	pressed := inpututil.IsKeyJustPressed(ebiten.KeyEnter) ||
		inpututil.IsKeyJustPressed(ebiten.KeySpace) ||
		inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft)
	if pressed && g.p.press() {
		return ebiten.Termination
	}
	return nil
}

func (g *game) Draw(screen *ebiten.Image) {
	op := &text.DrawOptions{}
	op.ColorScale.ScaleWithColor(color.White)

	y := margin
	for _, line := range g.p.box.visible() {
		op.GeoM.Reset()
		op.GeoM.Translate(margin, y)
		text.Draw(screen, line, g.face, op)
		y += g.lineHeight
	}
}

func (g *game) Layout(_, _ int) (int, int) { return g.width, g.height }
