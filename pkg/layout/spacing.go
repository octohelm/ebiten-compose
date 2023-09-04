package layout

import (
	"github.com/octohelm/ebiten-compose/pkg/patcher"
	"github.com/octohelm/ebiten-compose/pkg/unit"
)

type SpacingSetter interface {
	SetSpacing(dp unit.Dp) bool
}

type Spacer struct {
	Spacing unit.Dp
}

func (s *Spacer) SetSpacing(dp unit.Dp) bool {
	return patcher.WhenChanged(&s.Spacing, dp).Patch()
}
