package controller

import (
	"github.com/LearnedYet/Dungeon-Adventure-Server/engine"
)

func (c *GameController) handleGameChatStateRequest(playerId engine.PlayerId, request *Request, response *Response) []byte {
	wrapper, ok := c.engines.Find(playerId)
	if !ok {
		return response.WithStatus(ResponseStatusNotFound)
	}
	response.Data[DataKeyChat] = wrapper.ChatState()
	return response.WithStatus(ResponseStatusOk)
}
