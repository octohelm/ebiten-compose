package position

type Position int

const (
	Top Position = iota + 1
	Right
	Bottom
	Left
	TopLeft
	BottomLeft
	TopRight
	BottomRight
)
