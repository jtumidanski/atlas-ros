package _map

import "github.com/sirupsen/logrus"

func MonstersCount(_ logrus.FieldLogger) func(worldId byte, channelId byte, mapId uint32) uint32 {
	return func(worldId byte, channelId byte, mapId uint32) uint32 {
		// TODO
		return 0
	}
}

func MonsterCount(_ logrus.FieldLogger) func(worldId byte, channelId byte, mapId uint32, monsterId uint32) int {
	return func(worldId byte, channelId byte, mapId uint32, monsterId uint32) int {
		// TODO
		return 0
	}
}

func HitReactor(_ logrus.FieldLogger) func(worldId byte, channelId byte, mapId uint32, reactorName string) {
	return func(worldId byte, channelId byte, mapId uint32, reactorName string) {
		// TODO
	}
}

func MoveEnvironment(_ logrus.FieldLogger) func(worldId byte, channelId byte, mapId uint32, environment string, mode byte) {
	return func(worldId byte, channelId byte, mapId uint32, environment string, mode byte) {
		// TODO
	}
}

func ToggleEnvironment(_ logrus.FieldLogger) func(worldId byte, channelId byte, mapId uint32, environment string) {
	return func(worldId byte, channelId byte, mapId uint32, environment string) {
		// TODO
	}
}

func KillAllMonsters(_ logrus.FieldLogger) func(worldId byte, channelId byte, mapId uint32) {
	return func(worldId byte, channelId byte, mapId uint32) {
		// TODO
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
		// TODO
		return []uint32{}
	}
}