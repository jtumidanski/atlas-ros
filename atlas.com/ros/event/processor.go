package event

import (
	"github.com/sirupsen/logrus"
	"time"
)

// IdProvider is a function which provides an event id
type IdProvider func() uint32

// ParticipatingCharacterIdProvider retrieves an event id given a character who is presumably participating in the event
func ParticipatingCharacterIdProvider(characterId uint32) IdProvider {
	return func() uint32 {
		// TODO query the event id by character who is participating
		return 0
	}
}

// Get returns the event given the id provided by the IdProvider
func Get(l logrus.FieldLogger) func(provider IdProvider) (*Model, error) {
	return func(provider IdProvider) (*Model, error) {
		id := provider()
		return &Model{id: id}, nil
	}
}

func GetByParticipatingCharacter(l logrus.FieldLogger) func(characterId uint32) (*Model, error) {
	return func(characterId uint32) (*Model, error) {
		return Get(l)(ParticipatingCharacterIdProvider(characterId))
	}
}

func GetProperty(l logrus.FieldLogger) func(id uint32, name string) int32 {
	return func(id uint32, name string) int32 {
		return 0
	}
}

func GetStringProperty(l logrus.FieldLogger) func(id uint32, name string) string {
	return func(id uint32, name string) string {
		return ""
	}
}

func SetProperty(l logrus.FieldLogger) func(id uint32, name string, value int32) {
	return func(id uint32, name string, value int32) {

	}
}

func SetStringProperty(l logrus.FieldLogger) func(id uint32, name string, value string) {
	return func(id uint32, name string, value string) {

	}
}

func AllReactorsActivatedInMap(l logrus.FieldLogger) func(worldId byte, channelId byte, characterId uint32, mapId uint32, min uint32, max uint32) bool {
	return func(worldId byte, channelId byte, characterId uint32, mapId uint32, min uint32, max uint32) bool {
		return false
	}
}

func ShowClearEffect(l logrus.FieldLogger) func(worldId byte, channelId byte, characterId uint32, mapId uint32) {
	return func(worldId byte, channelId byte, characterId uint32, mapId uint32) {

	}
}

func ParticipatingInEvent(l logrus.FieldLogger) func(characterId uint32) bool {
	return func(characterId uint32) bool {
		_, err := Get(l)(ParticipatingCharacterIdProvider(characterId))
		if err != nil {
			l.WithError(err).Warnf("Unable to locate event by participating character %d. Assuming they're not participating.", characterId)
			return false
		}
		return true
	}
}

func BlueMessageParticipants(l logrus.FieldLogger) func(provider IdProvider, message string) {
	return func(provider IdProvider, message string) {
		_, err := Get(l)(provider)
		if err != nil {
			l.WithError(err).Errorf("Unable to message event participants, as the event could not be located.")
			return
		}
		// TODO
	}
}

func InvokeFunction(l logrus.FieldLogger) func(id uint32, name string) {
	return func(id uint32, name string) {
		
	}
}

func ClearPartyQuest(l logrus.FieldLogger) func(id uint32) {
	return func(id uint32) {

	}
}

func GiveParticipantsExperience(l logrus.FieldLogger) func(id uint32, amount int16) {
	return func(id uint32, amount int16) {

	}
}

func StartTimer(l logrus.FieldLogger) func(id uint32, duration time.Duration) {
	return func(id uint32, duration time.Duration) {

	}
}