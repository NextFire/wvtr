package databasecontroller

import "wvtrserv/data"

func GetArmorByID(id uint) *data.Armor {
	var armor *data.Armor = &data.Armor{}
	db.Preload("BlockScore").
		Preload("EvadeScore").
		Preload("BaseResistancesRange").
		Preload("Affixes").
		Find(&armor, id)

	for i := range armor.Affixes {
		armor.Affixes[i] = GetAffixByID(armor.Affixes[i].ID)
	}
	return armor
}

func SaveArmor(o *data.Armor) {
	db.Save(o)
}
