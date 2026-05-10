package databasecontroller

import "wvtrserv/data"

func GetInventoryByID(id uint) *data.Inventory {
	var res *data.Inventory = &data.Inventory{}
	// We can request equipment id later
	db.Preload("Weapons").
		Preload("Armors").
		Preload("Omamoris").
		Preload("Currencies").
		Find(&res, id)

	for i := range res.Weapons {
		res.Weapons[i] = GetWeaponByID(res.Weapons[i].ID)
	}

	for i := range res.Armors {
		res.Armors[i] = GetArmorByID(res.Armors[i].ID)
	}

	for i := range res.Omamoris {
		res.Omamoris[i] = GetOmamoriByID(res.Omamoris[i].ID)
	}

	for i := range res.Currencies {
		res.Currencies[i] = GetCurrencyOwnedByID(res.Currencies[i].ID)
	}

	return res
}
