package databasecontroller

import "wvtrserv/data"

func SaveHeroAttributes(a *data.HeroAttributes) {
	db.Save(a)
}

func GetHeroAttributesByID(id uint) *data.HeroAttributes {
	var inv *data.HeroAttributes = &data.HeroAttributes{}
	db.Find(&inv, id)

	return inv
}
