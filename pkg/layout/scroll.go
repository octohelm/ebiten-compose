package layout

import (
	"github.com/octohelm/ebiten-compose/pkg/patcher"
)

type Axis int

const (
	Horizontal Axis = iota
	Vertical
)

type ScrollableSetter interface {
	SetScrollable(axis Axis, enabled bool) bool
}

type Scrollable struct {
	Enabled bool
	Axis    Axis
}

func (s *Scrollable) SetScrollable(axis Axis, enabled bool) bool {
	bp := patcher.NewBatchPatcher()
	bp.DoPatch(patcher.WhenChanged(&s.Axis, axis))
	bp.DoPatch(patcher.WhenChanged(&s.Enabled, enabled))
	return bp.Patch()
}
