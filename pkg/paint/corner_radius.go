package paint

import (
	"github.com/octohelm/ebiten-compose/pkg/layout/position"
	"github.com/octohelm/ebiten-compose/pkg/patcher"
	"github.com/octohelm/ebiten-compose/pkg/unit"
)

type CornerRadiusSetter interface {
	SetCornerRadius(v unit.Dp, positions ...position.Position) bool
}

type CornerRadiusGetter interface {
	CornerRadius() CornerRadiusValues
}

type CornerRadiusValues struct {
	TopLeft     unit.Dp
	TopRight    unit.Dp
	BottomLeft  unit.Dp
	BottomRight unit.Dp
}

func (r CornerRadiusValues) CornerRadius() CornerRadiusValues {
	return r
}

func (r *CornerRadiusValues) SetCornerRadius(dp unit.Dp, positions ...position.Position) bool {
	s := patcher.NewBatchPatcher()

	for _, p := range positions {
		switch p {
		case position.TopLeft:
			s.DoPatch(patcher.WhenChanged(&r.TopLeft, dp))
		case position.BottomLeft:
			s.DoPatch(patcher.WhenChanged(&r.BottomLeft, dp))
		case position.TopRight:
			s.DoPatch(patcher.WhenChanged(&r.TopRight, dp))
		case position.BottomRight:
			s.DoPatch(patcher.WhenChanged(&r.BottomRight, dp))
		case position.Top:
			s.DoPatch(patcher.WhenChanged(&r.TopLeft, dp))
			s.DoPatch(patcher.WhenChanged(&r.TopRight, dp))
		case position.Bottom:
			s.DoPatch(patcher.WhenChanged(&r.BottomLeft, dp))
			s.DoPatch(patcher.WhenChanged(&r.BottomRight, dp))
		case position.Left:
			s.DoPatch(patcher.WhenChanged(&r.TopLeft, dp))
			s.DoPatch(patcher.WhenChanged(&r.BottomLeft, dp))
		case position.Right:
			s.DoPatch(patcher.WhenChanged(&r.TopRight, dp))
			s.DoPatch(patcher.WhenChanged(&r.BottomRight, dp))
		}
	}

	return s.Patch()
}
