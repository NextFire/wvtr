package databasecontroller

import "wvtrserv/data"

func GetInventoryByID(id uint) *data.Inventory {
	var inv *data.Inventory = &data.Inventory{}
	// We can request equipment id later
	db.Preload("Storage").
		Find(&inv, id)

	return inv
}
