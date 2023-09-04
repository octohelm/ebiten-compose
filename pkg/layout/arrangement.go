package layout

import (
	"github.com/octohelm/ebiten-compose/pkg/layout/arrangement"
	"github.com/octohelm/ebiten-compose/pkg/patcher"
)

type Arrangement = arrangement.Arrangement

type ArrangementSetter interface {
	SetArrangement(Arrangement Arrangement) bool
}

type Arrangementer struct {
	Arrangement Arrangement
}

var _ ArrangementSetter = &Arrangementer{}

func (a *Arrangementer) SetArrangement(arrangement Arrangement) bool {
	return patcher.WhenChanged(&a.Arrangement, arrangement).Patch()
}
