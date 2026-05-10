package databasecontroller

import "wvtrserv/data"

func GetGetRewardByID(id uint) *data.Reward {
	var res *data.Reward = nil

	db.Preload("Loot").
		Find(&res, id)

	if res.Loot != nil {
		res.Loot = GetInventoryByID(res.Loot.ID)
	}

	return res
}
