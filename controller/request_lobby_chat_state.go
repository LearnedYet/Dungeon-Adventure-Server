package controller

import (
	"github.com/LearnedYet/Dungeon-Adventure-Server/controller/users"
	"github.com/LearnedYet/Dungeon-Adventure-Server/engine"
)

func (c *GameController) handleLobbyChatStateRequest(playerId engine.PlayerId, request *Request, response *Response) []byte {
	status := c.users.GetUserStatus(playerId)
	if !status.Test(users.UserStatusInLobby | users.UserStatusInRoom) {
		return response.WithStatus(ResponseStatusNotAllowed)
	}
	response.Data[DataKeyChat] = c.lobbyChat.State()
	return response.WithStatus(ResponseStatusOk)
}
