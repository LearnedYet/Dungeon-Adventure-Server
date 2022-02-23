package test

import (
	"fmt"
	"github.com/LearnedYet/Dungeon-Adventure-Server/domain"
	"github.com/LearnedYet/Dungeon-Adventure-Server/util"
	"testing"
)

func TestAccumulateResistance(t *testing.T) {
	unit := newAgileHero(t)
	equipment, err := util.JsonToObject(
		&[]domain.Weapon{},
		`[
			{
			  "name": "The thing",
			  "equipped": true,
			  "modification": [
				{
				  "damage": {
					"cutting": 30,
					"crushing": 5,
					"stabbing": 0
				  },
				  "resistance": {
					"cold": 5,
					"cutting": -2
				  }
				}
			  ]
			}
		]`)
	if err != nil {
		t.Fatal()
	}
	unit.Inventory.Weapon = append(unit.Inventory.Weapon, *equipment.(*[]domain.Weapon)...)
	fmt.Printf("Total modification: {%v}\n", unit.TotalModification())
}
