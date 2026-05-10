package gamedata

import (
	"time"
	"wvtrserv/gamelogic/expedition"
)

var traval30s expedition.ExpeditionEvent = expedition.NewTravelEvent(time.Second*30, "Travel", travelReward1)
var traval40s expedition.ExpeditionEvent = expedition.NewTravelEvent(time.Second*40, "Travel", travelReward1)
var traval10s expedition.ExpeditionEvent = expedition.NewTravelEvent(time.Second*10, "Travel", travelReward1)
var goingToPlains expedition.ExpeditionEvent = expedition.NewTravelEvent(time.Second*30, "Going to the plains", travelReward1)
