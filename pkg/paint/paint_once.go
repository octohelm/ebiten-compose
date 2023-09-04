package paint

import (
	"sync/atomic"

	"gioui.org/layout"
	"gioui.org/op"
)

type PaintOnce struct {
	changed uint32
	call    op.CallOp
}

func (once *PaintOnce) Reset() {
	atomic.StoreUint32(&once.changed, 1)
}

func (once *PaintOnce) Paint(gtx layout.Context, paint func(gtx layout.Context)) {
	if atomic.LoadUint32(&once.changed) != 0 {
		newGtx := gtx
		newGtx.Ops = new(op.Ops)
		macro := op.Record(newGtx.Ops)
		paint(newGtx)
		once.call = macro.Stop()
	}

	once.call.Add(gtx.Ops)
	atomic.StoreUint32(&once.changed, 0)
}
