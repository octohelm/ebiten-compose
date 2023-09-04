package paint

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/octohelm/ebiten-compose/pkg/patcher"

	"github.com/hajimehoshi/ebiten/v2/vector"
)

type FillSetter interface {
	SetFill(c color.Color) bool
}

var _ FillSetter = &Fill{}

type Fill struct {
	Color color.Color
}

func (f *Fill) Transparent() bool {
	return f.Color == nil
}

func (f *Fill) SetFill(c color.Color) bool {
	return patcher.WhenChanged(&f.Color, c).Patch()
}

func (f *Fill) DrawTo(img *ebiten.Image, path *vector.Path) {
	if !f.Transparent() {
		vs, is := path.AppendVerticesAndIndicesForFilling(nil, nil)

		r, g, b, a := f.Color.RGBA()

		for i := range vs {
			vs[i].SrcX = 1
			vs[i].SrcY = 1
			vs[i].ColorR = float32(r) / 0xffff
			vs[i].ColorG = float32(g) / 0xffff
			vs[i].ColorB = float32(b) / 0xffff
			vs[i].ColorA = float32(a) / 0xffff
		}

		op := &ebiten.DrawTrianglesOptions{}
		op.ColorScaleMode = ebiten.ColorScaleModePremultipliedAlpha
		op.AntiAlias = false

		img.DrawTriangles(vs, is, whiteSubImage, op)
	}
}
