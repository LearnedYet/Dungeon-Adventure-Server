package controller

import (
	"github.com/LearnedYet/Dungeon-Adventure-Server/controller/users"
	"github.com/LearnedYet/Dungeon-Adventure-Server/engine"
)

func (c *GameController) handleExitLobbyRequest(playerId engine.PlayerId, request *Request, response *Response) []byte {
	status := c.users.GetUserStatus(playerId)
	if status != users.UserStatusInLobby {
		return response.WithStatus(ResponseStatusNotAllowed)
	}
	c.users.ChangeUserStatus(playerId, users.UserStatusJoined)
	return response.WithStatus(ResponseStatusOk)
}
