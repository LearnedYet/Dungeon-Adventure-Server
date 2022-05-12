package controller

import "github.com/LearnedYet/Dungeon-Adventure-Server/engine"

type DestroyGameRoomRequest struct {
	Request
}

func (c *GameController) handleDestroyGameRoomRequest(userId engine.UserId, requestRaw string, response *Response) string {
	request := parseRequest(&DestroyGameRoomRequest{}, requestRaw)
	if request == nil {
		return response.WithStatus(ResponseStatusMailformed)
	}
	if !c.rooms.ExistsForHostId(userId) {
		return response.WithStatus(ResponseStatusNotAllowed)
	}
	if _, ok := c.rooms.PopByHostId(userId); !ok {
		return response.WithStatus(ResponseStatusFailed)
	}
	return response.WithStatus(ResponseStatusOk)
}
