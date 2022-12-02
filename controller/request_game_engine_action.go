package controller

import (
	"github.com/LearnedYet/Dungeon-Adventure-Server/domain"
	"github.com/LearnedYet/Dungeon-Adventure-Server/engine"
)

type GameActionRequestData struct {
	domain.Action
}

func (c *GameController) handleGameActionRequest(userId engine.UserId, request *Request, response *Response) string {
	data := parseRequestData(&GameActionRequestData{}, request.Data)
	if data == nil {
		return response.WithStatus(ResponseStatusMalformed)
	}
	wrapper, ok := c.engines.Find(userId)
	if !ok {
		return response.WithStatus(ResponseStatusNotFound)
	}
	defer wrapper.Unlock()
	wrapper.Lock()
	result, broadcastUserIds, ok := wrapper.ExecuteUserAction(data.Action, userId)
	if !ok {
		return response.WithStatus(ResponseStatusNotAllowed)
	}
	response.Data[DataKeyAction] = request.Data
	response.Data[DataKeyActionResult] = result
	if len(broadcastUserIds) > 0 {
		c.broadcastGameAction(broadcastUserIds, result)
	}
	return response.WithStatus(ResponseStatusOk)
}
