package engine

import (
	"fmt"
	"github.com/LearnedYet/Dungeon-Adventure-Server/domain"
)

type GameUnit struct {
	domain.Unit
	FractionId uint   `json:"fractionId"`
	UserId     string `json:"userId,omitempty"`
}

func (u GameUnit) String() string {
	return fmt.Sprintf(
		"%v, fraction: %d, user id: %s",
		u.Unit,
		u.FractionId,
		u.UserId,
	)
}
