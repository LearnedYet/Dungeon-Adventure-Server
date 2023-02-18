package controller

import "github.com/LearnedYet/Dungeon-Adventure-Server/engine"

func (c *GameController) handleUserStatusRequest(playerId engine.PlayerId, request *Request, response *Response) string {
	user, ok := c.users.Get(playerId)
	if !ok {
		return response.WithStatus(ResponseStatusNotFound)
	}
	response.fillUserStatus(&user)
	return response.WithStatus(ResponseStatusOk)
}
