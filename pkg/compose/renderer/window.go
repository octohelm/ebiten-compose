package renderer

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/octohelm/ebiten-compose/pkg/canvas"
	"github.com/octohelm/ebiten-compose/pkg/compose/internal"
)

type window struct {
	deviceScaleFactor float64
	dirty             bool
	element           internal.Element
}

func newWindow() *window {
	ebiten.SetScreenClearedEveryFrame(false)

	return &window{
		deviceScaleFactor: ebiten.DeviceScaleFactor(),
		dirty:             true,
	}
}

func (w *window) Update() error {
	return nil
}

func (w *window) Draw(screen *ebiten.Image) {
	if !w.dirty {
		return
	}

	screen.Clear()

	root := canvas.NewRootCanvas(screen, float32(w.deviceScaleFactor))

	_ = w.element.Layout(root)

	root.Draw()

	w.dirty = false
}

func (w *window) Layout(outsideWidth, outsideHeight int) (int, int) {
	return int(float64(outsideWidth) * w.deviceScaleFactor), int(float64(outsideHeight) * w.deviceScaleFactor)
}
