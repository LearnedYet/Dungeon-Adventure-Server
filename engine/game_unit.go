package engine

import "github.com/LearnedYet/Dungeon-Adventure-Server/domain"

type GameUnit struct {
	domain.Unit
	FractionId uint `json:"fractionId"`
}
