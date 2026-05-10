package data

func (u User) UserHasAProblem() bool {
	return u.State.State == Error
}

func (u User) UserIsHome() bool {
	return u.State.State == Home
}

func (u *User) GetReward(expReward *Reward) {
	xpGained := expReward.XP / float64(len(u.CurrentTeam.Heroes))
	for _, h := range u.CurrentTeam.Heroes {
		h.GainXP(xpGained)
	}
	u.Inventory.Merge(expReward.Loot)
}
