package databasecontroller

import "wvtrserv/data"

func UpdateTeam(user *data.User) {
	db.Model(&user.CurrentTeam).
		Association("Heroes").
		Replace(user.CurrentTeam.Heroes)
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
