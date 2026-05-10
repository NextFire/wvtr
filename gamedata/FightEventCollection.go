package gamedata

import "wvtrserv/gamelogic/expedition"

// fights
// Plains
var plainFight expedition.ExpeditionEvent = expedition.NewFightEvent(EnemyPlainPool, "Plain fight", plainsRewardPool)
