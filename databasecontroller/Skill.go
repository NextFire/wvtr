package databasecontroller

import "wvtrserv/data"

func CreateSkill(skill *data.Skill) {
	db.Create(skill)
}

func GetSkills() []*data.Skill {
	res := []*data.Skill{}
	db.Find(&res)
	return res
}

func GetSkillByID(id uint) *data.Skill {
	var skill *data.Skill = &data.Skill{}

	db.Find(&skill, id)

	return skill
}
