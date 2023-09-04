package paint

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/hajimehoshi/ebiten/v2/vector"
	"github.com/octohelm/ebiten-compose/pkg/patcher"
	"github.com/octohelm/ebiten-compose/pkg/unit"
)

type BorderStrokeSetter interface {
	SetBorderStroke(width unit.Dp, c color.Color) bool
}

var _ BorderStrokeSetter = &BorderStroke{}

type BorderStroke struct {
	Width unit.Dp
	Fill
}

func (f *BorderStroke) SetBorderStroke(width unit.Dp, c color.Color) bool {
	s := patcher.NewBatchPatcher()

	s.DoPatch(patcher.WhenChanged(&f.Width, width))
	s.DoPatch(patcher.WhenChanged(&f.Color, c))

	return s.Patch()
}

func (f *BorderStroke) DrawTo(img *ebiten.Image, w float32, path *vector.Path) {
	// FIXME supported BorderStroke style like dashed, dotted

	if !f.Transparent() {
		op := &vector.StrokeOptions{}
		op.LineCap = vector.LineCapRound
		op.LineJoin = vector.LineJoinRound
		op.Width = w

		vs, is := path.AppendVerticesAndIndicesForStroke([]ebiten.Vertex{}, []uint16{}, op)

		r, g, b, a := f.Color.RGBA()

		for i := range vs {
			vs[i].SrcX = 1
			vs[i].SrcY = 1
			vs[i].ColorR = float32(r) / 0xffff
			vs[i].ColorG = float32(g) / 0xffff
			vs[i].ColorB = float32(b) / 0xffff
			vs[i].ColorA = float32(a) / 0xffff
		}

		img.DrawTriangles(vs, is, whiteSubImage, &ebiten.DrawTrianglesOptions{
			AntiAlias: false,
		})
	}
}
