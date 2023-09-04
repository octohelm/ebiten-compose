package paint

import (
	"github.com/octohelm/ebiten-compose/pkg/paint/size"
	"github.com/octohelm/ebiten-compose/pkg/unit"
)

type SizeSetter interface {
	SetSize(v unit.Dp, sizes ...size.Size) bool
}

type SizedChecker interface {
	Sized(sizes ...size.Size) bool
}

type FillMaxSetter interface {
	SetFillMax(sizes ...size.Size) bool
}

type FillMinSetter interface {
	SetFillMin(sizes ...size.Size) bool
}
