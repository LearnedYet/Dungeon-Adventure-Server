package domain

import (
	"github.com/LearnedYet/Dungeon-Adventure-Server/util"
)

type UnitAttributes struct {
	Strength     float32 `json:"strength" bson:"strength"`         // 增强穿刺、斩击、钝击和流血伤害
	Physique     float32 `json:"physique" bson:"physique"`         // 影响眩晕概率
	Agility      float32 `json:"agility" bson:"agility"`           // 影响攻击命中率与闪避概率
	Endurance    float32 `json:"endurance" bson:"endurance"`       // 体力恢复
	Intelligence float32 `json:"intelligence" bson:"intelligence"` // 增强火焰、寒冷、闪电、酸蚀、体力消耗、法力汲取、恐惧和诅咒伤害
	Initiative   float32 `json:"initiative" bson:"initiative"`     // 影响行动顺序
	Luck         float32 `json:"luck" bson:"luck"`                 // 影响暴击概率
}

func (a *UnitAttributes) Accumulate(attributes UnitAttributes) {
	a.Strength += attributes.Strength
	a.Physique += attributes.Physique
	a.Agility += attributes.Agility
	a.Endurance += attributes.Endurance
	a.Intelligence += attributes.Intelligence
	a.Initiative += attributes.Initiative
	a.Luck += attributes.Luck
}

func (a *UnitAttributes) Normalize() {
	a.Strength = max(a.Strength, 0)
	a.Physique = max(a.Physique, 0)
	a.Agility = max(a.Agility, 0)
	a.Endurance = max(a.Endurance, 0)
	a.Intelligence = max(a.Intelligence, 0)
	a.Initiative = max(a.Initiative, 0)
	a.Luck = max(a.Luck, 0)
}

func (a *UnitAttributes) MultiplyAll(factor float32) {
	a.Strength = util.MultiplyWithRounding(a.Strength, factor)
	a.Physique = util.MultiplyWithRounding(a.Physique, factor)
	a.Agility = util.MultiplyWithRounding(a.Agility, factor)
	a.Endurance = util.MultiplyWithRounding(a.Endurance, factor)
	a.Intelligence = util.MultiplyWithRounding(a.Intelligence, factor)
	a.Initiative = util.MultiplyWithRounding(a.Initiative, factor)
	a.Luck = util.MultiplyWithRounding(a.Luck, factor)
}

func (a *UnitAttributes) EnchanceAll(value float32) {
	a.Strength = util.AccumulateIfNotZeros(a.Strength, value)
	a.Physique = util.AccumulateIfNotZeros(a.Physique, value)
	a.Agility = util.AccumulateIfNotZeros(a.Agility, value)
	a.Endurance = util.AccumulateIfNotZeros(a.Endurance, value)
	a.Intelligence = util.AccumulateIfNotZeros(a.Intelligence, value)
	a.Initiative = util.AccumulateIfNotZeros(a.Initiative, value)
	a.Luck = util.AccumulateIfNotZeros(a.Luck, value)
}

func (a UnitAttributes) IsZero() bool {
	return a.Strength == 0 &&
		a.Physique == 0 &&
		a.Agility == 0 &&
		a.Endurance == 0 &&
		a.Intelligence == 0 &&
		a.Initiative == 0 &&
		a.Luck == 0
}
