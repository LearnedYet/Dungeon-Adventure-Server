package controller

import (
	"github.com/LearnedYet/Dungeon-Adventure-Server/controller/factory"
	"github.com/LearnedYet/Dungeon-Adventure-Server/controller/users"
	"github.com/LearnedYet/Dungeon-Adventure-Server/engine"
)

func (c *GameController) handleStartGameRequest(userId engine.UserId, request *Request, response *Response) string {
	room, ok := c.rooms.PopByHostId(userId)
	if !ok {
		return response.WithStatus(ResponseStatusNotAllowed)
	}
	// todo: handle scenario
	scenario := factory.NewTestScenario()
	actors := room.GetActors()
	engine := engine.NewGameEngine(scenario, actors)
	state := engine.NewGameEvent()
	userIds := engine.GetUserIds()
	for _, userId := range userIds {
		c.users.ChangeUserStatus(userId, users.UserStatusInGame)
	}
	c.broadcastGameState(userIds, state)
	c.broadcastRoomStatus(room.Uid)
	c.broadcastUsersStatus(userIds)
	c.engines.Add(userId, engine)
	return response.WithStatus(ResponseStatusOk)
}
