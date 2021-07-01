package reactor

import "atlas-ros/reactor/statistics"

type Model struct {
	uniqueId       uint32
	worldId        byte
	channelId      byte
	mapId          uint32
	id             uint32
	name           string
	statistics     statistics.Model
	state          byte
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

func (m Model) UniqueId() uint32 {
	return m.uniqueId
}

func (m Model) Id() uint32 {
	return m.id
}

func (m Model) Name() string {
	return m.name
}

func (m Model) Type() uint32 {
	return m.statistics.Type(m.State())
}

func (m Model) State() byte {
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
