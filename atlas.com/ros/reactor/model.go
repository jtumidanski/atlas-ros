package reactor

import "atlas-ros/reactor/statistics"

type Model struct {
	id             uint32
	worldId        byte
	channelId      byte
	mapId          uint32
	classification uint32
	name           string
	statistics     statistics.Model
	state          int8
	eventState     byte
	delay          uint32
	direction      byte
	x              int16
	y              int16
	alive          bool
	shouldCollect  bool
	attackHit      bool
	delayedRespawn bool
}

func (m Model) Id() uint32 {
	return m.id
}

func (m Model) Classification() uint32 {
	return m.classification
}

func (m Model) Name() string {
	return m.name
}

// Type 2 = only hit from right (kerning swamp plants), 00 is air left 02 is ground left
func (m Model) Type() int32 {
	return m.statistics.Type(m.State())
}

func (m Model) State() int8 {
	return m.state
}

func (m Model) EventState() byte {
	return m.eventState
}

func (m Model) X() int16 {
	return m.x
}

func (m Model) Y() int16 {
	return m.y
}

func (m Model) Delay() uint32 {
	return m.delay
}

func (m Model) FacingDirection() byte {
	return m.direction
}

func (m Model) Active() bool {
	return m.alive && m.Type() != -1
}

func (m Model) StateSize() byte {
	return m.statistics.StateSize(m.state)
}

func (m Model) ActiveSkills(i byte) []uint32 {
	return m.statistics.ActiveSkills(m.state, i)
}

func (m Model) NextState(b byte) int8 {
	return m.statistics.NextState(m.state, b)
}

func (m Model) WorldId() byte {
	return m.worldId
}

func (m Model) ChannelId() byte {
	return m.channelId
}

func (m Model) MapId() uint32 {
	return m.mapId
}

func (m Model) Timeout() int32 {
	return m.statistics.Timeout(m.state)
}

func (m Model) TimeoutState() int8 {
	return m.statistics.TimeoutState(m.state)
}

type Modifier func(m *Model)

func incrementState() Modifier {
	return func(m *Model) {
		m.state++
	}
}

func setState(state int8) Modifier {
	return func(m *Model) {
		m.state = state
	}
}

func advanceState(b byte) Modifier {
	return func(m *Model) {
		m.state = m.statistics.NextState(m.state, b)
	}
}

func shouldCollect(value bool) Modifier {
	return func(m *Model) {
		m.shouldCollect = value
	}
}
