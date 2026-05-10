package data

type CurrencyType int

const (
	Gold CurrencyType = iota
	MScrap
	LSCrap
	CScrap
	MaxCType
)

func NewCurrencyOwned(allCurrencies []*Currency) []*CurrencyOwned {

	res := make([]*CurrencyOwned, len(allCurrencies))
	for i, c := range allCurrencies {
		res[i] = &CurrencyOwned{
			NumberOwned: 0,
			Currency:    c,
		}
	}

	return res
}
