package controller

import "github.com/LearnedYet/Dungeon-Adventure-Server/engine"

func (c *GameController) handleLobbyStatusRequest(playerId engine.PlayerId, request *Request, response *Response) string {
	response.Data[DataKeyRooms] = c.rooms.GetAllRoomInfosList()
	response.Data[DataKeyUsersCount] = c.users.TotalCount()
	return response.WithStatus(ResponseStatusOk)
}
