package alignment

type Alignment int

const (
	TopStart Alignment = iota
	Top
	TopEnd
	End
	BottomStart
	Bottom
	BottomEnd
	Start
	Center

	Baseline
)

const Middle = Center
