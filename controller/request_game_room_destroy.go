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
	room, ok := c.rooms.PopByHostId(userId)
	if !ok {
		return response.WithStatus(ResponseStatusFailed)
	}
	for _, actor := range room.GetActors() {
		c.users.ChangeUserStatus(actor.UserId, UserStatusJoined)
	}
	c.broadcastLobbyStatusToAllExcept(userId)
	return response.WithStatus(ResponseStatusOk)
}
