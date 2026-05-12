package databasecontroller

import "wvtrserv/data"

func GetAffixByID(id uint) *data.Affix {
	var aff *data.Affix = &data.Affix{}
	db.Preload("Ranges").Find(&aff, id)
	return aff
}
