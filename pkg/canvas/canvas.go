package canvas

import (
	"image"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/octohelm/ebiten-compose/pkg/unit"
)

type Dimensions struct {
	Size     image.Point
	Baseline float32
}

type Canvas interface {
	Offset() image.Point
	Size() image.Point

	Translate(x int, y int)

	Image() *ebiten.Image

	PushOp(ops ...Op)

	Dp(dp unit.Dp) int
}

type Controller interface {
	Canvas
	Draw()
}

type Op interface {
	Draw()
}

func OpFunc(draw func()) Op {
	return &drawOpFunc{
		draw: draw,
	}
}

type drawOpFunc struct {
	draw func()
}

func (d *drawOpFunc) Draw() {
	d.draw()
}

func NewRootCanvas(img *ebiten.Image, deviceScaleFactor float32) Controller {
	return &canvas{
		position: position{
			size: img.Bounds().Max,
		},
		deviceScaleFactor: deviceScaleFactor,
		img:               img,
	}
}

type stack struct {
	ops []Op
}

func (c *stack) Draw() {
	for i := range c.ops {
		c.ops[i].Draw()
	}
	c.ops = nil
}

func (c *stack) PushOp(ops ...Op) {
	c.ops = append(c.ops, ops...)
}

type position struct {
	size   image.Point
	offset image.Point
}

func (c *position) Offset() image.Point {
	return c.offset
}

func (c *position) AddOffset(offset image.Point) image.Point {
	return image.Pt(offset.X+c.offset.X, offset.Y+c.offset.Y)
}

func (c *position) Translate(x int, y int) {
	c.offset.X += x
	c.offset.Y += y
}

func (c *position) Size() image.Point {
	return c.size
}

type canvas struct {
	stack
	position
	img               *ebiten.Image
	deviceScaleFactor float32
}

func (c *canvas) DeviceScaleFactor() float32 {
	return c.deviceScaleFactor
}

func (c *canvas) Dp(dp unit.Dp) int {
	return int(math.Round(float64(dp) * float64(c.deviceScaleFactor)))
}

func (c *canvas) Image() *ebiten.Image {
	return c.img
}

func NewCanvas(c Canvas, x int, y int) Canvas {
	return &canvasWithParent{
		parent: c,
		position: position{
			size: image.Pt(x, y),
		},
	}
}

type canvasWithParent struct {
	parent Canvas
	position
}

func (c *canvasWithParent) Dp(dp unit.Dp) int {
	return c.parent.Dp(dp)
}

func (c *canvasWithParent) Offset() image.Point {
	return c.position.AddOffset(c.parent.Offset())
}

func (c *canvasWithParent) Image() *ebiten.Image {
	return c.parent.Image()
}

func (c *canvasWithParent) PushOp(ops ...Op) {
	c.parent.PushOp(ops...)
}

func NewCanvasWithImage(c Canvas, img *ebiten.Image) Canvas {
	return &canvasImageWithParent{
		parent: c,
		img:    img,
		position: position{
			size: img.Bounds().Max,
		},
	}
}

type canvasImageWithParent struct {
	parent Canvas
	img    *ebiten.Image
	position
}

func (c *canvasImageWithParent) Dp(dp unit.Dp) int {
	return c.parent.Dp(dp)
}

func (c *canvasImageWithParent) Image() *ebiten.Image {
	return c.img
}

func (c *canvasImageWithParent) PushOp(ops ...Op) {
	c.parent.PushOp(ops...)
}
