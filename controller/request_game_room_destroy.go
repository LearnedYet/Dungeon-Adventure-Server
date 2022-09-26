package controller

import (
	"github.com/LearnedYet/Dungeon-Adventure-Server/controller/users"
	"github.com/LearnedYet/Dungeon-Adventure-Server/engine"
)

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
	userIds := room.GetUserIds()
	for _, userId := range userIds {
		c.users.ChangeUserStatus(userId, users.UserStatusJoined)
	}
	c.broadcastRoomStatus(room.Uid)
	c.broadcastUsersStatus(userIds)
	return response.WithStatus(ResponseStatusOk)
}
