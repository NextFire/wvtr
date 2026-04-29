package data

import (
	"fmt"
	"slices"
	"time"
)

func NewExpeditionResolveInfo(state EncounterState) *ExpeditionStepResolveInfo {
	return &ExpeditionStepResolveInfo{
		StepState: state,
		Timeline:  make([]*ExpeditionStepTimestamp, 0),
	}
}

func (e *ExpeditionStepResolveInfo) String() string {
	res := ""
	arr := e.Timeline
	for _, s := range arr {
		res += fmt.Sprintf("%s: %s\n", s.When, s.What)
		if s.WhatAction != nil {
			res += s.WhatAction.String()
			res += "\n"
		}
	}
	return res
}

func (e *ExpeditionStepResolveInfo) AddNewHappening(when time.Time, what string, fad *FieldActionDesc) {
	e.Timeline = append(e.Timeline, &ExpeditionStepTimestamp{
		When:       when,
		WhatAction: fad,
		What:       what})

	// order by time
	slices.SortFunc(e.Timeline, func(a, b *ExpeditionStepTimestamp) int {
		return a.When.Compare(b.When)
	})
}

func (e *ExpeditionStepResolveInfo) GetDuration() time.Duration {
	return e.Timeline[len(e.Timeline)-1].When.Sub(e.Timeline[0].When)
}
