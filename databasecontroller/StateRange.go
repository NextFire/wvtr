package databasecontroller

import "wvtrserv/data"

func GetStatsRangeByID(id uint) *data.StatsRange {
	var rng *data.StatsRange = &data.StatsRange{}
	db.Find(&rng, id)
	return rng
}
