package controller

import (
	"github.com/LearnedYet/Dungeon-Adventure-Server/controller/users"
	"github.com/LearnedYet/Dungeon-Adventure-Server/engine"
)

type LeaveGameRoomRequest struct {
	Request
}

func (c *GameController) handleLeaveGameRoomRequest(userId engine.UserId, requestRaw string, response *Response) string {
	request := parseRequest(&LeaveGameRoomRequest{}, requestRaw)
	if request == nil {
		return response.WithStatus(ResponseStatusMailformed)
	}
	roomUid, ok := c.rooms.GetUidByUserId(userId)
	if !ok {
		return response.WithStatus(ResponseStatusNotFound)
	}
	if !c.rooms.RemoveUser(userId) {
		return response.WithStatus(ResponseStatusFailed)
	}
	c.users.ChangeUserStatus(userId, users.UserStatusJoined)
	c.broadcastRoomStatus(roomUid)
	return response.WithStatus(ResponseStatusOk)
}
