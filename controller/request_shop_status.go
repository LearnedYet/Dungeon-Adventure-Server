package controller

import "github.com/LearnedYet/Dungeon-Adventure-Server/engine"

func (c *GameController) handleShopStatusRequest(userId engine.UserId, request *Request, response *Response) string {
	response.Data[DataKeyShop] = c.shop.GetStatus()
	return response.WithStatus(ResponseStatusOk)
}
