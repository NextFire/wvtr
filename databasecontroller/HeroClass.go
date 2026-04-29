package databasecontroller

import "wvtrserv/data"

func CreateHeroClass(hc *data.HeroClass) {
	db.Create(hc)
}

func GetHeroClasses() []*data.HeroClass {
	res := []*data.HeroClass{}
	db.Find(&res)
	return res
}

func GetHeroClassByID(id uint) *data.HeroClass {
	var class *data.HeroClass = &data.HeroClass{}
	
	db.Find(&class, id)

	return class
}
