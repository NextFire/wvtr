package data

type EquipableType int

const (
	// Weapon
	Sword EquipableType = iota
	Hammer
	Spear
	Bow
	Crossbow
	Wand
	Staff

	// Armor
	Robe
	Plate
	Leather

	// Omamori
)

type IEquipableBase interface {
	GenEquipment() IEquipable
}

// Don't need that in db, it's used by the back to create new item
type EquipableBase struct {
	BaseName    string
	IconURL     string
	AffixesPool []*Affix
}

type WeaponBase struct {
	EquipableBase
	BaseDamagesRange map[DamageType]*StatsRange
	BaseCritRate     *StatsRange
	BaseAttackSpeed  *StatsRange

	// Damage scaling (multiplicator between [0.01 and 2])
	StrScaling float64 `json:"strScaling"`
	IntScaling float64 `json:"intScaling"`
	DexScaling float64 `json:"dexScaling"`
	LckScaling float64 `json:"lckScaling"`
}

type ArmorBase struct {
	EquipableBase
	BaseBlockScore       *StatsRange
	BaseEvadeScore       *StatsRange
	BaseResistancesRange map[DamageType]*StatsRange
}

type OmamoriBase struct {
	EquipableBase
}

func (e *Equipable) GetTotalValueOfAffixInEquipment(af AffixType) *Affix {
	res := &Affix{
		Type: af,
	}

	for _, a := range e.Affixes {
		if a.Type == af {
			res.Ranges = a.Ranges
		}
	}
	return res
}

func (base *WeaponBase) GenEquipment() IEquipable {
	baseCritValue := base.BaseCritRate.RollValue()
	baseAttSpeed := base.BaseAttackSpeed.RollValue()
	resEqui := &Weapon{
		BaseDamage:      CreateDamageMapRange(base.BaseDamagesRange),
		BaseCritRate:    &baseCritValue,
		BaseAttackSpeed: &baseAttSpeed,
		Equipable: Equipable{
			Storable: Storable{
				IconURL: base.IconURL,
				Name:    base.BaseName,
			},
			Affixes: make([]*Affix, 0),
		},
		StrScaling: base.StrScaling,
		IntScaling: base.IntScaling,
		DexScaling: base.DexScaling,
		LckScaling: base.LckScaling,
	}

	// TODO : Item quality + Roll affix ?

	return resEqui
}

func (base *ArmorBase) GenEquipment() IEquipable {
	resistances := CreateDamageMapRange(base.BaseResistancesRange)
	baseEvade := base.BaseEvadeScore.RollValue()
	baseBlock := base.BaseBlockScore.RollValue()
	resEqui := &Armor{
		BaseResistancesRange: resistances,
		BlockScore:           &baseBlock,
		EvadeScore:           &baseEvade,
		Equipable: Equipable{
			Storable: Storable{
				IconURL: base.IconURL,
				Name:    base.BaseName,
			},
			Affixes: make([]*Affix, 0),
		},
	}

	// TODO : Item quality + Roll affix ?

	return resEqui
}

func (base *OmamoriBase) GenEquipment() IEquipable {
	resEqui := &Omamori{
		Equipable: Equipable{
			Storable: Storable{
				IconURL: base.IconURL,
				Name:    base.BaseName,
			},
			Affixes: make([]*Affix, 0),
		},
	}

	// TODO : Item quality + Roll affix ?

	return resEqui
}
