package databasecontroller

import "wvtrserv/data"

func GetWeaponByID(id uint) *data.Weapon {
	var weap *data.Weapon = &data.Weapon{}
	db.Preload("BaseDamage").
		Preload("BaseCritRate").
		Preload("BaseAttackSpeed").
		Preload("Affixes").
		Find(&weap, id)
	for i := range weap.Affixes {
		weap.Affixes[i] = GetAffixByID(weap.Affixes[i].ID)
	}
	return weap
}

func SaveWeapon(o *data.Weapon) {
	db.Save(o)
}
