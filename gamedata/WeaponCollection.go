package gamedata

import "wvtrserv/data"

// bases
var SwordBase *data.WeaponBase = &data.WeaponBase{
	BaseDamagesRange: map[data.DamageType]*data.StatsRange{
		data.Slash: {Min: 2, Max: 5},
	},
	BaseAttackSpeed: &data.StatsRange{Min: 1.7, Max: 2},
	BaseCritRate:    &data.StatsRange{Min: 5, Max: 10},
	EquipableBase: data.EquipableBase{
		AffixesPool: []*data.Affix{
			PercentRecupSpeed[1],
			FlatPhysicalDamage[1],
			FlatSlashDamage[1],
			CriticalChance[1],
			CriticalDamage[1],
		},
	},
	StrScaling: 0.50,
	IntScaling: 0.01,
	DexScaling: 0.20,
	LckScaling: 0.02,
}

var HammerBase *data.WeaponBase = &data.WeaponBase{
	BaseDamagesRange: map[data.DamageType]*data.StatsRange{
		data.Blunt: {Min: 5, Max: 10},
	},
	BaseAttackSpeed: &data.StatsRange{Min: 2.5, Max: 4},
	BaseCritRate:    &data.StatsRange{Min: 0, Max: 1},
	EquipableBase: data.EquipableBase{
		AffixesPool: []*data.Affix{
			PercentRecupSpeed[1],
			FlatPhysicalDamage[1],
			FlatBluntDamage[1],
			CriticalChance[1],
			CriticalDamage[1],
		},
	},
	StrScaling: 0.80,
	IntScaling: 0.01,
	DexScaling: 0.05,
	LckScaling: 0.01,
}

var SpearBase *data.WeaponBase = &data.WeaponBase{
	BaseDamagesRange: map[data.DamageType]*data.StatsRange{
		data.Pierce: {Min: 2, Max: 4},
	},
	BaseAttackSpeed: &data.StatsRange{Min: 1.2, Max: 1.9},
	BaseCritRate:    &data.StatsRange{Min: 6, Max: 12},
	EquipableBase: data.EquipableBase{
		AffixesPool: []*data.Affix{
			PercentRecupSpeed[1],
			FlatPhysicalDamage[1],
			FlatPierceDamage[1],
			CriticalChance[1],
			CriticalDamage[1],
		},
	},
	StrScaling: 0.05,
	IntScaling: 0.01,
	DexScaling: 0.80,
	LckScaling: 0.02,
}

var DaggerBase *data.WeaponBase = &data.WeaponBase{
	BaseDamagesRange: map[data.DamageType]*data.StatsRange{
		data.Pierce: {Min: 1, Max: 3},
		data.Slash:  {Min: 1, Max: 3},
	},
	BaseAttackSpeed: &data.StatsRange{Min: 1, Max: 1.6},
	BaseCritRate:    &data.StatsRange{Min: 15, Max: 25},
	EquipableBase: data.EquipableBase{
		AffixesPool: []*data.Affix{
			PercentRecupSpeed[1],
			FlatPhysicalDamage[1],
			FlatPierceDamage[1],
			CriticalChance[1],
			CriticalDamage[1],
		},
	},
	StrScaling: 0.01,
	IntScaling: 0.01,
	DexScaling: 1.00,
	LckScaling: 0.50,
}

var BowBase *data.WeaponBase = &data.WeaponBase{
	BaseDamagesRange: map[data.DamageType]*data.StatsRange{
		data.Pierce: {Min: 2, Max: 6},
	},
	BaseAttackSpeed: &data.StatsRange{Min: 1.2, Max: 1.9},
	BaseCritRate:    &data.StatsRange{Min: 5, Max: 10},
	EquipableBase: data.EquipableBase{
		AffixesPool: []*data.Affix{
			PercentRecupSpeed[1],
			FlatPhysicalDamage[1],
			FlatPierceDamage[1],
			CriticalChance[1],
			CriticalDamage[1],
		},
	},
	StrScaling: 0.05,
	IntScaling: 0.01,
	DexScaling: 0.80,
	LckScaling: 0.10,
}

var FistBase *data.WeaponBase = &data.WeaponBase{
	BaseDamagesRange: map[data.DamageType]*data.StatsRange{
		data.Blunt: {Min: 1, Max: 2},
	},
	BaseAttackSpeed: &data.StatsRange{Min: 1.8, Max: 3},
	BaseCritRate:    &data.StatsRange{Min: 0, Max: 0.5},
	EquipableBase: data.EquipableBase{
		AffixesPool: []*data.Affix{
			PercentRecupSpeed[1],
			FlatPhysicalDamage[1],
			FlatBluntDamage[1],
			CriticalChance[1],
			CriticalDamage[1],
		},
	},
	StrScaling: 0.50,
	IntScaling: 0.50,
	DexScaling: 0.50,
	LckScaling: 0.50,
}

var TuskBase *data.WeaponBase = &data.WeaponBase{
	BaseDamagesRange: map[data.DamageType]*data.StatsRange{
		data.Pierce: {Min: 2, Max: 4},
	},
	BaseAttackSpeed: &data.StatsRange{Min: 2.4, Max: 4},
	BaseCritRate:    &data.StatsRange{Min: 7, Max: 10},
	EquipableBase: data.EquipableBase{
		AffixesPool: []*data.Affix{
			PercentRecupSpeed[1],
			FlatPhysicalDamage[1],
			FlatPierceDamage[1],
			CriticalChance[1],
			CriticalDamage[1],
		},
	},
	StrScaling: 0.50,
	IntScaling: 0.01,
	DexScaling: 0.50,
	LckScaling: 0.05,
}

var ClawBase *data.WeaponBase = &data.WeaponBase{
	BaseDamagesRange: map[data.DamageType]*data.StatsRange{
		data.Slash: {Min: 3, Max: 6},
	},
	BaseAttackSpeed: &data.StatsRange{Min: 1.5, Max: 3},
	BaseCritRate:    &data.StatsRange{Min: 5, Max: 10},
	EquipableBase: data.EquipableBase{
		AffixesPool: []*data.Affix{
			PercentRecupSpeed[1],
			FlatPhysicalDamage[1],
			FlatSlashDamage[1],
			CriticalChance[1],
			CriticalDamage[1],
		},
	},
	StrScaling: 0.20,
	IntScaling: 0.01,
	DexScaling: 1.00,
	LckScaling: 0.10,
}

// Weapon
var GoblinSword *data.Weapon = data.CreateWeapon(SwordBase)
var WolfClaw *data.Weapon = data.CreateWeapon(ClawBase)
var BoarTusk *data.Weapon = data.CreateWeapon(TuskBase)
var SlimeAttack *data.Weapon = data.CreateWeapon(FistBase)
var RabbitHorn *data.Weapon = data.CreateWeapon(TuskBase)
