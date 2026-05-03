package expedition

import (
	"time"
	"wvtrserv/data"
)

type ExpToSendToFront struct {
	Key      string        `json:"key"`
	ImgURL   string        `json:"imgURL"`
	Duration time.Duration `json:"duration"`
}

type Expedition struct {
	StartedAt time.Time
	ImgURL    string
	HTeam     *data.Team
	Events    []ExpeditionEvent
}

func (e *Expedition) Solve(identifier string, pTeam *data.Team) *data.ExpeditionDB {
	var t time.Time = time.Now()
	happened := make([]*data.ExpeditionStepResolveInfo, 0)
	for _, ev := range e.Events {
		event := ev.CopyEvent()
		happened = append(happened, event.Solve(t, pTeam))
		t = t.Add(event.GetDuration())
		if pTeam.IsDefeated() {
			break
		}
	}
	edb := &data.ExpeditionDB{
		Identifier:   identifier,
		StartedAt:    t.UTC(),
		WhatHappened: happened,
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
