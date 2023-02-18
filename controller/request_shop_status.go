package controller

import "github.com/LearnedYet/Dungeon-Adventure-Server/engine"

func (c *GameController) handleShopStatusRequest(playerId engine.PlayerId, request *Request, response *Response) string {
	response.Data[DataKeyShop] = c.shop.GetStatus()
	return response.WithStatus(ResponseStatusOk)
}
