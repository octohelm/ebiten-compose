package patcher

func NewBatchPatcher() BatchPatcher {
	return &batchPatcher{}
}

type Patcher interface {
	Patch() bool
}

type BatchPatcher interface {
	Patcher

	DoPatch(action Patcher)
}

func WhenChanged[T comparable](addr *T, v T) Patcher {
	return &whenChanged[T]{
		addr: addr,
		v:    v,
	}
}

type whenChanged[T comparable] struct {
	addr *T
	v    T
}

func (w *whenChanged[T]) Patch() bool {
	o := *w.addr
	n := w.v
	if o == n {
		return false
	}
	*w.addr = n
	return true
}

type batchPatcher struct {
	changed bool
}

func (s *batchPatcher) DoPatch(action Patcher) {
	action.Patch()
}

func (s *batchPatcher) Patch() bool {
	return s.changed
}
