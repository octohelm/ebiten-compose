package shader

import (
	_ "embed"

	"github.com/hajimehoshi/ebiten/v2"
)

//go:embed blur.kage.go
var blurKage []byte

var Blur *ebiten.Shader

func init() {
	s, err := ebiten.NewShader(blurKage)
	if err != nil {
		panic(err)
	}
	Blur = s
}
