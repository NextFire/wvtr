package gamedata

import "wvtrserv/data"

var FlatLife map[int]*data.Affix = map[int]*data.Affix{
	1: {
		Name: "+{amount} to max life",
		Ranges: []*data.StatsRange{
			{Min: 5, Max: 10},
		},
		Type: data.FlatLife,
	},
}

var FlatStr map[int]*data.Affix = map[int]*data.Affix{
	1: {
		Name: "+{amount} to strength",
		Ranges: []*data.StatsRange{
			{Min: 1, Max: 5},
		},
		Type: data.FlatStr,
	},
}

var FlatInt map[int]*data.Affix = map[int]*data.Affix{
	1: {
		Name: "+{amount} to intelligence",
		Ranges: []*data.StatsRange{
			{Min: 1, Max: 5},
		},
		Type: data.FlatInt,
	},
}

var FlatDex map[int]*data.Affix = map[int]*data.Affix{
	1: {
		Name: "+{amount} to dexterity",
		Ranges: []*data.StatsRange{
			{Min: 1, Max: 5},
		},
		Type: data.FlatDex,
	},
}

var FlatLck map[int]*data.Affix = map[int]*data.Affix{
	1: {
		Name: "+{amount} to luck",
		Ranges: []*data.StatsRange{
			{Min: 1, Max: 5},
		},
		Type: data.FlatLck,
	},
}

var PercentLife map[int]*data.Affix = map[int]*data.Affix{
	1: {
		Name: "+{amount}`%` to max life",
		Ranges: []*data.StatsRange{
			{Min: 5, Max: 10},
		},
		Type: data.PercentLife,
	},
}

var PercentStr map[int]*data.Affix = map[int]*data.Affix{
	1: {
		Name: "+{amount}`%` to strength",
		Ranges: []*data.StatsRange{
			{Min: 5, Max: 10},
		},
		Type: data.PercentStr,
	},
}

var PercentInt map[int]*data.Affix = map[int]*data.Affix{
	1: {
		Name: "+{amount}`%` to intelligence",
		Ranges: []*data.StatsRange{
			{Min: 5, Max: 10},
		},
		Type: data.PercentInt,
	},
}

var PercentDex map[int]*data.Affix = map[int]*data.Affix{
	1: {
		Name: "+{amount}`%` to dexterity",
		Ranges: []*data.StatsRange{
			{Min: 5, Max: 10},
		},
		Type: data.PercentDex,
	},
}

var PercentLck map[int]*data.Affix = map[int]*data.Affix{
	1: {
		Name: "+{amount}`%` to luck",
		Ranges: []*data.StatsRange{
			{Min: 5, Max: 10},
		},
		Type: data.PercentLck,
	},
}

var PercentRecupSpeed map[int]*data.Affix = map[int]*data.Affix{
	1: {
		Name: "Reduce recuperation time by {amount}`%`",
		Ranges: []*data.StatsRange{
			{Min: 1, Max: 5},
		},
		Type: data.PercentRecupSpeed,
	},
}

var CriticalChance map[int]*data.Affix = map[int]*data.Affix{
	1: {
		Name: "Increased critical chance by {amount}`%`",
		Ranges: []*data.StatsRange{
			{Min: 5, Max: 10},
		},
		Type: data.CriticalChance,
	},
}

var CriticalDamage map[int]*data.Affix = map[int]*data.Affix{
	1: {
		Name: "Increased critical damage by {amount}%",
		Ranges: []*data.StatsRange{
			{Min: 5, Max: 10},
		},
		Type: data.CriticalDamage,
	},
}

var FlatPhysicalDamage map[int]*data.Affix = map[int]*data.Affix{
	1: {
		Name: "+{amount} to physical damage",
		Ranges: []*data.StatsRange{
			{Min: 1, Max: 5},
		},
		Type: data.FlatPhysDmg,
	},
}

var FlatElementalDamage map[int]*data.Affix = map[int]*data.Affix{
	1: {
		Name: "+{amount} to elemental damage",
		Ranges: []*data.StatsRange{
			{Min: 1, Max: 5},
		},
		Type: data.FlatElemDmg,
	},
}

var FlatSlashDamage map[int]*data.Affix = map[int]*data.Affix{
	1: {
		Name: "+{amount} to slash damage",
		Ranges: []*data.StatsRange{
			{Min: 4, Max: 8},
		},
		Type: data.FlatSlashDmg,
	},
}

var FlatBluntDamage map[int]*data.Affix = map[int]*data.Affix{
	1: {
		Name: "+{amount} to slash damage",
		Ranges: []*data.StatsRange{
			{Min: 4, Max: 8},
		},
		Type: data.FlatBluntDmg,
	},
}

var FlatPierceDamage map[int]*data.Affix = map[int]*data.Affix{
	1: {
		Name: "+{amount} to slash damage",
		Ranges: []*data.StatsRange{
			{Min: 4, Max: 8},
		},
		Type: data.FlatPierceDmg,
	},
}

var FlatFireDamage map[int]*data.Affix = map[int]*data.Affix{
	1: {
		Name: "+{amount} to fire damage",
		Ranges: []*data.StatsRange{
			{Min: 4, Max: 8},
		},
		Type: data.FlatFireDmg,
	},
}

var FlatFrostDamage map[int]*data.Affix = map[int]*data.Affix{
	1: {
		Name: "+{amount} to frost damage",
		Ranges: []*data.StatsRange{
			{Min: 4, Max: 8},
		},
		Type: data.FlatFrostDmg,
	},
}

var FlatLightningDamage map[int]*data.Affix = map[int]*data.Affix{
	1: {
		Name: "+{amount} to lightning damage",
		Ranges: []*data.StatsRange{
			{Min: 4, Max: 8},
		},
		Type: data.FlatLightningDmg,
	},
}

var PhysicalRes map[int]*data.Affix = map[int]*data.Affix{
	1: {
		Name: "+{amount} to physical resistance",
		Ranges: []*data.StatsRange{
			{Min: 1, Max: 5},
		},
		Type: data.PhysRes,
	},
}

var ElementalRes map[int]*data.Affix = map[int]*data.Affix{
	1: {
		Name: "+{amount} to elemental resistance",
		Ranges: []*data.StatsRange{
			{Min: 1, Max: 5},
		},
		Type: data.ElemRes,
	},
}

var SlashRes map[int]*data.Affix = map[int]*data.Affix{
	1: {
		Name: "+{amount} to slash resistance",
		Ranges: []*data.StatsRange{
			{Min: 4, Max: 8},
		},
		Type: data.SlashRes,
	},
}

var BluntRes map[int]*data.Affix = map[int]*data.Affix{
	1: {
		Name: "+{amount} to slash resistance",
		Ranges: []*data.StatsRange{
			{Min: 4, Max: 8},
		},
		Type: data.BluntRes,
	},
}

var PierceRes map[int]*data.Affix = map[int]*data.Affix{
	1: {
		Name: "+{amount} to slash resistance",
		Ranges: []*data.StatsRange{
			{Min: 4, Max: 8},
		},
		Type: data.PierceRes,
	},
}

var FireRes map[int]*data.Affix = map[int]*data.Affix{
	1: {
		Name: "+{amount} to fire resistance",
		Ranges: []*data.StatsRange{
			{Min: 4, Max: 8},
		},
		Type: data.FireRes,
	},
}

var FrostRes map[int]*data.Affix = map[int]*data.Affix{
	1: {
		Name: "+{amount} to frost resistance",
		Ranges: []*data.StatsRange{
			{Min: 4, Max: 8},
		},
		Type: data.FrostRes,
	},
}

var LightningRes map[int]*data.Affix = map[int]*data.Affix{
	1: {
		Name: "+{amount} to lightning resistance",
		Ranges: []*data.StatsRange{
			{Min: 4, Max: 8},
		},
		Type: data.LightningRes,
	},
}
