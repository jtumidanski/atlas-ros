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
	stateInfo   map[int8][]ReactorState
	timeoutInfo map[int8]int32
}

func NewModel() *Model {
	return &Model{
		tl:          Point{0, 0},
		br:          Point{0, 0},
		stateInfo:   make(map[int8][]ReactorState, 0),
		timeoutInfo: make(map[int8]int32),
	}
}

func (m Model) Tl() Point {
	return m.tl
}

func (m Model) AddState(state int8, data []ReactorState, timeout int32) *Model {
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

func (m Model) Type(state int8) int32 {
	return m.stateInfo[state][0].Type()
}

func (m Model) NextState(state int8, index byte) int8 {
	val, ok := m.stateInfo[state]
	if !ok {
		return -1
	}
	if len(val) < int(index)+1 {
		return -1
	}
	nsd := m.stateInfo[state][index]
	return nsd.NextState()
}

func (m Model) StateSize(state int8) byte {
	return byte(len(m.stateInfo[state]))
}

func (m Model) ActiveSkills(state int8, i byte) []uint32 {
	val, ok := m.stateInfo[state]
	if !ok {
		return make([]uint32, 0)
	}
	if len(val) < int(i)+1 {
		return make([]uint32, 0)
	}
	s := m.stateInfo[state][i]
	return s.ActiveSkills()
}

func (m Model) Timeout(state int8) int32 {
	if val, ok := m.timeoutInfo[state]; ok {
		return val
	}
	return -1
}

func (m Model) TimeoutState(state int8) int8 {
	if s, ok := m.stateInfo[state]; ok {
		if len(s) <= 0 {
			return -1
		}
		return s[len(s)-1].NextState()
	}
	return -1
}

type ReactorState struct {
	theType      int32
	reactorItem  *ReactorItem
	activeSkills []uint32
	nextState    int8
}

func (s ReactorState) Type() int32 {
	return s.theType
}

func (s ReactorState) NextState() int8 {
	return s.nextState
}

func (s ReactorState) ActiveSkills() []uint32 {
	return s.activeSkills
}

type ReactorItem struct {
	itemId   uint32
	quantity uint16
}
