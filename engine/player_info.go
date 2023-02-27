package engine

import "github.com/LearnedYet/Dungeon-Adventure-Server/domain"

type PlayerId string

const (
	PlayerIdEmpty PlayerId = ""
)

type PlayerInfo struct {
	Id        PlayerId         `json:"playerId"`
	Nickname  string           `json:"nickname"`
	Class     domain.UnitClass `json:"class"`
	Level     uint             `json:"level"`
	UnitUid   uint             `json:"unitUid,omitempty"`
	IsOffline bool             `json:"isOffline,omitempty"`
	IsReady   bool             `json:"isReady,omitempty"`
}

func (p PlayerInfo) Clone() *PlayerInfo {
	return &p
}
