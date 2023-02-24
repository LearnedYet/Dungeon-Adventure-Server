package model

import (
	"github.com/LearnedYet/Dungeon-Adventure-Server/domain"
	"github.com/LearnedYet/Dungeon-Adventure-Server/engine"
)

type UserModel struct {
	Model
	Email    string               `bson:"email"`
	Class    engine.GameUnitClass `json:"class,omitempty"`
	Nickname string               `bson:"nickname,omitempty"`
	Picture  string               `bson:"picture,omitempty"`
	Unit     *domain.Unit         `bson:"unit,omitempty"`
}
