package script

import "github.com/sirupsen/logrus"

type Context struct {
	WorldId         byte
	ChannelId       byte
	MapId           uint32
	CharacterId     uint32
	ReactorId       uint32
	ReactorUniqueId uint32
}

type ActFunc func(l logrus.FieldLogger, c Context)

type HitFunc func(l logrus.FieldLogger, c Context)

type TouchFunc func(l logrus.FieldLogger, c Context)

type ReleaseFunc func(l logrus.FieldLogger, c Context)

type Script interface {
	ReactorId() uint32

	Act(l logrus.FieldLogger, c Context)

	Hit(l logrus.FieldLogger, c Context)

	Touch(l logrus.FieldLogger, c Context)

	Release(l logrus.FieldLogger, c Context)
}

type Action func(l logrus.FieldLogger) func(c Context) func(script Script)

func InvokeAct(l logrus.FieldLogger) func(c Context) func(script Script) {
	return func(c Context) func(script Script) {
		return func(script Script) {
			script.Act(l, c)
		}
	}
}

func InvokeHit(l logrus.FieldLogger) func(c Context) func(script Script) {
	return func(c Context) func(script Script) {
		return func(script Script) {
			script.Hit(l, c)
		}
	}
}

func InvokeTouch(l logrus.FieldLogger) func(c Context) func(script Script) {
	return func(c Context) func(script Script) {
		return func(script Script) {
			script.Touch(l, c)
		}
	}
}

func InvokeRelease(l logrus.FieldLogger) func(c Context) func(script Script) {
	return func(c Context) func(script Script) {
		return func(script Script) {
			script.Release(l, c)
		}
	}
}
