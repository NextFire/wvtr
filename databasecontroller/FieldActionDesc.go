package databasecontroller

import "wvtrserv/data"

func GetFieldActionDescByID(id uint) *data.FieldActionDesc {
	var res *data.FieldActionDesc = &data.FieldActionDesc{}
	db.Preload("FromH").
		Preload("UsedSkill").
		Preload("TargetH").
		Find(&res, id)

	if res.FromH != nil {
		res.FromH = GetHeroByID(res.FromH.ID)
	}
	if res.TargetH != nil {
		res.TargetH = GetHeroByID(res.TargetH.ID)
	}
	if res.UsedSkill != nil {
		res.UsedSkill = GetSkillByID(res.UsedSkill.ID)
	}

	return res
}
