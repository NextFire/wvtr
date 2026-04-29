package data

func (he *HeroEquipment) GetTotalValueOfAffixInEquipment(af AffixType) *Affix {
	res := &Affix{Type: af}
	if he.Weapon != nil {
		res.Add(he.Weapon.GetTotalValueOfAffixInEquipment(af))
	}
	if he.Armor != nil {
		res.Add(he.Armor.GetTotalValueOfAffixInEquipment(af))
	}
	if he.Omamori != nil {
		res.Add(he.Omamori.GetTotalValueOfAffixInEquipment(af))
	}

	// No affixes have been found
	if len(res.Ranges) == 0 {
		res.Ranges = []*StatsRange{
			{Value: 0},
		}
	}

	return res
}
