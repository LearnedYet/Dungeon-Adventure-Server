package controller

import "github.com/LearnedYet/Dungeon-Adventure-Server/engine"

func (c *GameController) handleLobbyStatusRequest(playerId engine.PlayerId, request *Request, response *Response) []byte {
	response.Data[DataKeyRooms] = c.rooms.GetAllRoomInfos()
	return response.WithStatus(ResponseStatusOk)
}
