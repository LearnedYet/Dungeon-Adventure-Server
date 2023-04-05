package controller

import (
	"github.com/LearnedYet/Dungeon-Adventure-Server/controller/users"
	"github.com/LearnedYet/Dungeon-Adventure-Server/engine"
)

type ApplyForAJobRequestData struct {
	Code engine.PlayerJobCode `json:"code"`
}

func (c *GameController) handleApplyForAJobRequest(playerId engine.PlayerId, request *Request, response *Response) []byte {
	data := parseRequestData(&ApplyForAJobRequestData{}, request.Data)
	if data == nil {
		return response.WithStatus(ResponseStatusMalformed)
	}
	user, ok := c.users.Get(playerId)
	if !ok {
		return response.WithStatus(ResponseStatusNotFound)
	}
	jobStatus, ok := c.employment.ApplyForAJob(&user, data.Code)
	if !ok {
		return response.WithStatus(ResponseStatusNotAllowed)
	}
	if jobStatus != nil {
		c.persistJobStatus(user.Email, *jobStatus)
	}
	c.users.ChangeUserStatus(playerId, users.UserStatusAtJob)
	return response.WithStatus(ResponseStatusOk)
}
