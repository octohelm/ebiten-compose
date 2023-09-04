package internal

import (
	"reflect"
)

func H(component Component, modifiers ...Modifier) VNode {
	return &vnode{
		typ:       component,
		modifiers: modifiers,
	}
}

type Key string

type VNode interface {
	Children(vnodes ...VNode) VNode
}

type VNodeAccessor interface {
	Key() Key
	Type() Component
	Modifiers() []Modifier
	ChildVNodes() []VNode
	PutChildVNodes(childVNodes ...VNode)
	ChildVNodeAccessors() []VNodeAccessor

	IsRoot() bool
	MountedNode() Element
	Node() Element
	Mount(node Element)
	BindParent(va VNodeAccessor)

	OnUpdate(fn func(va VNodeAccessor))
	Update()

	WillMount(oldVNode VNodeAccessor)
	WillRender(oldVNode VNodeAccessor)

	Use(hook Hook) Hook

	DidMount()

	Destroy() error

	Clone() VNodeAccessor
	Hooks() Hooks
}

var _ VNodeAccessor = &vnode{}

type vnode struct {
	key         Key
	typ         Component
	modifiers   []Modifier
	childVNodes []VNode

	childVNodeAccessors []VNodeAccessor
	parent              VNodeAccessor
	node                Element

	update func(vn VNodeAccessor)

	hooks Hooks
}

func (v *vnode) Hooks() Hooks {
	return v.hooks
}

func (v vnode) Clone() VNodeAccessor {
	return &v
}

func (v vnode) Children(children ...VNode) VNode {
	v.childVNodes = children
	return &v
}

func (v *vnode) ChildVNodes() []VNode {
	return v.childVNodes
}

func (v *vnode) ChildVNodeAccessors() []VNodeAccessor {
	return v.childVNodeAccessors
}

func (v *vnode) BindParent(va VNodeAccessor) {
	v.parent = va
}

func (v *vnode) Modifiers() []Modifier {
	return v.modifiers
}

func (v *vnode) PutChildVNodes(childVNodes ...VNode) {
	v.childVNodeAccessors = make([]VNodeAccessor, 0, len(childVNodes))

	for i := range childVNodes {
		if va, ok := childVNodes[i].(VNodeAccessor); ok {
			va.BindParent(v)
			v.childVNodeAccessors = append(v.childVNodeAccessors, va)
		}
	}
}

func (v *vnode) Type() Component {
	return v.typ
}

func (v *vnode) Key() Key {
	return v.key
}

func (v *vnode) Node() Element {
	return v.node
}

func (v *vnode) Mount(node Element) {
	v.node = node
}

func (v *vnode) MountedNode() Element {
	mounted := v.node
	if mounted == nil && v.parent != nil {
		return v.parent.MountedNode()
	}
	return mounted
}

func (v *vnode) IsRoot() bool {
	if r, ok := v.typ.(interface{ IsRoot() bool }); ok {
		return r.IsRoot()
	}

	return false
}

func (v *vnode) OnUpdate(fn func(vn VNodeAccessor)) {
	v.update = fn
}

func (v *vnode) Update() {
	if v.update != nil {
		v.update(v)
	}
}

func (v *vnode) Use(hook Hook) Hook {
	return v.hooks.use(hook)
}

func (v *vnode) WillMount(oldVNode VNodeAccessor) {
	if oldVNode != nil {
		v.node = oldVNode.Node()
	}
}

func (v *vnode) WillRender(oldVNode VNodeAccessor) {
	if oldVNode != nil {
		v.hooks = oldVNode.Hooks()
	}
	v.hooks.Reset()
}

func (v *vnode) DidMount() {
	v.hooks.commit()
}

func (v *vnode) Destroy() error {
	v.node = nil
	v.hooks.destroy()
	return nil
}

func SameComponent(type1 Component, type2 Component) bool {
	if typeE1, ok := type1.(Element); ok {
		if typeE2, ok := type2.(Element); ok {
			return typeE1 == typeE2
		}
	}

	t1 := reflect.TypeOf(type1)
	for t1.Kind() == reflect.Ptr {
		t1 = t1.Elem()
	}

	t2 := reflect.TypeOf(type2)
	for t2.Kind() == reflect.Ptr {
		t2 = t2.Elem()
	}

	return t1 == t2
}

func UseHook[T Hook](v VNodeAccessor, h T) T {
	return v.Use(h).(T)
}
