package databasecontroller

import (
	"time"
	"wvtrserv/data"
	"wvtrserv/logger"
)

func GetGameState(gameState *data.GameState) *data.GameState {
	db.Preload("CurrentExpedition").
		Find(&gameState)

	gameState.CurrentExpedition = GetCurrentExpedition(gameState.CurrentExpedition)

	return gameState
}

func GetGameStateByID(id uint) *data.GameState {
	var gs *data.GameState = &data.GameState{}
	// We can request equipment id later
	db.Preload("CurrentExpedition").
		Find(&gs, id)

	if gs.CurrentExpedition != nil {
		gs.CurrentExpedition = GetExpeditionDBByID(gs.CurrentExpedition.ID)
	}
	return gs
}

func UpdateGameState(state *data.GameState) {
	db.Save(state)
}

func UpdateGameStateWithTime(g *data.GameState, t *time.Time) *data.ExpeditionStepResolveInfo {
	idx := -1
	var currentExpStep *data.ExpeditionStepResolveInfo = nil
	if g.CurrentExpedition != nil {
		idx = GetCurrentExpeditionStepIdx(*g.CurrentExpedition, t)
		currentExpStep = UpdateGameStateWithIndex(g, idx)
		logger.DumpLog.Println("idx : ", idx)
		logger.DumpLog.Println("current exp step : ", g.State)
	}
	return currentExpStep
}

func UpdateGameStateWithIndex(g *data.GameState, idx int) *data.ExpeditionStepResolveInfo {
	var currentExpStep *data.ExpeditionStepResolveInfo = nil

	if idx < 0 || idx >= len(g.CurrentExpedition.WhatHappened) {
		g.State = data.Home
		g.CurrentExpedition = nil
		return nil
	}

	currentExpStep = g.CurrentExpedition.WhatHappened[idx]

	g.State = currentExpStep.StepState
	// if g.State == data.Fight {
	// 	glTeam := gamedata.GetEnemyTeamForEvent(g.CurrentExpedition.Identifier, idx)
	// 	g.ETeam = (*data.Team)(glTeam)
	// }
	return currentExpStep
}
