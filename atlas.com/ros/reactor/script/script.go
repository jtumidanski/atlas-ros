package script

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Context struct {
	WorldId               byte
	ChannelId             byte
	MapId                 uint32
	CharacterId           uint32
	ReactorId             uint32
	ReactorClassification uint32
}

type ActFunc func(l logrus.FieldLogger, db *gorm.DB, c Context)

type HitFunc func(l logrus.FieldLogger, db *gorm.DB, c Context)

type TouchFunc func(l logrus.FieldLogger, db *gorm.DB, c Context)

type ReleaseFunc func(l logrus.FieldLogger, db *gorm.DB, c Context)

type Script interface {
	ReactorClassification() uint32

	Act(l logrus.FieldLogger, db *gorm.DB, c Context)

	Hit(l logrus.FieldLogger, db *gorm.DB, c Context)

	Touch(l logrus.FieldLogger, db *gorm.DB, c Context)

	Release(l logrus.FieldLogger, db *gorm.DB, c Context)
}

type Action func(l logrus.FieldLogger, db *gorm.DB) func(c Context) func(script Script)

func InvokeAct(l logrus.FieldLogger, db *gorm.DB) func(c Context) func(script Script) {
	return func(c Context) func(script Script) {
		return func(script Script) {
			script.Act(l, db, c)
		}
	}
}

func InvokeHit(l logrus.FieldLogger, db *gorm.DB) func(c Context) func(script Script) {
	return func(c Context) func(script Script) {
		return func(script Script) {
			script.Hit(l, db, c)
		}
	}
}

func InvokeTouch(l logrus.FieldLogger, db *gorm.DB) func(c Context) func(script Script) {
	return func(c Context) func(script Script) {
		return func(script Script) {
			script.Touch(l, db, c)
		}
	}
}

func InvokeRelease(l logrus.FieldLogger, db *gorm.DB) func(c Context) func(script Script) {
	return func(c Context) func(script Script) {
		return func(script Script) {
			script.Release(l, db, c)
		}
	}
}
