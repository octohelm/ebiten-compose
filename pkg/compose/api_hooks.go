package compose

import "github.com/octohelm/ebiten-compose/pkg/compose/internal"

func UseState[T comparable](ctx BuildContext, defaultState T) internal.State[T] {
	vn := ctx.(internal.BuildContextAccessor).VNode()

	return internal.UseHook(vn, &internal.StateHook[T]{
		State:         defaultState,
		OnStateChange: vn.Update,
	})
}

func UseEffect(ctx BuildContext, setup func() func(), deps []interface{}) {
	vn := ctx.(internal.BuildContextAccessor).VNode()

	internal.UseHook(vn, &internal.EffectHook{
		Setup: setup,
		Deps:  deps,
	})
}

func UseMemo[T any](ctx BuildContext, setup func() T, deps []interface{}) T {
	vn := ctx.(internal.BuildContextAccessor).VNode()

	h := internal.UseHook(vn, &internal.MemoHook[T]{
		Setup: setup,
		Deps:  deps,
	})

	return h.Memorised()
}

func UseRef[T any](ctx BuildContext, initialValue T) *internal.Ref[T] {
	vn := ctx.(internal.BuildContextAccessor).VNode()

	h := internal.UseHook(vn, &internal.RefHook[T]{
		Ref: internal.Ref[T]{
			Current: initialValue,
		},
	})

	return &h.Ref
}
