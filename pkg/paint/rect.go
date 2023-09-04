package paint

import (
	"image"

	"github.com/octohelm/ebiten-compose/pkg/canvas"
	"github.com/octohelm/ebiten-compose/pkg/paint/size"
	"github.com/octohelm/ebiten-compose/pkg/patcher"

	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/octohelm/ebiten-compose/pkg/unit"
)

func NewRect() *Rect {
	return &Rect{
		Width:  -1,
		Height: -1,
	}
}

type Rect struct {
	Width         unit.Dp
	Height        unit.Dp
	FillMaxWidth  bool
	FillMaxHeight bool
}

func (r *Rect) Sized(sizes ...size.Size) bool {
	sized := true

	for _, s := range sizes {
		switch s {
		case size.Width:
			if r.Width < 0 {
				sized = false
			}
		case size.Height:
			if r.Height < 0 {
				sized = false
			}
		}
	}

	return sized
}

var _ FillMaxSetter = &Rect{}

func (r *Rect) SetFillMax(sizes ...size.Size) bool {
	if len(sizes) == 0 {
		return false
	}

	bp := patcher.NewBatchPatcher()

	for _, s := range sizes {
		switch s {
		case size.Width:
			bp.DoPatch(patcher.WhenChanged(&r.FillMaxWidth, true))
			bp.DoPatch(patcher.WhenChanged(&r.Width, -1))
		case size.Height:
			bp.DoPatch(patcher.WhenChanged(&r.FillMaxHeight, true))
			bp.DoPatch(patcher.WhenChanged(&r.Height, -1))
		}
	}

	return bp.Patch()
}

var _ FillMinSetter = &Rect{}

func (r *Rect) SetFillMin(sizes ...size.Size) bool {
	if len(sizes) == 0 {
		return false
	}

	bp := patcher.NewBatchPatcher()

	for _, s := range sizes {
		switch s {
		case size.Width:
			bp.DoPatch(patcher.WhenChanged(&r.FillMaxWidth, false))
			bp.DoPatch(patcher.WhenChanged(&r.Width, -1))
		case size.Height:
			bp.DoPatch(patcher.WhenChanged(&r.FillMaxHeight, false))
			bp.DoPatch(patcher.WhenChanged(&r.Height, -1))
		}
	}

	return bp.Patch()
}

var _ SizeSetter = &Rect{}
var _ FillMaxSetter = &Rect{}

func (r *Rect) SetSize(dp unit.Dp, sizes ...size.Size) bool {
	if len(sizes) == 0 {
		return false
	}

	bp := patcher.NewBatchPatcher()

	for _, s := range sizes {
		switch s {
		case size.Width:
			bp.DoPatch(patcher.WhenChanged(&r.Width, dp))
		case size.Height:
			bp.DoPatch(patcher.WhenChanged(&r.Height, dp))
		}
	}

	return bp.Patch()
}

func (r *Rect) Rectangle(c canvas.Canvas) image.Rectangle {
	rr := image.Rectangle{}

	if r.Width > -1 {
		rr.Max.X = c.Dp(r.Width)
	} else {
		if r.FillMaxWidth {
			rr.Max.X = c.Size().X
		}
	}

	if r.Height > -1 {
		rr.Max.Y = c.Dp(r.Height)
	} else {
		if r.FillMaxHeight {
			rr.Max.Y = c.Size().Y
		}
	}

	return rr
}

func (r *Rect) Path(c canvas.Canvas) *vector.Path {
	offset := c.Offset()

	rect := r.Rectangle(c)

	l, t := offset.X, offset.Y
	w, h := rect.Dx(), rect.Dy()

	p := new(vector.Path)

	p.MoveTo(float32(l), float32(t))
	p.LineTo(float32(l), float32(t+h))
	p.LineTo(float32(l+w), float32(t+h))
	p.LineTo(float32(l+w), float32(t))
	p.Close()

	return p
}
