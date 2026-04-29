package gamedata

import (
	"wvtrserv/data"
)

func GetRandomHeroClass(classes []*data.HeroClass) *data.HeroClass {
	weights := make([]float64, 0)
	for _, c := range classes {
		weights = append(weights, c.Weight)
	}
	weights = data.NormalizeArray(weights)
	idx := data.RollInArrayWithRate(data.NaturalRoll(0, 1), weights)
	return classes[idx]
}
