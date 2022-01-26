package controller

import (
	"github.com/LearnedYet/Dungeon-Adventure-Server/engine"
	"github.com/LearnedYet/Dungeon-Adventure-Server/util"
)

type GameController struct {
	uidGen           util.UidGen
	engines          map[uint]*engine.GameEngine
	userIdToEngineId map[string]uint
}

func (c *GameController) HandleRequest(requestRaw string) string {
	request := c.parseRequest(requestRaw)
	if request == nil {
		return ""
	}
	response := Response{}
	switch request.Type {
	case RequestCreateGameEngine:
		c.createGameEngine(request, &response)
	}
	return util.ObjectToJson(response)
}

func (c *GameController) parseRequest(requestRaw string) *Request {
	if request, ok := util.JsonToObject(&Request{}, requestRaw); ok {
		return request.(*Request)
	}
	return nil
}
