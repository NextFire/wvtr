package gamedata

import (
	"time"
	"wvtrserv/data"
	"wvtrserv/gamelogic/expedition"
)

var traval30s expedition.ExpeditionEvent = expedition.NewTravelEvent(time.Second*30, "Travel")
var traval40s expedition.ExpeditionEvent = expedition.NewTravelEvent(time.Second*40, "Travel")
var traval10s expedition.ExpeditionEvent = expedition.NewTravelEvent(time.Second*10, "Travel")
var goingToPlains expedition.ExpeditionEvent = expedition.NewTravelEvent(time.Second*30, "Going to the plains")
var plainFight expedition.ExpeditionEvent = expedition.NewFightEvent(PlainPool, "Plain fight")

var nothing10s expedition.ExpeditionEvent = expedition.NewNeutralEvent(time.Second*10, "Neutral",
	expedition.HappeningType(func(t *data.Team, e *data.ExpeditionStepResolveInfo) {}))

var selfTraining expedition.ExpeditionEvent = expedition.NewNeutralEvent(time.Second*30, "Self training (30 minutes)",
	expedition.HappeningType(func(t *data.Team, e *data.ExpeditionStepResolveInfo) {
		for _, h := range t.Heroes {
			h.GainXP(5)
			e.AddNewHappening(time.Now(), h.Name+" Gained 5 xp points.", nil)
		}
	}))

var Expeditions = map[string]expedition.Expedition{
	"Travel 30 sec": {
		ImgURL: DOMAIN_NAME + "/imgs/expeditions/base_expedition.png",
		Events: []expedition.ExpeditionEvent{
			traval30s,
		},
	},
	"Travel 40 sec": {
		ImgURL: DOMAIN_NAME + "/imgs/expeditions/base_expedition.png",
		Events: []expedition.ExpeditionEvent{
			traval40s,
		},
	},
	"Travel and do nothing": {
		ImgURL: DOMAIN_NAME + "/imgs/expeditions/base_expedition.png",
		Events: []expedition.ExpeditionEvent{
			traval10s,
			nothing10s,
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
}

func GetAvailableExpeditions() []*expedition.ExpToSendToFront {
	res := make([]*expedition.ExpToSendToFront, 0)
	for k, v := range Expeditions {
		res = append(res, &expedition.ExpToSendToFront{
			Key:      k,
			ImgURL:   v.ImgURL,
			Duration: v.GetMinimumTotalTime(),
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
