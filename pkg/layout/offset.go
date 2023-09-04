package layout

import (
	"github.com/octohelm/ebiten-compose/pkg/layout/direction"
	"github.com/octohelm/ebiten-compose/pkg/patcher"
	"github.com/octohelm/ebiten-compose/pkg/unit"
)

type OffsetSetter interface {
	SetOffset(dp unit.Dp, directions ...direction.Direction) bool
}

type Offset struct {
	X unit.Dp
	Y unit.Dp
}

func (c *Offset) IsZero() bool {
	return c.X == 0 && c.Y == 0
}

func (c *Offset) SetOffset(dp unit.Dp, directions ...direction.Direction) bool {
	if len(directions) == 0 {
		return false
	}

	bp := patcher.NewBatchPatcher()

	for _, p := range directions {
		switch p {
		case direction.X:
			bp.DoPatch(patcher.WhenChanged(&c.X, dp))
		case direction.Y:
			bp.DoPatch(patcher.WhenChanged(&c.Y, dp))
		}
	}

	return bp.Patch()
}
