package databasecontroller

import "wvtrserv/data"

func GetCurrencyOwnedByID(id uint) *data.CurrencyOwned {
	var res *data.CurrencyOwned = &data.CurrencyOwned{}
	// We can request equipment id later
	db.Preload("Currency").
		Find(&res, id)

	if res.Currency != nil {
		res.Currency = GetCurrencyByID(res.Currency.ID)
	}

	return res
}
