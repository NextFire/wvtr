package databasecontroller

import (
	"time"
	"wvtrserv/data"
)

func UpdateTeam(team *data.Team) {
	db.Model(&team).
		Association("Heroes").
		Replace(team.Heroes)
}

func SaveTeam(team *data.Team) {
	for _, h := range team.Heroes {
		SaveHero(h)
	}
}

func GetTeam(team *data.Team) *data.Team {
	db.Preload("Heroes").
		Find(&team)
	return team
}

func GetTeamByID(id uint) *data.Team {
	var res *data.Team
	db.Preload("Heroes").
		Find(&res, id)

	for i := range res.Heroes {
		res.Heroes[i] = GetHeroByID(res.Heroes[i].ID)
	}

	return res
}

func UpdateTeamWithExpAndTime(team *data.Team, expdb data.ExpeditionDB, t time.Time) {
	// get array of all esri to apply between last update and t
	timeStart := team.UpdatedAt
	timeEnd := t
	timeIntervalStartIdx := GetCurrentExpeditionStepIdx(expdb, &timeStart)
	timeIntervalEnd := GetCurrentExpeditionStepIdx(expdb, &timeEnd)
	if timeIntervalEnd < len(expdb.WhatHappened) {
		timeIntervalEnd++
	}

	toApply := expdb.WhatHappened[timeIntervalStartIdx:timeIntervalEnd]

	for _, a := range toApply {
		team = team.ApplyESRI(a, timeStart, timeEnd)
	}

	SaveTeam(team)
}
