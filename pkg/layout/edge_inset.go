package layout

import (
	"github.com/octohelm/ebiten-compose/pkg/canvas"
	"github.com/octohelm/ebiten-compose/pkg/layout/position"
	"github.com/octohelm/ebiten-compose/pkg/patcher"
	"github.com/octohelm/ebiten-compose/pkg/unit"
)

type EdgeInsetSetter interface {
	SetEdgeInset(dp unit.Dp, positions ...position.Position) bool
}

type EdgeInset struct {
	Top, Bottom, Left, Right unit.Dp
}

func (c *EdgeInset) SetEdgeInset(dp unit.Dp, positions ...position.Position) bool {
	if len(positions) == 0 {
		return false
	}

	bp := patcher.NewBatchPatcher()

	for _, p := range positions {
		switch p {
		case position.Right:
			bp.DoPatch(patcher.WhenChanged(&c.Right, dp))
		case position.Top:
			bp.DoPatch(patcher.WhenChanged(&c.Top, dp))
		case position.Bottom:
			bp.DoPatch(patcher.WhenChanged(&c.Bottom, dp))
		case position.Left:
			bp.DoPatch(patcher.WhenChanged(&c.Left, dp))
		}
	}

	return bp.Patch()
}

func (c *EdgeInset) IsZero() bool {
	return c.Top == 0 && c.Right == 0 && c.Bottom == 0 && c.Left == 0
}

func (c *EdgeInset) NewContext(parent canvas.Canvas) canvas.Canvas {
	size := parent.Size()

	if size.X == 0 || size.Y == 0 {
		return parent
	}

	cc := canvas.NewCanvas(
		parent,
		size.X-(parent.Dp(c.Left)+parent.Dp(c.Right)),
		size.Y-(parent.Dp(c.Top)+parent.Dp(c.Bottom)),
	)

	cc.Translate(
		parent.Dp(c.Left),
		parent.Dp(c.Top),
	)

	return cc
}
