package layout

import "github.com/octohelm/ebiten-compose/pkg/patcher"

type WeightSetter interface {
	SetWeight(weight float32) bool
}

type WeightGetter interface {
	Weight() (float32, bool)
}

var _ WeightGetter = &FlexWeight{}
var _ WeightSetter = &FlexWeight{}

type FlexWeight struct {
	weight float32
}

func (w *FlexWeight) Weight() (float32, bool) {
	return w.weight, w.weight > 0
}

func (w *FlexWeight) SetWeight(weight float32) bool {
	return patcher.WhenChanged(&w.weight, weight).Patch()
}
