package test

import (
	"fmt"
	"github.com/LearnedYet/Dungeon-Adventure-Server/controller"
	"testing"
)

func TestMake(t *testing.T) {
	controller := controller.NewGameController()
	result := controller.HandleRequest(`{
		"userId": "user-id-1",
		"type": "createBattleRoom",
		"data": {
			"allowedUsers": ["user-id-1", "user-id-2"]
		}
	}`)
	fmt.Println(result)
}
