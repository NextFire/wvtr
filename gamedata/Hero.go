package gamedata

import (
	"wvtrserv/data"
	"wvtrserv/nanapi/client"
)

func CreateNewHeroFromDBWaifuInfos(wc *client.JoinWC, classes []*data.HeroClass, skills []*data.Skill) *data.Hero {
	// select class
	class := GetRandomHeroClass(classes)
	attributes := data.NewHeroAttribute(class, data.GenerateGrowthRateFromRank(wc.Rank))
	uniqueSkill := GetRandomUniqueSkill(skills)
	resHero := &data.Hero{
		ImageUrl:       wc.ImageLarge,
		Name:           wc.NameUserPreferred,
		Class:          class,
		Rank:           wc.Rank,
		Attributes:     attributes,
		WeaponAttack:   GetAttackSkill(),
		UniqueSkill:    uniqueSkill,
		WaifuID:        wc.ID,
		AnilistCharaID: uint(wc.IdAl),
		Equipment: &data.HeroEquipment{
			Weapon: data.CreateWeapon(FistBase),
		},
	}

	resHero.LevelUp()
	resHero.Attributes.XPToLvlUp = resHero.Attributes.LevelThreshold()
	resHero.Attributes.CurrentHP = resHero.Attributes.MaxHP
	return resHero
}
