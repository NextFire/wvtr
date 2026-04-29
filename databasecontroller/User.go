package databasecontroller

import (
	"wvtrserv/data"
	"wvtrserv/logger"
)

func CreateNewUser(user *data.User) *data.User {
	logger.DumpLog.Print("CreateNewUser")
	db.Create(user)
	return user
}

func GetUserByID(id uint) *data.User {
	var res *data.User = nil

	db.Preload("Inventory").
		Preload("State").
		Preload("CurrentTeam").
		Preload("OwnedHeroes").
		Find(&res, id)

	if res.Inventory != nil {
		res.Inventory = GetInventoryByID(res.Inventory.ID)
	}
	if res.State != nil {
		res.State = GetGameStateByID(res.State.ID)
	}
	if res.CurrentTeam != nil {
		res.CurrentTeam = GetTeamByID(res.CurrentTeam.ID)
	}
	for i := range res.OwnedHeroes {
		res.OwnedHeroes[i] = GetHeroByID(res.OwnedHeroes[i].ID)
	}

	logger.DumpLog.Println("Get user by id ", id)
	logger.DumpLog.Println(res)

	return res
}

func GetUserByDiscordID(did string) *data.User {
	var res *data.User = nil
	db.Where("discord_id = ?", did).Find(&res)
	if res != nil {
		res := GetUserByID(res.ID)
		logger.DumpLog.Println("GetUserByDiscordID: ", did, " | ", res.Name)
	} else {
		logger.DumpLog.Println("GetUserByDiscordID: ", did, " | user not found.")
	}
	return res
}

func UpdateUser(user *data.User) {
	db.Save(user)
}

func GetUserGameState(user *data.User) *data.User {
	user.State = GetGameState(user.State)
	return user
}

func GetUserCurrentTeam(user *data.User) *data.User {
	user.CurrentTeam = GetTeam(user.CurrentTeam)
	return user
}

func GetUserOwnedHeroes(user *data.User) *data.User {
	db.Preload("OwnedHeroes").
		Find(&user)

	for i := range user.OwnedHeroes {
		user.OwnedHeroes[i] = GetHeroByID(user.OwnedHeroes[i].ID)
	}
	return user
}
