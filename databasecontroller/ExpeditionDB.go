package databasecontroller

import (
	"time"
	"wvtrserv/data"
)

func CreateExpeditionDB(edb *data.ExpeditionDB) {
	db.Save(edb)
}

func GetExpeditionDBByID(id uint) *data.ExpeditionDB {
	var ex *data.ExpeditionDB = &data.ExpeditionDB{}
	// We can request equipment id later
	db.Preload("WhatHappened").
		Find(&ex, id)

	for i := range ex.WhatHappened {
		ex.WhatHappened[i] = GetExpeditionStepResolveInfoByID(ex.WhatHappened[i].ID)
	}
	return ex
}

func GetCurrentExpedition(exp *data.ExpeditionDB) *data.ExpeditionDB {
	db.Preload("WhatHappened").
		Find(&exp)
	return exp
}

func GetCurrentExpeditionStepIdx(e data.ExpeditionDB, t *time.Time) int {
	for i, step := range e.WhatHappened {
		if step.Timeline[len(step.Timeline)-1].When.After(*t) {
			return i
		}
	}
	return len(e.WhatHappened)
}
