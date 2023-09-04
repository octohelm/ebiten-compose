package compose

import (
	"github.com/octohelm/ebiten-compose/pkg/canvas"
	"github.com/octohelm/ebiten-compose/pkg/layout"
	"github.com/octohelm/ebiten-compose/pkg/layout/position"
	"github.com/octohelm/ebiten-compose/pkg/paint"
	"github.com/octohelm/ebiten-compose/pkg/unit"
)

func newContainer() *container {
	return &container{
		shape: &shape{SizedShape: paint.NewRoundedRect()},
	}
}

type container struct {
	*shape

	layout.FlexWeight
	layout.Offset
	layout.EdgeInset
}

func (c *container) Wrap(p ElementPainter) ElementPainter {
	return ElementPainterFunc(func(parent canvas.Canvas) canvas.Dimensions {
		shapePt := c.shape.Rectangle(parent).Size()
		if shapePt.X == 0 || shapePt.Y == 0 {
			return canvas.Dimensions{
				Size: shapePt,
			}
		}

		cc := canvas.NewCanvas(parent, shapePt.X, shapePt.Y)

		if !c.Offset.IsZero() {
			cc.Translate(parent.Dp(c.Offset.X), parent.Dp(c.Offset.Y))
		}

		cc.PushOp(c.shape.DrawOp(cc))

		if !c.EdgeInset.IsZero() {
			cc = c.EdgeInset.NewContext(cc)
		}

		_ = p.Layout(cc)

		return canvas.Dimensions{
			Size: shapePt,
		}
	})
}

type shape struct {
	paint.SizedShape
	paint.Shadow
	paint.Fill
	paint.BorderStroke
}

var _ paint.CornerRadiusSetter = &shape{}

func (s *shape) SetCornerRadius(v unit.Dp, positions ...position.Position) bool {
	if setter, ok := s.SizedShape.(paint.CornerRadiusSetter); ok {
		return setter.SetCornerRadius(v, positions...)
	}
	return false
}

func (s *shape) DrawOp(cc canvas.Canvas) canvas.Op {
	return canvas.OpFunc(func() {
		path := s.Path(cc)

		// paint shadow if exists
		s.Shadow.Paint(cc, s.SizedShape)

		// paint background if exists
		s.Fill.DrawTo(cc.Image(), path)

		// paint then border
		s.BorderStroke.DrawTo(cc.Image(), float32(cc.Dp(s.BorderStroke.Width)), path)
	})

}
