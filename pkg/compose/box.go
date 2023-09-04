package compose

import (
	"github.com/octohelm/ebiten-compose/pkg/canvas"
	"github.com/octohelm/ebiten-compose/pkg/layout"
	"github.com/octohelm/ebiten-compose/pkg/layout/alignment"
	"github.com/octohelm/ebiten-compose/pkg/node"
)

func Box(modifiers ...Modifier) VNode {
	return H(
		&box{},
		modifiers...,
	)
}

var _ Element = &box{}

type box struct {
	node.Element

	*container

	layout.Aligner
}

func (c box) New() Element {
	return &box{
		Element: node.Element{
			Name: "Box",
		},
		container: newContainer(),
	}
}

func (c *box) Build(ctx BuildContext) VNode {
	return nil
}

func (c *box) Layout(ptx canvas.Canvas) canvas.Dimensions {
	return c.container.Wrap(ElementPainterFunc(func(parent canvas.Canvas) canvas.Dimensions {
		pt := parent.Size()
		w, h := pt.X, pt.Y

		for child := c.FirstChild(); child != nil; child = child.NextSibling() {
			if e, ok := child.(Element); ok {
				cc := canvas.NewCanvas(parent, w, h)

				childSize := e.Layout(cc).Size

				align := alignment.Center

				if alignGetter, ok := e.(layout.AlignGetter); ok {
					align = alignGetter.Align()
				}

				cw, ch := childSize.X, childSize.Y

				switch align {
				case alignment.TopStart:
					cc.Translate(0, 0)
				case alignment.Top:
					cc.Translate((w-cw)/2, 0)
				case alignment.TopEnd:
					cc.Translate(w-cw, 0)
				case alignment.Start:
					cc.Translate(0, (h-ch)/2)
				case alignment.Center:
					cc.Translate((w-cw)/2, (h-ch)/2)
				case alignment.End:
					cc.Translate(w-cw, (h-ch)/2)
				case alignment.BottomStart:
					cc.Translate(0, h-cw)
				case alignment.Bottom:
					cc.Translate((w-cw)/2, h-ch)
				case alignment.BottomEnd:
					cc.Translate(w-cw, h-ch)
				}
			}
		}

		return canvas.Dimensions{Size: pt}
	})).Layout(ptx)
}
