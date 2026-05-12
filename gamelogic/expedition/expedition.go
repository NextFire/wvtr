package expedition

import (
	"time"
	"wvtrserv/data"
)

type ExpToSendToFront struct {
	Key           string        `json:"key"`
	ImgURL        string        `json:"imgURL"`
	Duration      time.Duration `json:"duration"`
	CostName      string        `json:"costName"`
	CostNumber    int           `json:"costNumber"`
	CanBeLaunched bool          `json:"canBeLaunched"`
}

type Expedition struct {
	StartedAt  time.Time
	ImgURL     string
	HTeam      *data.Team
	Cost       data.IStorable
	CostNumber int
	Events     []ExpeditionEvent
}

func (e *Expedition) CanEnter(user *data.User) bool {
	if e.CostNumber == 0 || user.Inventory.IsInInventory(e.Cost, e.CostNumber) {
		return true
	}
	return false
}

func (e *Expedition) Solve(identifier string, pTeam *data.Team, co []*data.CurrencyOwned) *data.ExpeditionDB {
	var t time.Time = time.Now()
	happened := make([]*data.ExpeditionStepResolveInfo, 0)
	for _, ev := range e.Events {
		happened = append(happened, ev.Solve(t, pTeam))
		t = t.Add(ev.GetDuration())
		if pTeam.IsDefeated() {
			break
		}
	}
	loot := data.NewInventory(co)
	re := e.GetReward()
	loot.StoreReward(re.Loot, re.Currencies)
	edb := &data.ExpeditionDB{
		Identifier:   identifier,
		StartedAt:    t.UTC(),
		WhatHappened: happened,
		ExpeditionRewards: &data.Reward{
			XP:   re.XP,
			Loot: loot,
		},
	}
	return edb
}

func (e Expedition) GetMinimumTotalTime() time.Duration {
	var res time.Duration = 0
	for _, ev := range e.Events {
		res += ev.GetDuration()
	}
	return res
}

func (e Expedition) GetReward() *Reward {
	res := NewReward(nil)
	for _, ev := range e.Events {
		res.MergeReward(ev.GetReward())
	}
	return res
}

func (e Expedition) GetCopy() Expedition {
	res := Expedition{
		StartedAt:  e.StartedAt,
		ImgURL:     e.ImgURL,
		HTeam:      e.HTeam,
		Cost:       e.Cost,
		CostNumber: e.CostNumber,
		Events:     make([]ExpeditionEvent, len(e.Events)),
	}
	for i, ev := range e.Events {
		res.Events[i] = ev.CopyEvent()
	}

	return res
}
