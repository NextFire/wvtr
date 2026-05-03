package databasecontroller

import "wvtrserv/data"

func UpdateTeam(team *data.Team) {
	db.Model(&team).
		Association("Heroes").
		Replace(team.Heroes)
}

func SaveTeam(team *data.Team) {
	for _, h := range team.Heroes {
		SaveHero(h)
	}
}

func GetTeam(team *data.Team) *data.Team {
	db.Preload("Heroes").
		Find(&team)
	return team
}

func GetTeamByID(id uint) *data.Team {
	var res *data.Team
	db.Preload("Heroes").
		Find(&res, id)

	for i := range res.Heroes {
		res.Heroes[i] = GetHeroByID(res.Heroes[i].ID)
	}

	return res
}
