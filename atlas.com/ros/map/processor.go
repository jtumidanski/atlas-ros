package _map

import "github.com/sirupsen/logrus"

func MonstersCount(l logrus.FieldLogger) func(worldId byte, channelId byte, mapId uint32) uint32 {
	return func(worldId byte, channelId byte, mapId uint32) uint32 {
		return 0
	}
}

func MonsterCount(l logrus.FieldLogger) func(worldId byte, channelId byte, mapId uint32, monsterId uint32) int {
	return func(worldId byte, channelId byte, mapId uint32, monsterId uint32) int {
		return 0
	}
}

func HitReactor(l logrus.FieldLogger) func(worldId byte, channelId byte, mapId uint32, reactorName string) {
	return func(worldId byte, channelId byte, mapId uint32, reactorName string) {

	}
}

func MoveEnvironment(l logrus.FieldLogger) func(worldId byte, channelId byte, mapId uint32, environment string, mode byte) {
	return func(worldId byte, channelId byte, mapId uint32, environment string, mode byte) {

	}
}

func ToggleEnvironment(l logrus.FieldLogger) func(worldId byte, channelId byte, mapId uint32, environment string) {
	return func(worldId byte, channelId byte, mapId uint32, environment string) {

	}
}

func KillAllMonsters(l logrus.FieldLogger) func(worldId byte, channelId byte, mapId uint32) {
	return func(worldId byte, channelId byte, mapId uint32) {

	}
}

func SetSummonState(l logrus.FieldLogger) func(worldId byte, channelId byte, mapId uint32, state bool) {
	return func(worldId byte, channelId byte, mapId uint32, state bool) {

	}
}

func SummonState(l logrus.FieldLogger) func(worldId byte, channelId byte, mapId uint32) bool {
	return func(worldId byte, channelId byte, mapId uint32) bool {
		return false
	}
}