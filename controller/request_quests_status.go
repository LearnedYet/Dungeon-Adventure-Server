package controller

import "github.com/LearnedYet/Dungeon-Adventure-Server/engine"

func (c *GameController) handleQuestsStatusRequest(playerId engine.PlayerId, request *Request, response *Response) []byte {
	user, ok := c.users.Get(playerId)
	if !ok {
		return response.WithStatus(ResponseStatusNotFound)
	}
	response.Data[DataKeyQuests] = c.quests.GetStatus(&user.Unit.Unit)
	return response.WithStatus(ResponseStatusOk)
}
