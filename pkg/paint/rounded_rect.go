package paint

import (
	"math"

	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/octohelm/ebiten-compose/pkg/canvas"
)

func NewRoundedRect() *RoundedRect {
	return &RoundedRect{
		Rect: NewRect(),
	}
}

type RoundedRect struct {
	*Rect
	CornerRadiusValues
}

func (rr *RoundedRect) Path(c canvas.Canvas) *vector.Path {
	if rr.TopLeft == 0 && rr.TopRight == 0 && rr.BottomLeft == 0 && rr.BottomRight == 0 {
		return rr.Rect.Path(c)
	}

	offset := c.Offset()

	size := rr.Rectangle(c).Size()

	const q = 4 * (math.Sqrt2 - 1) / 3
	const iq = 1 - q

	x, y := float32(size.X), float32(size.Y)

	bl, br, tl, tr := float32(c.Dp(rr.BottomLeft)), float32(c.Dp(rr.BottomRight)), float32(c.Dp(rr.TopLeft)), float32(c.Dp(rr.TopRight))

	l, t, r, b := float32(offset.X), float32(offset.Y), float32(offset.X)+x, float32(offset.Y)+y

	p := new(vector.Path)

	p.MoveTo(l+tl, t)
	p.LineTo(r-tr, t)
	p.CubicTo(r-tr*iq, t, r, t+tr*iq, r, t+tr)
	p.LineTo(r, b-bl)
	p.CubicTo(r, b-bl*iq, r-bl*iq, b, r-bl, b)
	p.LineTo(l+br, b)
	p.CubicTo(l+br*iq, b, l, b-br*iq, l, b-br)
	p.LineTo(l, t+tl)
	p.CubicTo(l, t+tl*iq, l+tl*iq, t, l+tl, t)

	p.Close()

	return p
}
