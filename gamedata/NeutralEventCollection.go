package gamedata

import (
	"fmt"
	"time"
	"wvtrserv/data"
	"wvtrserv/gamelogic/expedition"
)

// Nothing events
var nothing10s expedition.ExpeditionEvent = expedition.NewNeutralEvent(time.Second*10, "Neutral", nil,
	expedition.HappeningType(func(selfEvent expedition.ExpeditionEvent, t *data.Team, e *data.ExpeditionStepResolveInfo) {}))

// train events
func trainingEventFactory(duration time.Duration, quantity float64) expedition.ExpeditionEvent {
	return expedition.NewNeutralEvent(duration, "Training", nil,
		expedition.HappeningType(func(selfEvent expedition.ExpeditionEvent, t *data.Team, e *data.ExpeditionStepResolveInfo) {
			selfEvent.GetReward().AddXP(quantity)
			e.AddNewHappening(time.Now().Add(duration-2*time.Millisecond), fmt.Sprintf("Team has trained and got %fxp", quantity), nil)
		}))
}

var selfTraining expedition.ExpeditionEvent = trainingEventFactory(10*time.Second, 5)

// rest events
func restEventFactory(duration time.Duration, quantity float64) expedition.ExpeditionEvent {
	return expedition.NewNeutralEvent(duration, "Resting", nil,
		expedition.HappeningType(func(selfEvent expedition.ExpeditionEvent, t *data.Team, e *data.ExpeditionStepResolveInfo) {
			for _, h := range t.Heroes {
				healed := h.Rest(quantity)
				e.AddNewHappening(time.Now().Add(duration-2*time.Millisecond), fmt.Sprintf("%s healed for %f hp points.", h.Name, healed), &data.FieldActionDesc{
					FromH:        h,
					FromPVChange: -healed,
				})
			}
		}))
}

var testsmallRest1 expedition.ExpeditionEvent = restEventFactory(30*time.Second, 10)
var testsmallRest2 expedition.ExpeditionEvent = restEventFactory(30*time.Second, 15)
var testsmallRest3 expedition.ExpeditionEvent = restEventFactory(30*time.Second, 20)

var smallRest1 expedition.ExpeditionEvent = restEventFactory(30*time.Minute, 10)
var smallRest2 expedition.ExpeditionEvent = restEventFactory(30*time.Minute, 15)
var smallRest3 expedition.ExpeditionEvent = restEventFactory(30*time.Minute, 20)

var mediumRest1 expedition.ExpeditionEvent = restEventFactory(2*time.Hour, 30)
var mediumRest2 expedition.ExpeditionEvent = restEventFactory(2*time.Hour, 50)
var mediumRest3 expedition.ExpeditionEvent = restEventFactory(2*time.Hour, 60)

var bigRest1 expedition.ExpeditionEvent = restEventFactory(23*time.Hour, 150)
var bigRest2 expedition.ExpeditionEvent = restEventFactory(23*time.Hour, 200)
var bigRest3 expedition.ExpeditionEvent = restEventFactory(23*time.Hour, 250)

// Work events
func workEventFactory(duration time.Duration, quantity float64) expedition.ExpeditionEvent {
	return expedition.NewNeutralEvent(duration, "Working",
		&expedition.RewardPool{
			CurrencyPool: map[data.CurrencyType]data.StatsRange{
				data.Gold: {Min: quantity, Max: quantity},
			},
		},
		expedition.HappeningType(func(selfEvent expedition.ExpeditionEvent, t *data.Team, e *data.ExpeditionStepResolveInfo) {
			e.AddNewHappening(time.Now().Add(duration-2*time.Millisecond), fmt.Sprintf("Gained %f gold", quantity), nil)
		}))
}

var workShort expedition.ExpeditionEvent = workEventFactory(10*time.Second, 5)
