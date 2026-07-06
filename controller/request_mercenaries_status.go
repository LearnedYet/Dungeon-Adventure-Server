package controller

import "github.com/LearnedYet/Dungeon-Adventure-Server/engine"

func (c *GameController) handleMercenariesStatusRequest(playerId engine.PlayerId, request *Request, response *Response) []byte {
	response.Data[DataKeyMercenaries] = c.mercenaries.GetStatus()
	return response.WithStatus(ResponseStatusOk)
}
