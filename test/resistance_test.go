package test

import (
	"fmt"
	"github.com/LearnedYet/Dungeon-Adventure-Server/domain"
	"github.com/LearnedYet/Dungeon-Adventure-Server/util"
	"testing"
)

func TestAccumulateResistance(t *testing.T) {
	unit := domain.Unit{}
	equipment, ok := util.JsonToObject(
		&[]domain.Weapon{},
		`[{
			"name": "The thing",
			"equipped": true,
			"enhancement": [
				{
					"cutting": 30,
					"crushing": 5,
					"stabbing": 0
				},
				{
					"cold": 5,
					"cutting": -2
				}
			]
		}]`)
	if !ok {
		t.Fatal()
	}
	for _, v := range *equipment.(*[]domain.Weapon) {
		unit.Items = append(unit.Items, v)
	}
	fmt.Printf("Total resistance: {%v}\n", unit.TotalResistance())
}
