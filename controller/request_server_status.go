package controller

import "github.com/LearnedYet/Dungeon-Adventure-Server/engine"

func (c *GameController) handleServerStatusRequest(playerId engine.PlayerId, request *Request, response *Response) []byte {
	response.Data[DataKeyUsersNumber] = c.users.Total()
	return response.WithStatus(ResponseStatusOk)
}
