package expedition

import (
	"time"
	"wvtrserv/data"
)

type FightEvent struct {
	EEvent
	EnemyPool []*data.Hero
}

func NewFightEvent(areaEnemyPool []*data.Hero, name string) *FightEvent {
	res := &FightEvent{
		EEvent: EEvent{
			duration: 0,
			name:     name,
		},
		EnemyPool: areaEnemyPool,
	}
	res.EventSolvingInfo = data.NewExpeditionResolveInfo(res.GetEventType())

	return res
}

func (e FightEvent) GetEventType() data.EncounterState {
	return data.Fight
}

func (e FightEvent) GenerateTeamToFightFromAreaPool() *data.Team {
	var teamHeroes []*data.Hero = make([]*data.Hero, 0)
	numberOfEnemies := int(data.NaturalRoll(1, 4))
	for range numberOfEnemies {
		ene := e.EnemyPool[int(data.NaturalRoll(0, float64(len(e.EnemyPool))))]
		ene.Equipment.Weapon.RollBaseStats()
		ene.ClearAllStatusAndSetToFullLife()
		teamHeroes = append(teamHeroes, ene)
	}
	return &data.Team{
		Heroes: teamHeroes,
	}
}

func (e *FightEvent) Solve(startAt time.Time, heroTeam *data.Team) *data.ExpeditionStepResolveInfo {
	resExp := data.NewExpeditionResolveInfo(e.GetEventType())

	resExp.AddNewHappening(startAt, "Fight start", nil)
	resExp.ETeam = e.GenerateTeamToFightFromAreaPool()

	heroTeam.Fight(resExp.ETeam, resExp)
	e.duration = resExp.GetDuration()
	resExp.AddNewHappening(startAt.Add(e.GetDuration()), "Fight End", nil)
	return resExp
}

func (e FightEvent) CopyEvent() ExpeditionEvent {
	return &FightEvent{
		EEvent: EEvent{
			duration:         e.duration,
			EventSolvingInfo: &data.ExpeditionStepResolveInfo{},
			name:             e.name,
		},
		EnemyPool: e.EnemyPool,
	}
}

// func Fight(heroTeam *data.Team, enemyTeam *data.Team, infos *data.ExpeditionStepResolveInfo) {
// 	startTime := infos.Timeline[0].When

// 	turnOrder := NewFightTurnOrderTimeline(*heroTeam, *enemyTeam, startTime)
// 	escapeTime := 30 * time.Minute
// 	time := startTime
// 	for !heroTeam.IsDefeated() && !enemyTeam.IsDefeated() {
// 		if turnOrder.fightTotalDuration(startTime) < escapeTime {
// 			// too long, the team escape
// 			return
// 		}
// 		a := turnOrder.getNextAction(time)
// 		recupTime := a.Who.Play(heroTeam, enemyTeam)
// 		turnOrder.addAction(&Action{
// 			When: a.When.Add(recupTime),
// 			Who:  a.Who,
// 		})
// 		time = a.When
// 	}

// 	// while !fightFinished
// 	// Initiativelist[current]
// }
