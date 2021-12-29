package _map

import "github.com/sirupsen/logrus"

func MonstersCount(_ logrus.FieldLogger) func(worldId byte, channelId byte, mapId uint32) uint32 {
	return func(worldId byte, channelId byte, mapId uint32) uint32 {
		// TODO AT-10
		return 0
	}
}

func MonsterCount(_ logrus.FieldLogger) func(worldId byte, channelId byte, mapId uint32, monsterId uint32) int {
	return func(worldId byte, channelId byte, mapId uint32, monsterId uint32) int {
		// TODO AT-10
		return 0
	}
}

func HitReactor(_ logrus.FieldLogger) func(worldId byte, channelId byte, mapId uint32, reactorName string) {
	return func(worldId byte, channelId byte, mapId uint32, reactorName string) {
		// TODO AT-11
	}
}

func MoveEnvironment(_ logrus.FieldLogger) func(worldId byte, channelId byte, mapId uint32, environment string, mode byte) {
	return func(worldId byte, channelId byte, mapId uint32, environment string, mode byte) {
		// TODO AT-12
	}
}

func ToggleEnvironment(_ logrus.FieldLogger) func(worldId byte, channelId byte, mapId uint32, environment string) {
	return func(worldId byte, channelId byte, mapId uint32, environment string) {
		// TODO AT-13
	}
}

func KillAllMonsters(_ logrus.FieldLogger) func(worldId byte, channelId byte, mapId uint32) {
	return func(worldId byte, channelId byte, mapId uint32) {
		// TODO AT-14
	}
}

func SetSummonState(_ logrus.FieldLogger) func(worldId byte, channelId byte, mapId uint32, state bool) {
	return func(worldId byte, channelId byte, mapId uint32, state bool) {
		// TODO
	}
}

func SummonState(_ logrus.FieldLogger) func(worldId byte, channelId byte, mapId uint32) bool {
	return func(worldId byte, channelId byte, mapId uint32) bool {
		// TODO
		return false
	}
}

func CharactersInMap(_ logrus.FieldLogger) func(worldId byte, channelId byte, mapId uint32) []uint32 {
	return func(worldId byte, channelId byte, mapId uint32) []uint32 {
		// TODO AT-15
		return []uint32{}
	}
}
