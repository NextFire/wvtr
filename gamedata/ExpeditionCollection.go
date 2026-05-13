package gamedata

import (
	"slices"
	"wvtrserv/data"
	"wvtrserv/gamelogic/expedition"
)

var ExpeditionsJobs = map[string]expedition.Expedition{
	"Work 10 sec": {
		ImgURL: DOMAIN_NAME + "/imgs/expeditions/base_expedition.png",
		Events: []expedition.ExpeditionEvent{
			workShort,
		},
		Order: 0,
	},
}

var ExpeditionsHeal = map[string]expedition.Expedition{
	"Cost(10g) heal": {
		ImgURL:     DOMAIN_NAME + "/imgs/expeditions/base_expedition.png",
		Cost:       allCurrencies[data.Gold],
		CostNumber: 10,
		Events: []expedition.ExpeditionEvent{
			testsmallRest1,
		},
		Order: 1,
	},
	"Rest": {
		ImgURL: DOMAIN_NAME + "/imgs/expeditions/base_expedition.png",
		Events: []expedition.ExpeditionEvent{
			testsmallRest1,
		},
		Order: 0,
	},
}

var ExpeditionsQuests = map[string]expedition.Expedition{
	"Plain quest": {
		ImgURL: DOMAIN_NAME + "/imgs/expeditions/base_expedition.png",
		Events: []expedition.ExpeditionEvent{
			goingToPlains,
			plainFight,
			goingToPlains,
		},
		Order: 0,
	},
}

var ExpeditionsTest = map[string]expedition.Expedition{
	"Traveling 10 sec": {
		ImgURL: DOMAIN_NAME + "/imgs/expeditions/base_expedition.png",
		Events: []expedition.ExpeditionEvent{
			traval10s,
		},
		Order: 0,
	},
	"Training": {
		ImgURL: DOMAIN_NAME + "/imgs/expeditions/self_training.png",
		Events: []expedition.ExpeditionEvent{
			selfTraining,
		},
		Order: 1,
	},
	"Craft Weapon": {
		ImgURL: DOMAIN_NAME + "/imgs/expeditions/base_expedition.png",
		Events: []expedition.ExpeditionEvent{
			craftWeapon1,
		},
		Order: 2,
	},
}

var Expeditions = map[string]map[string]expedition.Expedition{
	"Jobs":   ExpeditionsJobs,
	"Heal":   ExpeditionsHeal,
	"Quests": ExpeditionsQuests,
	"Tests":  ExpeditionsTest,
}

func GetAvailableExpeditions(user *data.User) []*expedition.ExpToSendToFront {
	res := make([]*expedition.ExpToSendToFront, 0)

	for category, exps := range Expeditions {
		for k, v := range exps {
			name := ""
			cbl := true
			if v.Cost != nil {
				name = v.Cost.GetName()
				cbl = user.Inventory.IsInInventory(v.Cost, v.CostNumber)
			}
			res = append(res, &expedition.ExpToSendToFront{
				Category:      category,
				Key:           k,
				ImgURL:        v.ImgURL,
				Duration:      v.GetMinimumTotalTime(),
				CostName:      name,
				CostNumber:    v.CostNumber,
				CanBeLaunched: cbl,
				Order:         v.Order,
			})
		}
	}
	// order by Order value
	slices.SortFunc(res, func(a, b *expedition.ExpToSendToFront) int {
		if a.Order < b.Order {
			return -1
		}
		return 1
	})
	return res
}
