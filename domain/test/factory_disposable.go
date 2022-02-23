package test

import (
	"github.com/LearnedYet/Dungeon-Adventure-Server/domain"
	"github.com/LearnedYet/Dungeon-Adventure-Server/util"
	"testing"
)

func newSmallHealthPotion(t *testing.T) *domain.Disposable {
	disposable, err := util.JsonToObject(
		&domain.Disposable{},
		`{
			"uid": 2000,
			"type": "disposable",
			"name": "Small Health Potion",
			"description": "Small health potion",
			"quantity": 1,
			"modification": [
				{
					"recovery": {
						"health": 35
					}
				},
				{
					"chance": 50,
					"duration": 5,
					"recovery": {
						"health": 2
					}
				}
			]
		}`)
	if err != nil {
		t.Fatal()
	}
	return disposable.(*domain.Disposable)
}
