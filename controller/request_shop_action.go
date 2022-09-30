package controller

import (
	"github.com/LearnedYet/Dungeon-Adventure-Server/domain"
	"github.com/LearnedYet/Dungeon-Adventure-Server/engine"
)

type ShopActionRequest struct {
	Request
	Data domain.Action `json:"data"`
}

func (c *GameController) handleShopActionRequest(userId engine.UserId, requestRaw string, response *Response) string {
	request := parseRequest(&ShopActionRequest{}, requestRaw)
	if request == nil {
		return response.WithStatus(ResponseStatusMalformed)
	}
	user, ok := c.users.Get(userId)
	if !ok {
		return response.WithStatus(ResponseStatusNotFound)
	}
	actionResult := c.shop.ExecuteAction(request.Data, &user)
	if actionResult.Result == domain.ResultAccomplished {
		response.Data[DataKeyUnit] = user.Unit
	}
	response.Data[DataKeyAction] = request.Data
	response.Data[DataKeyActionResult] = actionResult
	return response.WithStatus(ResponseStatusOk)
}
