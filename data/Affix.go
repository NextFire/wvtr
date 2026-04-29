package data

type AffixType int

const (
	FlatLife AffixType = iota
	FlatStr
	FlatInt
	FlatDex
	FlatLck
	PercentLife
	PercentStr
	PercentInt
	PercentDex
	PercentLck
	PercentRecupSpeed
	CriticalChance
	CriticalDamage
	FlatPhysDmg
	FlatElemDmg
	FlatSlashDmg
	FlatBluntDmg
	FlatPierceDmg
	FlatFireDmg
	FlatFrostDmg
	FlatLightningDmg

	// res
	PhysRes
	ElemRes
	SlashRes
	BluntRes
	PierceRes
	FireRes
	FrostRes
	LightningRes
)

func (a *Affix) RollValue() {
	for _, r := range a.Ranges {
		r.RollValue()
	}
}

func (a *Affix) Add(b *Affix) *Affix {
	if len(b.Ranges) == 0 {
		return a
	}

	if len(a.Ranges) == 0 && a.Type == b.Type {
		a.Ranges = b.Ranges
		return a
	}

	if a.Type == b.Type && a.Ranges != nil && len(a.Ranges) == len(b.Ranges) {
		for i := range a.Ranges {
			a.Ranges[i].Value += b.Ranges[i].Value
		}
		return a
	}

	return nil
}
