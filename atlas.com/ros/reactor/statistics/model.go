package statistics

type Point struct {
	x int16
	y int16
}

func (p Point) X() int16 {
	return p.x
}

func (p Point) Y() int16 {
	return p.y
}

type Model struct {
	tl          Point
	br          Point
	stateInfo   map[byte][]ReactorState
	timeoutInfo map[byte]int32
}

func NewModel() *Model {
	return &Model{
		tl:          Point{0, 0},
		br:          Point{0, 0},
		stateInfo:   make(map[byte][]ReactorState, 0),
		timeoutInfo: make(map[byte]int32),
	}
}

func (m Model) Tl() Point {
	return m.tl
}

func (m Model) AddState(state byte, data []ReactorState, timeout int32) *Model {
	nm := &Model{
		tl:          m.tl,
		br:          m.br,
		stateInfo:   m.stateInfo,
		timeoutInfo: m.timeoutInfo,
	}

	nm.stateInfo[state] = data
	if timeout > -1 {
		nm.timeoutInfo[state] = timeout
	}
	return nm
}

func (m Model) SetTL(x int32, y int32) *Model {
	nm := &Model{
		tl:          Point{int16(x), int16(y)},
		br:          m.br,
		stateInfo:   m.stateInfo,
		timeoutInfo: m.timeoutInfo,
	}
	return nm
}

func (m Model) SetRB(x int32, y int32) *Model {
	nm := &Model{
		tl:          m.tl,
		br:          Point{int16(x), int16(y)},
		stateInfo:   m.stateInfo,
		timeoutInfo: m.timeoutInfo,
	}
	return nm
}

func (m Model) Type(state byte) uint32 {
	return m.stateInfo[state][0].Type()
}

type ReactorState struct {
	theType      uint32
	reactorItem  *ReactorItem
	activeSkills []uint32
	nextState    byte
}

func (s ReactorState) Type() uint32 {
	return s.theType
}

type ReactorItem struct {
	itemId   uint32
	quantity uint16
}
