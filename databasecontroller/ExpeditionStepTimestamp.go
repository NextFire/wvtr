package databasecontroller

import "wvtrserv/data"

func GetExpeditionStepTimestampByID(id uint) *data.ExpeditionStepTimestamp {
	var inv *data.ExpeditionStepTimestamp = &data.ExpeditionStepTimestamp{}
	// We can request equipment id later
	db.Find(&inv, id)

	return inv
}
