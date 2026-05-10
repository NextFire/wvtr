package gamedata

import "wvtrserv/data"

var allCurrencies []*data.Currency = []*data.Currency{
	{
		Storable: data.Storable{
			Name:    "Gold",
			IconURL: DOMAIN_NAME + "/imgs/currencies/gold.ico",
		},
		Type: data.Gold,
	},
	{
		Storable: data.Storable{
			Name:    "Metal scrap",
			IconURL: DOMAIN_NAME + "/imgs/currencies/mscrap.ico",
		},
		Type: data.MScrap,
	},
	{
		Storable: data.Storable{
			Name:    "Cloth scrap",
			IconURL: DOMAIN_NAME + "/imgs/currencies/cscrap.ico",
		},
		Type: data.CScrap,
	},
	{
		Storable: data.Storable{
			Name:    "Lether scrap",
			IconURL: DOMAIN_NAME + "/imgs/currencies/lscrap.ico",
		},
		Type: data.LSCrap,
	},
}

func GetAllCurrencies() []*data.Currency {
	return allCurrencies
}
