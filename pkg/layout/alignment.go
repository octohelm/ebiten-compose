package layout

import (
	"github.com/octohelm/ebiten-compose/pkg/layout/alignment"
	"github.com/octohelm/ebiten-compose/pkg/patcher"
)

type Alignment = alignment.Alignment

type AlignSetter interface {
	SetAlign(align Alignment) bool
}

type AlignGetter interface {
	Align() Alignment
}

type Aligner struct {
	Alignment Alignment
}

var _ AlignSetter = &Aligner{}

func (a *Aligner) SetAlign(alignment Alignment) bool {
	return patcher.WhenChanged(&a.Alignment, alignment).Patch()
}

func (a *Aligner) Align() Alignment {
	return a.Alignment
}
