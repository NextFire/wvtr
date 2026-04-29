package databasecontroller

import "wvtrserv/data"

func GetExpeditionStepResolveInfoByID(id uint) *data.ExpeditionStepResolveInfo {
	var ex *data.ExpeditionStepResolveInfo = &data.ExpeditionStepResolveInfo{}
	// We can request equipment id later
	db.Preload("Timeline").
		Preload("ETeam").
		Find(&ex, id)

	if ex.ETeam != nil {
		ex.ETeam = GetTeamByID(ex.ETeam.ID)
	}
	for i := range ex.Timeline {
		ex.Timeline[i] = GetExpeditionStepTimestampByID(ex.Timeline[i].ID)
	}
	return ex
}
