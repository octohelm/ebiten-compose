package internal

import "fmt"

type State[T comparable] interface {
	Value() T
	Update(value T)
	UpdateFunc(func(prev T) T)
}

type StateHook[T comparable] struct {
	State         T
	OnStateChange func()
}

func (s *StateHook[T]) String() string {
	return fmt.Sprintf("UseState: %v", s.State)
}

func (s *StateHook[T]) Value() T {
	return s.State
}

func (s *StateHook[T]) Update(v T) {
	if v != s.State {
		s.State = v
		s.OnStateChange()
	}
}

func (s *StateHook[T]) UpdateFunc(fn func(prev T) T) {
	if nextState := fn(s.State); nextState != s.State {
		s.State = nextState
		s.OnStateChange()
	}
}

func (s *StateHook[T]) UpdateHook(next Hook) {
	if n, ok := next.(*StateHook[T]); ok {
		// context may change, should bind the latest callback
		s.OnStateChange = n.OnStateChange
	}
}
