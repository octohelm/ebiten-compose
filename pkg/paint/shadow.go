package paint

import (
	"image"
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/octohelm/ebiten-compose/pkg/bezier"
	"github.com/octohelm/ebiten-compose/pkg/canvas"
	"github.com/octohelm/ebiten-compose/pkg/paint/shader"
	"github.com/octohelm/ebiten-compose/pkg/patcher"
	"github.com/octohelm/ebiten-compose/pkg/unit"
)

type ShadowSetter interface {
	SetShadow(elevation unit.Dp) bool
}

type Shadow struct {
	elevation unit.Dp

	ambient  Fill
	penumbra Fill
	umbra    Fill
}

func (s *Shadow) SetShadow(elevation unit.Dp) bool {
	return patcher.WhenChanged(&s.elevation, elevation).Patch()
}

func (s *Shadow) Paint(cc canvas.Canvas, shape SizedShape) {
	if s.elevation <= 0 {
		return
	}

	if s.ambient.Transparent() {
		s.ambient.Color = color.NRGBA{A: 0x10}
	}
	if s.umbra.Transparent() {
		s.umbra.Color = color.NRGBA{A: 0x25}
	}
	if s.penumbra.Transparent() {
		s.penumbra.Color = color.NRGBA{A: 0x15}
	}

	s.paint(cc, shape, s.ambient.Color, ambientShadow)
	s.paint(cc, shape, s.umbra.Color, umbraShadow)
	s.paint(cc, shape, s.penumbra.Color, penumbraShadow)
}

func (s *Shadow) paint(parentCanvas canvas.Canvas, shape SizedShape, c color.Color, se *shadowElevation) {
	offsetY, blurRadius, spread := se.Calc(float64(s.elevation))

	blurSpread := parentCanvas.Dp(unit.Dp(blurRadius))

	rr := s.withSpread(parentCanvas, shape, unit.Dp(spread/parentCanvas.Dp(1)))

	w, h := parentCanvas.Dp(rr.Width), parentCanvas.Dp(rr.Height)
	maxW, maxH := w+blurSpread*2, h+blurSpread*2

	container := canvas.NewCanvasWithImage(parentCanvas, ebiten.NewImage(maxW, maxH))
	container.Translate(-blurSpread, -blurSpread)
	defer func() {
		opi := &ebiten.DrawImageOptions{}
		opi.GeoM.Translate(canvas.XYAs[float64](container.Offset()))
		opi.GeoM.Translate(canvas.XYAs[float64](parentCanvas.Offset().Add(image.Pt(0, offsetY))))

		parentCanvas.Image().DrawImage(container.Image(), opi)
	}()

	shadowContainer := canvas.NewCanvasWithImage(container, ebiten.NewImage(maxW, maxH))
	defer func() {
		if blurRadius > 0 {
			ops := &ebiten.DrawRectShaderOptions{}
			ops.Uniforms = map[string]any{
				"Radius": float32(blurRadius),
			}
			ops.Images[0] = shadowContainer.Image()
			bounds := shadowContainer.Image().Bounds()
			container.Image().DrawRectShader(bounds.Dy(), bounds.Dx(), shader.Blur, ops)
		}
	}()

	shadowGraph := canvas.NewCanvasWithImage(container, ebiten.NewImage(w, h))
	(&Fill{Color: c}).DrawTo(shadowGraph.Image(), rr.Path(shadowGraph))
	defer func() {
		ops := &ebiten.DrawImageOptions{}
		ops.GeoM.Translate(float64(blurSpread), float64(blurSpread))
		shadowContainer.Image().DrawImage(shadowGraph.Image(), ops)
	}()
}

func (s *Shadow) withSpread(c canvas.Canvas, shape Shape, spread unit.Dp) *RoundedRect {
	pt := shape.Rectangle(c).Size()

	rrect := &RoundedRect{
		Rect: &Rect{
			Width:  unit.Dp(pt.X/c.Dp(1)) + spread,
			Height: unit.Dp(pt.X/c.Dp(1)) + spread,
		},
	}

	if g, ok := shape.(CornerRadiusGetter); ok {
		cr := g.CornerRadius()

		rrect.TopRight = cr.TopRight + spread/2
		rrect.TopLeft = cr.TopLeft + spread/2
		rrect.BottomLeft = cr.BottomLeft + spread/2
		rrect.BottomRight = cr.BottomRight + spread/2

		if rrect.TopRight < 0 {
			rrect.TopRight = 0
		}
		if rrect.TopLeft < 0 {
			rrect.TopLeft = 0
		}
		if rrect.BottomLeft < 0 {
			rrect.BottomLeft = 0
		}
		if rrect.BottomRight < 0 {
			rrect.BottomRight = 0
		}
	}

	return rrect
}

var (
	ambientShadow = &shadowElevation{
		maxElevation:      24,
		bezierCurveY:      bezier.Easing(0.4, 0, 1, 0.8),
		bezierCurveBlur:   bezier.Easing(0.4, 0, 1, 0.8),
		bezierCurveSpread: bezier.Easing(0.4, 0, 1, 0.8),
		yBoundaries:       [2]int{0, 4},
		blurBoundaries:    [2]int{0, 64},
		spreadBoundaries:  [2]int{1, 4},
	}

	umbraShadow = &shadowElevation{
		maxElevation:      24,
		bezierCurveY:      bezier.Easing(0.4, 0, 1, 0.8),
		bezierCurveBlur:   bezier.Easing(0.4, 0, 1, 0.8),
		bezierCurveSpread: bezier.Easing(0.4, 0, 1, 0.8),
		yBoundaries:       [2]int{1, 16},
		blurBoundaries:    [2]int{2, 16},
		spreadBoundaries:  [2]int{1, 2},
		negativeSpread:    true,
	}

	penumbraShadow = &shadowElevation{
		maxElevation:      24,
		bezierCurveY:      bezier.Easing(0.4, 0, 1, 0.8),
		bezierCurveBlur:   bezier.Easing(0.4, 0, 1, 0.8),
		bezierCurveSpread: bezier.Easing(0.4, 0, 1, 0.8),
		yBoundaries:       [2]int{0, 32},
		blurBoundaries:    [2]int{0, 64},
		spreadBoundaries:  [2]int{1, 5},
	}
)

type shadowElevation struct {
	maxElevation      float64
	yBoundaries       boundaries
	blurBoundaries    boundaries
	spreadBoundaries  boundaries
	bezierCurveY      bezier.EasingFunc
	bezierCurveBlur   bezier.EasingFunc
	bezierCurveSpread bezier.EasingFunc
	negativeSpread    bool
}

type boundaries [2]int

func (b boundaries) At(p float64) int {
	return int(math.Round(float64(b[1]-b[0])*p)) + b[0]
}

func (se *shadowElevation) Calc(elevation float64) (y int, blur int, spread int) {
	p := elevation * 1 / (se.maxElevation - 1)

	y = se.yBoundaries.At(se.bezierCurveY(p))
	blur = se.blurBoundaries.At(se.bezierCurveBlur(p))
	spread = se.spreadBoundaries.At(se.bezierCurveSpread(p))

	if se.negativeSpread {
		spread = -spread
	}

	return
}
