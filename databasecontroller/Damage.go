package databasecontroller

import "wvtrserv/data"

func GetDamageByID(id uint) *data.Damage {
	var dmg *data.Damage = &data.Damage{}
	db.Find(&dmg, id)
	return dmg
}
