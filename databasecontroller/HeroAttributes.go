package databasecontroller

import "wvtrserv/data"

func GetHeroAttributesByID(id uint) *data.HeroAttributes {
	var inv *data.HeroAttributes = &data.HeroAttributes{}
	// We can request equipment id later
	db.Find(&inv, id)

	return inv
}
