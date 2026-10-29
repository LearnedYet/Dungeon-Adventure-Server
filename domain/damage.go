package domain

import (
	"github.com/LearnedYet/Dungeon-Adventure-Server/util"
)

type Damage struct {
	Stabbing       float32 `json:"stabbing,omitzero" bson:"stabbing,omitempty"`     // 影响生命值
	Cutting        float32 `json:"cutting,omitzero" bson:"cutting,omitempty"`       // 影响生命值
	Crushing       float32 `json:"crushing,omitzero" bson:"crushing,omitempty"`     // 影响生命值
	Fire           float32 `json:"fire,omitzero" bson:"fire,omitempty"`             // 影响生命值
	Cold           float32 `json:"cold,omitzero" bson:"cold,omitempty"`             // 影响生命值
	Lightning      float32 `json:"lightning,omitzero" bson:"lightning,omitempty"`   // 影响生命值
	Acid           float32 `json:"acid,omitzero" bson:"acid,omitempty"`             // 影响生命值
	Poison         float32 `json:"poison,omitzero" bson:"poison,omitempty"`         // 影响生命值
	Exhaustion     float32 `json:"exhaustion,omitzero" bson:"exhaustion,omitempty"` // 影响体力
	ManaDrain      float32 `json:"manaDrain,omitzero" bson:"manaDrain,omitempty"`   // 影响法力
	Bleeding       float32 `json:"bleeding,omitzero" bson:"bleeding,omitempty"`     // 影响生命值
	Fear           float32 `json:"fear,omitzero" bson:"fear,omitempty"`             // 影响压力
	Curse          float32 `json:"curse,omitzero" bson:"curse,omitempty"`           // 影响压力
	Madness        float32 `json:"madness,omitzero" bson:"madness,omitempty"`       // 影响压力
	IsCritical     bool    `json:"isCritical,omitzero" bson:"-"`                    // 暴击伤害标志
	IsCriticalMiss bool    `json:"isCriticalMiss,omitzero" bson:"-"`                // 严重失误伤害标志
	WithStun       bool    `json:"withStun,omitzero" bson:"-"`                      // 眩晕标志
}

func (d *Damage) Accumulate(damage Damage) {
	d.Stabbing += damage.Stabbing
	d.Cutting += damage.Cutting
	d.Crushing += damage.Crushing
	d.Fire += damage.Fire
	d.Cold += damage.Cold
	d.Lightning += damage.Lightning
	d.Acid += damage.Acid
	d.Poison += damage.Poison
	d.Exhaustion += damage.Exhaustion
	d.ManaDrain += damage.ManaDrain
	d.Bleeding += damage.Bleeding
	d.Fear += damage.Fear
	d.Curse += damage.Curse
	d.Madness += damage.Madness
}

func (d *Damage) Reduce(damage Damage) {
	d.Stabbing -= damage.Stabbing
	d.Cutting -= damage.Cutting
	d.Crushing -= damage.Crushing
	d.Fire -= damage.Fire
	d.Cold -= damage.Cold
	d.Lightning -= damage.Lightning
	d.Acid -= damage.Acid
	d.Poison -= damage.Poison
	d.Exhaustion -= damage.Exhaustion
	d.ManaDrain -= damage.ManaDrain
	d.Bleeding -= damage.Bleeding
	d.Fear -= damage.Fear
	d.Curse -= damage.Curse
	d.Madness -= damage.Madness
}

func (d *Damage) PhysicalDamage() float32 {
	return d.Stabbing + d.Cutting + d.Crushing + d.Fire + d.Cold + d.Lightning + d.Acid
}

func (d *Damage) HasPhysicalEffect() bool {
	return d.Stabbing != 0 ||
		d.Cutting != 0 ||
		d.Crushing != 0 ||
		d.Fire != 0 ||
		d.Cold != 0 ||
		d.Lightning != 0 ||
		d.Acid != 0
}

func (d *Damage) HasEffect() bool {
	return d.HasPhysicalEffect() ||
		d.Poison != 0 ||
		d.Exhaustion != 0 ||
		d.ManaDrain != 0 ||
		d.Bleeding != 0 ||
		d.Fear != 0 ||
		d.Curse != 0 ||
		d.Madness != 0
}

func (d *Damage) Enchance(attributes UnitAttributes, damage Damage) {
	d.Stabbing = util.AccumulateIfNotZeros(d.Stabbing, attributes.Strength)
	d.Cutting = util.AccumulateIfNotZeros(d.Cutting, attributes.Strength)
	d.Crushing = util.AccumulateIfNotZeros(d.Crushing, attributes.Strength)
	d.Fire = util.AccumulateIfNotZeros(d.Fire, attributes.Intelligence)
	d.Cold = util.AccumulateIfNotZeros(d.Cold, attributes.Intelligence)
	d.Lightning = util.AccumulateIfNotZeros(d.Lightning, attributes.Intelligence)
	d.Acid = util.AccumulateIfNotZeros(d.Acid, attributes.Intelligence)
	d.Exhaustion = util.AccumulateIfNotZeros(d.Exhaustion, attributes.Intelligence)
	d.ManaDrain = util.AccumulateIfNotZeros(d.ManaDrain, attributes.Intelligence)
	d.Bleeding = util.AccumulateIfNotZeros(d.Bleeding, attributes.Strength)
	d.Fear = util.AccumulateIfNotZeros(d.Fear, attributes.Intelligence)
	d.Curse = util.AccumulateIfNotZeros(d.Curse, attributes.Intelligence)
	d.Madness = util.AccumulateIfNotZeros(d.Madness, attributes.Intelligence)

	d.Stabbing = util.AccumulateIfNotZeros(d.Stabbing, damage.Stabbing)
	d.Cutting = util.AccumulateIfNotZeros(d.Cutting, damage.Cutting)
	d.Crushing = util.AccumulateIfNotZeros(d.Crushing, damage.Crushing)
	d.Fire = util.AccumulateIfNotZeros(d.Fire, damage.Fire)
	d.Cold = util.AccumulateIfNotZeros(d.Cold, damage.Cold)
	d.Lightning = util.AccumulateIfNotZeros(d.Lightning, damage.Lightning)
	d.Acid = util.AccumulateIfNotZeros(d.Acid, damage.Acid)
	d.Poison = util.AccumulateIfNotZeros(d.Poison, damage.Poison)
	d.Exhaustion = util.AccumulateIfNotZeros(d.Exhaustion, damage.Exhaustion)
	d.ManaDrain = util.AccumulateIfNotZeros(d.ManaDrain, damage.ManaDrain)
	d.Bleeding = util.AccumulateIfNotZeros(d.Bleeding, damage.Bleeding)
	d.Fear = util.AccumulateIfNotZeros(d.Fear, damage.Fear)
	d.Curse = util.AccumulateIfNotZeros(d.Curse, damage.Curse)
	d.Madness = util.AccumulateIfNotZeros(d.Madness, damage.Madness)
}

func (d *Damage) MultiplyAll(factor float32) {
	d.Stabbing = util.MultiplyWithRounding(d.Stabbing, factor)
	d.Cutting = util.MultiplyWithRounding(d.Cutting, factor)
	d.Crushing = util.MultiplyWithRounding(d.Crushing, factor)
	d.Fire = util.MultiplyWithRounding(d.Fire, factor)
	d.Cold = util.MultiplyWithRounding(d.Cold, factor)
	d.Lightning = util.MultiplyWithRounding(d.Lightning, factor)
	d.Acid = util.MultiplyWithRounding(d.Acid, factor)
	d.Poison = util.MultiplyWithRounding(d.Poison, factor)
	d.Exhaustion = util.MultiplyWithRounding(d.Exhaustion, factor)
	d.ManaDrain = util.MultiplyWithRounding(d.ManaDrain, factor)
	d.Bleeding = util.MultiplyWithRounding(d.Bleeding, factor)
	d.Fear = util.MultiplyWithRounding(d.Fear, factor)
	d.Curse = util.MultiplyWithRounding(d.Curse, factor)
	d.Madness = util.MultiplyWithRounding(d.Madness, factor)
}

func (d *Damage) EnchanceAll(value float32) {
	d.Stabbing = util.AccumulateIfNotZeros(d.Stabbing, value)
	d.Cutting = util.AccumulateIfNotZeros(d.Cutting, value)
	d.Crushing = util.AccumulateIfNotZeros(d.Crushing, value)
	d.Fire = util.AccumulateIfNotZeros(d.Fire, value)
	d.Cold = util.AccumulateIfNotZeros(d.Cold, value)
	d.Lightning = util.AccumulateIfNotZeros(d.Lightning, value)
	d.Acid = util.AccumulateIfNotZeros(d.Acid, value)
	d.Poison = util.AccumulateIfNotZeros(d.Poison, value)
	d.Exhaustion = util.AccumulateIfNotZeros(d.Exhaustion, value)
	d.ManaDrain = util.AccumulateIfNotZeros(d.ManaDrain, value)
	d.Bleeding = util.AccumulateIfNotZeros(d.Bleeding, value)
	d.Fear = util.AccumulateIfNotZeros(d.Fear, value)
	d.Curse = util.AccumulateIfNotZeros(d.Curse, value)
	d.Madness = util.AccumulateIfNotZeros(d.Madness, value)
}

func (d *Damage) Normalize() {
	d.Stabbing = max(d.Stabbing, 0)
	d.Cutting = max(d.Cutting, 0)
	d.Crushing = max(d.Crushing, 0)
	d.Fire = max(d.Fire, 0)
	d.Cold = max(d.Cold, 0)
	d.Lightning = max(d.Lightning, 0)
	d.Acid = max(d.Acid, 0)
	d.Poison = max(d.Poison, 0)
	d.Exhaustion = max(d.Exhaustion, 0)
	d.ManaDrain = max(d.ManaDrain, 0)
	d.Bleeding = max(d.Bleeding, 0)
	d.Fear = max(d.Fear, 0)
	d.Curse = max(d.Curse, 0)
	d.Madness = max(d.Madness, 0)
}

func (d *Damage) Apply(state *UnitState) {
	state.Health -= d.Stabbing + d.Cutting + d.Crushing + d.Fire + d.Cold + d.Lightning + d.Acid + d.Poison + d.Bleeding
	state.Stamina -= d.Exhaustion
	state.Mana -= d.ManaDrain
	state.Stress += d.Fear + d.Curse + d.Madness

	state.Health = max(state.Health, 0)
	state.Stamina = max(state.Stamina, 0)
	state.Mana = max(state.Mana, 0)
}

func (d Damage) IsZero() bool {
	return d.Stabbing == 0 &&
		d.Cutting == 0 &&
		d.Crushing == 0 &&
		d.Fire == 0 &&
		d.Cold == 0 &&
		d.Lightning == 0 &&
		d.Acid == 0 &&
		d.Poison == 0 &&
		d.Exhaustion == 0 &&
		d.ManaDrain == 0 &&
		d.Bleeding == 0 &&
		d.Fear == 0 &&
		d.Curse == 0 &&
		d.Madness == 0
}
