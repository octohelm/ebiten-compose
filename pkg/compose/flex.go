package compose

import (
	"image"
	"math"

	"github.com/octohelm/ebiten-compose/pkg/layout/alignment"
	"github.com/octohelm/ebiten-compose/pkg/layout/arrangement"

	"github.com/octohelm/ebiten-compose/pkg/canvas"
	"github.com/octohelm/ebiten-compose/pkg/compose/internal"
	"github.com/octohelm/ebiten-compose/pkg/layout"
	"github.com/octohelm/ebiten-compose/pkg/layout/direction"
	"github.com/octohelm/ebiten-compose/pkg/node"
	"github.com/octohelm/ebiten-compose/pkg/paint"
	"github.com/octohelm/ebiten-compose/pkg/paint/size"
)

func Column(modifiers ...Modifier) VNode {
	return H(&flex{
		Element: node.Element{
			Name: "Column",
		},
		Axis: direction.Vertical,
	}, modifiers...)
}

func Row(modifiers ...Modifier) VNode {
	return H(&flex{
		Element: node.Element{
			Name: "Row",
		},
		Axis: direction.Horizontal,
	}, modifiers...)
}

var _ Element = &flex{}

type flex struct {
	node.Element

	*container

	Axis direction.Axis

	layout.Spacer
	layout.Aligner
	layout.Arrangementer

	layout.Scrollable
}

func (f *flex) Build(context internal.BuildContext) internal.VNode {
	return nil
}

func (f *flex) New() Element {
	return &flex{
		Element: node.Element{
			Name: f.Name,
		},
		Axis:      f.Axis,
		container: newContainer(),
	}
}

func (f *flex) Layout(parent canvas.Canvas) canvas.Dimensions {
	return f.container.Wrap(ElementPainterFunc(func(parentCc canvas.Canvas) canvas.Dimensions {
		if f.Scrollable.Enabled {
		}

		children := make([]*flexChild, 0)

		idx := 0
		addFlexChild := func(child *flexChild) {
			if idx > 0 && f.Spacing != 0 {
				children = append(children, sized(ElementPainterFunc(func(cc canvas.Canvas) canvas.Dimensions {
					if f.Axis == direction.Vertical {
						return canvas.Dimensions{
							Size: image.Pt(0, parentCc.Dp(f.Spacing)),
						}
					}
					return canvas.Dimensions{
						Size: image.Pt(parentCc.Dp(f.Spacing), 0),
					}
				})))
			}

			children = append(children, child)
			idx++
		}

		totalWeight := float32(0)

		for child := f.FirstChild(); child != nil; child = child.NextSibling() {
			if elem, ok := child.(Element); ok {
				if sc, ok := elem.(paint.SizedChecker); ok {
					switch f.Axis {
					case direction.Vertical:
						if sc.Sized(size.Height) {
							addFlexChild(sized(elem))
							continue
						}
					case direction.Horizontal:
						if sc.Sized(size.Width) {
							addFlexChild(sized(elem))
							continue
						}
					}
				}

				weight := float32(1)

				if getter, ok := elem.(layout.WeightGetter); ok {
					if v, ok := getter.Weight(); ok {
						weight = v
					}
				}

				totalWeight += weight

				if sc, ok := elem.(paint.FillMaxSetter); ok {
					switch f.Axis {
					case direction.Horizontal:
						sc.SetFillMax(size.Width)
					case direction.Vertical:
						sc.SetFillMax(size.Height)
					}
				}

				addFlexChild(flexed(weight, elem))
			}
		}

		pt := parentCc.Size()

		if len(children) == 0 {
			return canvas.Dimensions{
				Size: pt,
			}
		}

		exactlySize := 0

		for i := range children {
			c := children[i]

			if !c.flexed() {
				c.Draw(parentCc)

				switch f.Axis {
				case direction.Horizontal:
					exactlySize += c.dimensions.Size.X
				case direction.Vertical:
					exactlySize += c.dimensions.Size.Y
				}
			}
		}

		remainSize := 0

		switch f.Axis {
		case direction.Horizontal:
			remainSize = pt.X
		case direction.Vertical:
			remainSize = pt.Y
		}

		remainSize -= exactlySize

		offset := 0

		for i := range children {
			c := children[i]

			if c.flexed() {
				b := parent.Size()

				flexedSize := int(math.Round(float64(float32(remainSize) * c.weight / totalWeight)))

				switch f.Axis {
				case direction.Horizontal:
					b.X = flexedSize
				case direction.Vertical:
					b.Y = flexedSize
				}

				c.Draw(canvas.NewCanvas(parentCc, b.X, b.Y))
			}

			switch f.Axis {
			case direction.Horizontal:
				c.canvas.Translate(offset, 0)
				offset += c.dimensions.Size.X
			case direction.Vertical:
				c.canvas.Translate(0, offset)
				offset += c.dimensions.Size.Y
			}

			c.canvas.Translate(f.offsetOfAlignment(pt, c.dimensions.Size))
			c.canvas.Translate(f.offsetOfArrangement(remainSize, len(children), i))
		}

		return canvas.Dimensions{
			Size: pt,
		}
	})).Layout(parent)
}

func (f *flex) offsetOfAlignment(parent image.Point, child image.Point) (int, int) {
	switch f.Alignment {
	case alignment.End:
		if f.Axis == direction.Horizontal {
			return 0, parent.Y - child.Y
		}
		return parent.X - child.X, 0
	case alignment.Center:
		if f.Axis == direction.Horizontal {
			return 0, (parent.Y - child.Y) / 2
		}
		return (parent.X - child.X) / 2, 0
	}
	return 0, 0
}

func (f *flex) offsetOfArrangement(spacing int, n int, idx int) (int, int) {
	if spacing <= 0 {
		return 0, 0
	}

	switch f.Arrangement {

	case arrangement.SpaceEvenly:
		if f.Axis == direction.Horizontal {
			return spacing / (n + 1) * (idx + 1), 0
		}
		return 0, spacing / (n + 1) * (idx + 1)
	case arrangement.SpaceAround:
		if f.Axis == direction.Horizontal {
			return spacing / (n * 2) * (idx*2 + 1), 0
		}
		return 0, spacing / (n * 2) * (idx*2 + 1)
	case arrangement.SpaceBetween:
		if (n - 1) == 0 {
			return 0, 0
		}
		if f.Axis == direction.Horizontal {
			return spacing / (n - 1) * idx, 0
		}
		return 0, spacing / (n - 1) * idx
	case arrangement.Start:
		if f.Axis == direction.Horizontal {
			return 0, 0
		}
		return 0, 0
	case arrangement.End:
		if f.Axis == direction.Horizontal {
			return spacing, 0
		}
		return 0, spacing
	case arrangement.Center:
		if f.Axis == direction.Horizontal {
			return spacing / 2, 0
		}
		return 0, spacing / 2
	}
	return 0, 0
}

func sized(e ElementPainter) *flexChild {
	return &flexChild{element: e}
}

func flexed(weight float32, e ElementPainter) *flexChild {
	return &flexChild{weight: weight, element: e}
}

type flexChild struct {
	weight  float32
	element ElementPainter

	dimensions canvas.Dimensions
	canvas     canvas.Canvas
}

func (c *flexChild) flexed() bool {
	return c.weight > 0
}

func (c *flexChild) Draw(cc canvas.Canvas) {
	s := cc.Size()

	c.canvas = canvas.NewCanvas(cc, s.X, s.Y)
	c.dimensions = c.element.Layout(c.canvas)
}
