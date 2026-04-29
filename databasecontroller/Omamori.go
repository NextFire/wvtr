package databasecontroller

import "wvtrserv/data"

func GetOmamoriByID(id uint) *data.Omamori {
	var oma *data.Omamori = &data.Omamori{}
	// Weapon Without States make no sens
	db.Preload("Affixes").
		Find(&oma, id)
	for i := range oma.Affixes {
		oma.Affixes[i] = GetAffixByID(oma.Affixes[i].ID)
	}
	return oma
}
