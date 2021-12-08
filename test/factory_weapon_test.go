package test

import (
	"github.com/LearnedYet/Dungeon-Adventure-Server/domain"
	"github.com/LearnedYet/Dungeon-Adventure-Server/util"
	"testing"
)

func newBaseSword(t *testing.T) *domain.Weapon {
	weapon, ok := util.JsonToObject(
		&domain.Weapon{},
		`{
			"uid": 4000,
			"type": "weapon",
			"name": "Sword",
			"description": "Base one hand sword",
			"condition": 100,
			"slot": "weapon",
			"slotsNumber": 1,
			"requirements": {
				"strength": 5
			},
			"impact": [
				{
					"chance": 40,
					"cutting": 30,
					"crushing": 5
				},
				{
					"duration": 3,
					"chance": 15,
					"bleeding": 3
				}
			],
			"enhancement": [
				{
					"damage": {
						"cutting": 30
					}
				}	
			]
		}`)
	if !ok {
		t.Fatal()
	}
	return weapon.(*domain.Weapon)
}
