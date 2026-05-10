package data

func NewReward(c []*CurrencyOwned) *Reward {
	return &Reward{
		XP:   0,
		Loot: NewInventory(c),
	}
}

func (r *Reward) AddXP(number float64) {
	r.XP += number
}

func (r *Reward) AddLoot(loot IStorable, number int) {
	r.Loot.Store(loot, number)
}
