package renderer

import (
	"context"
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"

	"github.com/octohelm/ebiten-compose/pkg/compose"
	"github.com/octohelm/ebiten-compose/pkg/compose/internal"
	"github.com/octohelm/ebiten-compose/pkg/compose/modifier"
	"github.com/octohelm/ebiten-compose/pkg/node"
)

func CreateRoot() *Root {
	rootVNode := compose.Box(
		modifier.DisplayName("Root"),
		modifier.FillMaxSize(),
		modifier.BackgroundColor(color.White),
	).(internal.VNodeAccessor)

	initialVNode(context.Background(), rootVNode)

	r := &Root{
		window: newWindow(),
		Root:   Portal(rootVNode.Node()).(internal.VNodeAccessor),
	}

	r.cq.Start()
	return r
}

type Root struct {
	Root internal.VNodeAccessor

	cq commitQueue

	window *window
}

func (r *Root) Act(fn func()) {
	fn()
	r.cq.ForceCommit()
}

func (r *Root) Close() error {
	return r.cq.Close()
}

func (r *Root) Wait() error {
	r.window.element = r.Root.Node()
	return ebiten.RunGame(r.window)
}

func (r *Root) Render(ctx context.Context, vnode internal.VNode) {
	nextRoot := Portal(r.Root.Node()).Children(vnode).(internal.VNodeAccessor)
	r.patchVNode(ctx, r.Root, nextRoot)
	r.Root = nextRoot
	r.cq.ForceCommit()
}

func (r *Root) sameVNode(vnode1 internal.VNodeAccessor, vnode2 internal.VNodeAccessor) bool {
	return vnode1 == vnode2 || internal.SameComponent(vnode1.Type(), vnode2.Type()) && vnode1.Key() == vnode2.Key()
}

func (r *Root) patchVNode(ctx context.Context, oldVNode internal.VNodeAccessor, vnode internal.VNodeAccessor) {
	if vnode == oldVNode {
		return
	}

	switch vnode.Type().(type) {
	case internal.Element, internal.Fragment, internal.Portal:
		// trigger Build to apply data in context
		_ = vnode.Type().Build(&buildContext{
			vnode:   vnode,
			Context: ctx,
		})
		vnode.PutChildVNodes(vnode.ChildVNodes()...)

		r.mount(ctx, oldVNode, vnode)
		r.cq.Dispatch(func() {
			vnode.DidMount()
		})
	default:
		// only component need to render
		var doRender func(ctx context.Context, oldVNode internal.VNodeAccessor, vnode internal.VNodeAccessor)

		doRender = func(ctx context.Context, oldVNode internal.VNodeAccessor, vnode internal.VNodeAccessor) {
			childCtx := ctx

			if cp, ok := vnode.Type().(internal.ContextProvider); ok {
				childCtx = cp.GetChildContext(ctx)
			}

			vnode.WillRender(oldVNode)

			child := vnode.Type().Build(&buildContext{
				Context: childCtx,
				vnode:   vnode,
			})

			vnode.PutChildVNodes(compose.Fragment().Children(child))

			r.mount(childCtx, oldVNode, vnode)

			rendered := vnode.Clone()
			vnode.OnUpdate(func(vn internal.VNodeAccessor) {
				doRender(ctx, rendered, vn)
			})

			r.cq.Dispatch(func() {
				vnode.DidMount()

				// FIXME trigger render
				//r.window.Update()
			})
		}

		doRender(ctx, oldVNode, vnode)
	}
}

func (r *Root) mount(ctx context.Context, oldVNode internal.VNodeAccessor, vnode internal.VNodeAccessor) {
	if oldVNode == nil {
		initialVNode(ctx, vnode)

		mounted := vnode.MountedNode()
		if mounted == nil {
			panic(fmt.Errorf("required mount point: %v", vnode))
		}

		r.addVNodes(
			ctx,
			mounted,
			nil,
			vnode.ChildVNodeAccessors(),
			0,
			len(vnode.ChildVNodeAccessors())-1,
		)

		return
	}

	vnode.WillMount(oldVNode)

	oldChildVNodes := oldVNode.ChildVNodeAccessors()
	childVNodes := vnode.ChildVNodeAccessors()

	mounted := vnode.MountedNode()

	if mounted == nil {
		panic(fmt.Errorf("required mount point: %v", vnode))
	}

	if len(oldChildVNodes) != 0 && len(childVNodes) != 0 {
		r.patchVNodes(ctx, mounted, oldChildVNodes, childVNodes)
	} else if len(childVNodes) != 0 {
		r.addVNodes(ctx, mounted, nil, childVNodes, 0, len(childVNodes)-1)
	} else if len(oldChildVNodes) != 0 {
		r.removeVNodes(ctx, mounted, oldChildVNodes, 0, len(oldChildVNodes)-1)
	}
}

func initialVNode(ctx context.Context, vnode internal.VNodeAccessor) {
	if w, ok := vnode.Type().(internal.Element); ok {
		vnode.Mount(w.New())

		// TODO when change need to trigger re-paint
		_ = compose.ApplyModifiers(ctx, vnode.Node(), vnode.Modifiers()...)
	}
}

func indexVNodeList(l []internal.VNodeAccessor, n, i int) internal.VNodeAccessor {
	if maxIndex := n - 1; maxIndex >= i && i >= 0 {
		return l[i]
	}
	return nil
}

func (r *Root) patchVNodes(ctx context.Context, parentNode internal.Element, oldChildren []internal.VNodeAccessor, newChildren []internal.VNodeAccessor) {
	oldStartIdx := 0
	newStartIdx := 0
	oldN := len(oldChildren)
	oldEndIdx := oldN - 1
	oldStartVNode := indexVNodeList(oldChildren, oldN, 0)
	oldEndVNode := indexVNodeList(oldChildren, oldN, oldEndIdx)
	newN := len(newChildren)
	newEndIdx := newN - 1
	newStartVNode := indexVNodeList(newChildren, newN, 0)
	newEndVNode := indexVNodeList(newChildren, newN, newEndIdx)

	var oldKeyToIdx map[internal.Key]int
	var idxInOld int
	var elmToMove internal.VNodeAccessor
	var beforeNode internal.Element

	for oldStartIdx <= oldEndIdx && newStartIdx <= newEndIdx {
		if oldStartVNode == nil {
			oldStartIdx++
			oldStartVNode = indexVNodeList(oldChildren, oldN, oldStartIdx) // VNode might have been moved left
		} else if oldEndVNode == nil {
			oldEndIdx--
			oldEndVNode = indexVNodeList(oldChildren, oldN, oldEndIdx)
		} else if newStartVNode == nil {
			newStartIdx++
			newStartVNode = indexVNodeList(newChildren, newN, newStartIdx)
		} else if newEndVNode == nil {
			newEndIdx--
			newEndVNode = indexVNodeList(newChildren, newN, newEndIdx)
		} else if r.sameVNode(oldStartVNode, newStartVNode) {
			r.patchVNode(ctx, oldStartVNode, newStartVNode)
			oldStartIdx++
			oldStartVNode = indexVNodeList(oldChildren, oldN, oldStartIdx)
			newStartIdx++
			newStartVNode = indexVNodeList(newChildren, newN, newStartIdx)
		} else if r.sameVNode(oldEndVNode, newEndVNode) {
			r.patchVNode(ctx, oldEndVNode, newEndVNode)
			oldEndIdx--
			oldEndVNode = indexVNodeList(oldChildren, oldN, oldEndIdx)
			newEndIdx--
			newEndVNode = indexVNodeList(newChildren, newN, newEndIdx)
		} else if r.sameVNode(oldStartVNode, newEndVNode) {
			// VNode moved right
			r.patchVNode(ctx, oldStartVNode, newEndVNode)
			r.insertBefore("move right", parentNode, oldStartVNode.Node(), oldEndVNode.Node().NextSibling())
			oldStartIdx++
			oldStartVNode = indexVNodeList(oldChildren, oldN, oldStartIdx)
			newEndIdx--
			newEndVNode = indexVNodeList(newChildren, newN, newEndIdx)
		} else if r.sameVNode(oldEndVNode, newStartVNode) {
			// VNode moved left
			r.patchVNode(ctx, oldEndVNode, newStartVNode)
			r.insertBefore("move left", parentNode, oldEndVNode.Node(), oldStartVNode.Node())

			oldEndIdx--
			oldEndVNode = indexVNodeList(oldChildren, oldN, oldEndIdx)
			newStartIdx++
			newStartVNode = indexVNodeList(newChildren, newN, newStartIdx)
		} else {
			if oldKeyToIdx == nil {
				oldKeyToIdx = createKeyToOldIdx(oldChildren, oldStartIdx, oldEndIdx)
			}

			idxInOld = oldKeyToIdx[newStartVNode.Key()]

			if idxInOld == 0 {
				r.patchVNode(ctx, nil, newStartVNode)
				r.insertBefore("first", parentNode, newStartVNode.Node(), oldStartVNode.Node())
			} else {
				elmToMove = indexVNodeList(oldChildren, oldN, idxInOld)
				if !internal.SameComponent(elmToMove.Type(), newStartVNode.Type()) {
					r.patchVNode(ctx, nil, newStartVNode)
					r.insertBefore("not same component", parentNode, newStartVNode.Node(), oldStartVNode.Node())
				} else {
					r.patchVNode(ctx, elmToMove, newStartVNode)
					oldChildren[idxInOld] = nil
					r.insertBefore("same component", parentNode, elmToMove.Node(), oldStartVNode.Node())
				}
			}
			newStartIdx++
			newStartVNode = indexVNodeList(newChildren, newN, newStartIdx)
		}
	}

	if oldStartIdx <= oldEndIdx || newStartIdx <= newEndIdx {
		if oldStartIdx > oldEndIdx {
			beforeNode = nil
			if n := indexVNodeList(newChildren, newN, newEndIdx+1); n != nil {
				beforeNode = n.Node()
			}
			r.addVNodes(ctx, parentNode, beforeNode, newChildren, newStartIdx, newEndIdx)
		} else {
			r.removeVNodes(ctx, parentNode, oldChildren, oldStartIdx, oldEndIdx)
		}
	}
}

func (r *Root) addVNodes(ctx context.Context, parentNode internal.Element, beforeNode internal.Element, vnodes []internal.VNodeAccessor, startIdx int, endIdx int) {
	for startIdx <= endIdx {
		vn := vnodes[startIdx]
		r.patchVNode(ctx, nil, vn)
		if !vn.IsRoot() {
			r.insertBefore("add", parentNode, vn.Node(), beforeNode)
		}
		startIdx++
	}
}

func (r *Root) removeVNodes(ctx context.Context, parentNode internal.Element, vnodes []internal.VNodeAccessor, startIdx int, endIdx int) {
	for startIdx <= endIdx {
		vn := vnodes[startIdx]

		if vn.Node() == nil {
			r.removeVNodes(ctx, parentNode, vn.ChildVNodeAccessors(), 0, len(vn.ChildVNodeAccessors())-1)
		} else {
			if !vn.IsRoot() {
				r.removeChild(parentNode, vn.Node())
			}
		}
		_ = vn.Destroy()

		startIdx++
	}
}

func createKeyToOldIdx(childVNodes []internal.VNodeAccessor, beginIdx int, endIdx int) map[internal.Key]int {
	keyToIndexMap := map[internal.Key]int{}
	for i := beginIdx; i <= endIdx; i++ {
		vn := childVNodes[i]
		key := vn.Key()
		if key != "" {
			keyToIndexMap[key] = i
		}
	}
	return keyToIndexMap
}

func (r *Root) insertBefore(cause string, parent, new, old node.Node) {
	r.cq.Dispatch(func() {
		parent.InsertBefore(new, old)
	})
}

func (r *Root) removeChild(parent, old node.Node) {
	r.cq.Dispatch(func() {
		parent.RemoveChild(old)
	})
}

type buildContext struct {
	context.Context

	vnode internal.VNodeAccessor
}

func (b *buildContext) VNode() internal.VNodeAccessor {
	return b.vnode
}

func (b *buildContext) ChildVNodes() []internal.VNode {
	return b.vnode.ChildVNodes()
}
