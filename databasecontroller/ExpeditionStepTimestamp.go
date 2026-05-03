package databasecontroller

import "wvtrserv/data"

func GetExpeditionStepTimestampByID(id uint) *data.ExpeditionStepTimestamp {
	var res *data.ExpeditionStepTimestamp = &data.ExpeditionStepTimestamp{}
	// We can request equipment id later
	db.Preload("WhatAction").Find(&res, id)

	if res.WhatAction != nil {
		res.WhatAction = GetFieldActionDescByID(res.WhatAction.ID)
	}

	return res
}
