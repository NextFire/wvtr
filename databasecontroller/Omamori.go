package databasecontroller

import "wvtrserv/data"

func GetOmamoriByID(id uint) *data.Omamori {
	var oma *data.Omamori = &data.Omamori{}
	db.Preload("Affixes").
		Find(&oma, id)
	for i := range oma.Affixes {
		oma.Affixes[i] = GetAffixByID(oma.Affixes[i].ID)
	}
	return oma
}

func SaveOmamori(o *data.Omamori) {
	db.Save(o)
}
