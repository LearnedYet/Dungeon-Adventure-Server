package controller

import (
	"github.com/LearnedYet/Dungeon-Adventure-Server/domain"
	"github.com/LearnedYet/Dungeon-Adventure-Server/engine"
)

type ConfiguratorActionRequestData struct {
	domain.Action
}

func (c *GameController) handleConfiguratorActionRequest(playerId engine.PlayerId, request *Request, response *Response) string {
	data := parseRequestData(&ConfiguratorActionRequestData{}, request.Data)
	if data == nil {
		return response.WithStatus(ResponseStatusMalformed)
	}
	user, ok := c.users.Get(playerId)
	if !ok {
		return response.WithStatus(ResponseStatusNotFound)
	}
	actionResult := c.configurator.ExecuteAction(data.Action, &user.Unit.Unit)
	if actionResult.Result == domain.ResultAccomplished {
		response.Data[DataKeyUnit] = user.Unit
	}
	response.Data[DataKeyAction] = request.Data
	response.Data[DataKeyActionResult] = actionResult
	return response.WithStatus(ResponseStatusOk)
}
