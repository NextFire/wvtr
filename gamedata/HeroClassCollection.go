package gamedata

import "wvtrserv/data"

// should be inserted in db once
var heroClasses []*data.HeroClass = []*data.HeroClass{
	{
		Name:         "Warrior",
		Identifier:   data.Warrior,
		Descritpion:  "Relying mostly on strength",
		ClassIconURL: "",
		Weight:       100,
		//12 points (5hp = 1pt) base hp = 10
		MaxHP:        20,
		Strength:     5,
		Intelligence: 1,
		Dexterity:    2,
		Luck:         2,
		// 14 points
		HPgt: 0.6,
		Sgt:  0.5,
		Igt:  0.05,
		Dgt:  0.15,
		Lgt:  0.1,
	},
	{
		Name:         "Mage",
		Identifier:   data.Mage,
		Descritpion:  "Relying mostly on intelligence",
		ClassIconURL: "",
		Weight:       100,
		//12 points (5hp = 1pt) base hp = 10
		MaxHP:        10,
		Strength:     1,
		Intelligence: 6,
		Dexterity:    2,
		Luck:         3,
		// 14 points
		HPgt: 0.45,
		Sgt:  0.05,
		Igt:  0.6,
		Dgt:  0.1,
		Lgt:  0.2,
	},
	{
		Name:         "Thief",
		Identifier:   data.Thief,
		Descritpion:  "Relying mostly on dexterity",
		ClassIconURL: "",
		Weight:       100,
		//12 points (5hp = 1pt) base hp = 10
		MaxHP:        15,
		Strength:     1,
		Intelligence: 2,
		Dexterity:    6,
		Luck:         2,
		// 14 points
		HPgt: 0.5,
		Sgt:  0.1,
		Igt:  0.15,
		Dgt:  0.5,
		Lgt:  0.15,
	},
	{
		Name:         "Jester",
		Identifier:   data.Jester,
		Descritpion:  "He is lucky I guess...",
		ClassIconURL: "",
		Weight:       10,
		//12 points (5hp = 1pt) base hp = 10
		MaxHP:        10,
		Strength:     1,
		Intelligence: 2,
		Dexterity:    2,
		Luck:         7,
		// 14 points
		HPgt: 0.4,
		Sgt:  0.05,
		Igt:  0.15,
		Dgt:  0.15,
		Lgt:  0.65,
	},
}

func GetHeroClasses() []*data.HeroClass {
	return heroClasses
}
