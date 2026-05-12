package databasecontroller

import "wvtrserv/data"

func GetCurrencyOwnedByID(id uint) *data.CurrencyOwned {
	var res *data.CurrencyOwned = &data.CurrencyOwned{}
	db.Preload("Currency").
		Find(&res, id)

	if res.Currency != nil {
		res.Currency = GetCurrencyByID(res.Currency.ID)
	}

	return res
}

func SaveCurrencyOwned(c *data.CurrencyOwned) {
	db.Save(c)
}
