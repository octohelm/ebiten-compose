package text

import (
	"image/color"

	"github.com/octohelm/ebiten-compose/pkg/unit"
)

type Alignment int

const (
	Left Alignment = iota
	Right
	Center
)

type FontStyle int

const (
	Normal FontStyle = iota
	Italy
)

type Style struct {
	FontFamily string
	FontStyle  FontStyle
	FontWeight int
	FontSize   unit.Dp
	LineHeight unit.Dp
	Color      color.Color
}

func (s Style) Merge(style Style) *Style {
	return &s
}
