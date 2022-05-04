package controller

import (
	"github.com/LearnedYet/Dungeon-Adventure-Server/engine"
	"github.com/LearnedYet/Dungeon-Adventure-Server/util"
	"sync"
)

type GameEngines struct {
	sync.RWMutex
	rndGen  *util.RndGen
	engines []*engine.GameEngine
}

func NewGameEngines() *GameEngines {
	e := &GameEngines{}
	e.rndGen = util.NewRndGen()
	return e
}

func (e *GameEngines) Add(userUids []uint, engine *engine.GameEngine) {

}
