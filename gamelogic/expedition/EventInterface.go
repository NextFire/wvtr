package expedition

import (
	"time"
	"wvtrserv/data"
)

type ExpeditionEvent interface {
	EndAt(startAt time.Time) time.Time
	GetEventType() data.EncounterState
	Solve(startAt time.Time, t *data.Team) *data.ExpeditionStepResolveInfo
	GetDuration() time.Duration
	GetName() string
	CopyEvent() ExpeditionEvent
}

type EEvent struct {
	duration         time.Duration
	EventSolvingInfo *data.ExpeditionStepResolveInfo
	name             string
}

func (e EEvent) EndAt(startAt time.Time) time.Time {
	return startAt.Add(e.duration)
}

func (e EEvent) GetDuration() time.Duration {
	return e.duration
}

func (e EEvent) GetName() string {
	return e.name
}
