package gamedata

import (
	"wvtrserv/data"
	"wvtrserv/gamelogic/expedition"
)

var Expeditions = map[string]expedition.Expedition{
	"Cost(10g) heal": {
		ImgURL:     DOMAIN_NAME + "/imgs/expeditions/base_expedition.png",
		Cost:       allCurrencies[data.Gold],
		CostNumber: 10,
		Events: []expedition.ExpeditionEvent{
			testsmallRest1,
		},
	},
	"Traveling 10 sec": {
		ImgURL: DOMAIN_NAME + "/imgs/expeditions/base_expedition.png",
		Events: []expedition.ExpeditionEvent{
			traval10s,
		},
	},
	"Rest": {
		ImgURL: DOMAIN_NAME + "/imgs/expeditions/base_expedition.png",
		Events: []expedition.ExpeditionEvent{
			testsmallRest1,
		},
	},
	"Plain quest": {
		ImgURL: DOMAIN_NAME + "/imgs/expeditions/base_expedition.png",
		Events: []expedition.ExpeditionEvent{
			goingToPlains,
			plainFight,
			goingToPlains,
		},
	},
	"Training": {
		ImgURL: DOMAIN_NAME + "/imgs/expeditions/self_training.png",
		Events: []expedition.ExpeditionEvent{
			selfTraining,
		},
	},
	"Work 10 sec": {
		ImgURL: DOMAIN_NAME + "/imgs/expeditions/base_expedition.png",
		Events: []expedition.ExpeditionEvent{
			workShort,
		},
	},
}

func GetAvailableExpeditions(user *data.User) []*expedition.ExpToSendToFront {
	res := make([]*expedition.ExpToSendToFront, 0)

	for k, v := range Expeditions {
		name := ""
		cbl := true
		if v.Cost != nil {
			name = v.Cost.GetName()
			cbl = user.Inventory.IsInInventory(v.Cost, v.CostNumber)
		}
		res = append(res, &expedition.ExpToSendToFront{
			Key:           k,
			ImgURL:        v.ImgURL,
			Duration:      v.GetMinimumTotalTime(),
			CostName:      name,
			CostNumber:    v.CostNumber,
			CanBeLaunched: cbl,
		})
	}
	return res
}

// func GetEnemyTeamForEvent(identifier string, idx int) *data.Team {
// 	fEvent, ok := Expeditions[identifier].Events[idx].(expedition.FightEvent)
// 	if !ok {
// 		logger.ErrLog.Printf("can't cast %dth event from %s expedition event into a fight event.", idx, identifier)
// 		return nil
// 	}
// 	return fEvent.ETeam
// }
