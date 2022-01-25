package engine

import "github.com/LearnedYet/Dungeon-Adventure-Server/domain"

type EndTurnResult struct {
	Damage   map[uint]domain.Damage       `json:"damage"`
	Recovery map[uint]domain.UnitRecovery `json:"recovery"`
}
