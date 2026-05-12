package databasecontroller

import (
	"wvtrserv/data"
)

func CreateCurrency(currency *data.Currency) {
	db.Create(currency)
}

func GetAllCurrencies() []*data.Currency {
	res := []*data.Currency{}
	db.Find(&res)
	return res
}

func GetCurrencyByID(id uint) *data.Currency {
	var res *data.Currency = &data.Currency{}
	db.Find(&res, id)

	return res
}
