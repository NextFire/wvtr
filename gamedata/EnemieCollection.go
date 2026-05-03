package gamedata

import "wvtrserv/data"

/////////// \\\\\\\\\\\\
/////////// \\\\\\\\\\\\
//////////   \\\\\\\\\\\
/////////  |  \\\\\\\\\\
////////   |   \\\\\\\\\
///////    0    \\\\\\\\
/////////// \\\\\\\\\\\\
/*     Enemies names must be unique    */

const DOMAIN_NAME = "https://tama.rhiobet.sh"

var GreenSlime *data.Hero = &data.Hero{
	ImageUrl: DOMAIN_NAME + "/imgs/enemies/green_slime.jpg",
	Name:     "Green slime",
	Attributes: &data.HeroAttributes{
		Level:        1,
		MaxHP:        10,
		Strength:     2,
		Intelligence: 1,
		Dexterity:    1,
		Luck:         1,
		Slash:        2,
		Lightning:    2,
	},
	Equipment: &data.HeroEquipment{
		Weapon: SlimeAttack,
	},
	WeaponAttack: GetAttackSkill(),
}

var BlueSlime *data.Hero = &data.Hero{
	ImageUrl: DOMAIN_NAME + "/imgs/enemies/blue_slime.jpg",
	Name:     "Blue slime",
	Attributes: &data.HeroAttributes{
		Level:        1,
		MaxHP:        10,
		Strength:     2,
		Intelligence: 1,
		Dexterity:    1,
		Luck:         1,
		Slash:        2,
		Frost:        2,
	},
	Equipment: &data.HeroEquipment{
		Weapon: SlimeAttack,
	},
	WeaponAttack: GetAttackSkill(),
}

var RedSlime *data.Hero = &data.Hero{
	ImageUrl: DOMAIN_NAME + "/imgs/enemies/red_slime.jpg",
	Name:     "Red slime",
	Attributes: &data.HeroAttributes{
		Level:        1,
		MaxHP:        10,
		Strength:     2,
		Intelligence: 1,
		Dexterity:    1,
		Luck:         1,
		Slash:        2,
		Fire:         2,
	},
	Equipment: &data.HeroEquipment{
		Weapon: SlimeAttack,
	},
	WeaponAttack: GetAttackSkill(),
}

var HornRabbit *data.Hero = &data.Hero{
	ImageUrl: DOMAIN_NAME + "/imgs/enemies/horn_rabbit.png",
	Name:     "Horn Rabbit",
	Attributes: &data.HeroAttributes{
		Level:        1,
		MaxHP:        8,
		Strength:     1,
		Intelligence: 1,
		Dexterity:    3,
		Luck:         2,
	},
	Equipment: &data.HeroEquipment{
		Weapon: RabbitHorn,
	},
	WeaponAttack: GetAttackSkill(),
}

var Boar *data.Hero = &data.Hero{
	ImageUrl: DOMAIN_NAME + "/imgs/enemies/boar.png",
	Name:     "Boar",
	Attributes: &data.HeroAttributes{
		Level:        1,
		MaxHP:        15,
		Strength:     3,
		Intelligence: 1,
		Dexterity:    2,
		Luck:         1,
	},
	Equipment: &data.HeroEquipment{
		Weapon: BoarTusk,
	},
	WeaponAttack: GetAttackSkill(),
}

var Wolf *data.Hero = &data.Hero{
	ImageUrl: DOMAIN_NAME + "/imgs/enemies/wolf.png",
	Name:     "Wolf",
	Attributes: &data.HeroAttributes{
		Level:        2,
		MaxHP:        15,
		Strength:     2,
		Intelligence: 2,
		Dexterity:    2,
		Luck:         1,
	},
	Equipment: &data.HeroEquipment{
		Weapon: WolfClaw,
	},
	WeaponAttack: GetAttackSkill(),
}

var Goblin *data.Hero = &data.Hero{
	ImageUrl: DOMAIN_NAME + "/imgs/enemies/goblin.png",
	Name:     "Goblin",
	Attributes: &data.HeroAttributes{
		Level:        2,
		MaxHP:        15,
		Strength:     3,
		Intelligence: 3,
		Dexterity:    2,
		Luck:         1,
	},
	Equipment: &data.HeroEquipment{
		Weapon: GoblinSword,
	},
	WeaponAttack: GetAttackSkill(),
}

var EnemyPool map[int][]*data.Hero = map[int][]*data.Hero{
	1: {
		BlueSlime,
		RedSlime,
		GreenSlime,
		HornRabbit,
		Boar,
	},
	2: {
		Wolf,
		Goblin,
	},
}

var PlainPool []*data.Hero = EnemyPool[1]

var EveryEnemies []*data.Hero = []*data.Hero{
	BlueSlime,
	GreenSlime,
	RedSlime,
	HornRabbit,
	Boar,
	Wolf,
	Goblin,
}

func GetEveryEnemies() []*data.Hero {
	return EveryEnemies
}
