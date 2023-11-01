package controller

import (
	"testing"

	"github.com/jbillote/fgo-planner-api/model"
)

var mapUnsortedKeys = map[string]string{
	"o": "test",
	"g": "test",
	"f": "test",
}

var mapSortedKeys = map[string]string{
	"f": "test",
	"g": "test",
	"o": "test",
}

var arcAppends = []appendPassive{
	{
		Skill: model.Skill{
			Name: "Extra Attack Boost",
			Icon: "https://static.atlasacademy.io/JP/SkillIcons/skill_00301.png",
		},
	},
	{
		Skill: model.Skill{
			Name: "Load Magical Energy",
			Icon: "https://static.atlasacademy.io/JP/SkillIcons/skill_00601.png",
		},
	},
	{
		Skill: model.Skill{
			Name: "Anti-Foreigner (ATK Up)",
			Icon: "https://static.atlasacademy.io/JP/SkillIcons/skill_00300.png",
		},
	},
}

var arcAscMaterialsUnsorted = map[string]materials{
	"1": materials{
		Items: []item{
			{
				Details: itemDetails{
					ID:   7005,
					Name: "Caster Piece",
					Icon: "https://static.atlasacademy.io/JP/Items/7005.png",
				},
				Amount: 12,
			},
			{
				Details: itemDetails{
					ID:   7007,
					Name: "Berserker Piece",
					Icon: "https://static.atlasacademy.io/JP/Items/7007.png",
				},
				Amount: 12,
			},
		},
		QP: 300000,
	},
	"0": materials{
		Items: []item{
			{
				Details: itemDetails{
					ID:   7005,
					Name: "Caster Piece",
					Icon: "https://static.atlasacademy.io/JP/Items/7005.png",
				},
				Amount: 5,
			},
			{
				Details: itemDetails{
					ID:   7007,
					Name: "Berserker Piece",
					Icon: "https://static.atlasacademy.io/JP/Items/7007.png",
				},
				Amount: 5,
			},
		},
		QP: 100000,
	},
	"3": materials{
		Items: []item{
			{
				Details: itemDetails{
					ID:   7105,
					Name: "Caster Monument",
					Icon: "https://static.atlasacademy.io/JP/Items/7105.png",
				},
				Amount: 12,
			},
			{
				Details: itemDetails{
					ID:   7107,
					Name: "Berserker Monument",
					Icon: "https://static.atlasacademy.io/JP/Items/7107.png",
				},
				Amount: 12,
			},
		},
		QP: 3000000,
	},
	"2": materials{
		Items: []item{
			{
				Details: itemDetails{
					ID:   7105,
					Name: "Caster Monument",
					Icon: "https://static.atlasacademy.io/JP/Items/7105.png",
				},
				Amount: 5,
			},
			{
				Details: itemDetails{
					ID:   7107,
					Name: "Berserker Monument",
					Icon: "https://static.atlasacademy.io/JP/Items/7107.png",
				},
				Amount: 5,
			},
		},
		QP: 1000000,
	},
}

var arcAscMaterialsSorted = map[string]materials{
	"0": materials{
		Items: []item{
			{
				Details: itemDetails{
					ID:   7005,
					Name: "Caster Piece",
					Icon: "https://static.atlasacademy.io/JP/Items/7005.png",
				},
				Amount: 5,
			},
			{
				Details: itemDetails{
					ID:   7007,
					Name: "Berserker Piece",
					Icon: "https://static.atlasacademy.io/JP/Items/7007.png",
				},
				Amount: 5,
			},
		},
		QP: 100000,
	},
	"1": materials{
		Items: []item{
			{
				Details: itemDetails{
					ID:   7005,
					Name: "Caster Piece",
					Icon: "https://static.atlasacademy.io/JP/Items/7005.png",
				},
				Amount: 12,
			},
			{
				Details: itemDetails{
					ID:   7007,
					Name: "Berserker Piece",
					Icon: "https://static.atlasacademy.io/JP/Items/7007.png",
				},
				Amount: 12,
			},
		},
		QP: 300000,
	},
	"2": materials{
		Items: []item{
			{
				Details: itemDetails{
					ID:   7105,
					Name: "Caster Monument",
					Icon: "https://static.atlasacademy.io/JP/Items/7105.png",
				},
				Amount: 5,
			},
			{
				Details: itemDetails{
					ID:   7107,
					Name: "Berserker Monument",
					Icon: "https://static.atlasacademy.io/JP/Items/7107.png",
				},
				Amount: 5,
			},
		},
		QP: 1000000,
	},
	"3": materials{
		Items: []item{
			{
				Details: itemDetails{
					ID:   7105,
					Name: "Caster Monument",
					Icon: "https://static.atlasacademy.io/JP/Items/7105.png",
				},
				Amount: 12,
			},
			{
				Details: itemDetails{
					ID:   7107,
					Name: "Berserker Monument",
					Icon: "https://static.atlasacademy.io/JP/Items/7107.png",
				},
				Amount: 12,
			},
		},
		QP: 3000000,
	},
}

var processedArcAscMaterials = []model.MaterialList{
	{
		Materials: []model.Material{
			{
				ID:     7005,
				Name:   "Caster Piece",
				Icon:   "https://static.atlasacademy.io/JP/Items/7005.png",
				Amount: 5,
			},
			{
				ID:     7007,
				Name:   "Berserker Piece",
				Icon:   "https://static.atlasacademy.io/JP/Items/7007.png",
				Amount: 5,
			},
		},
		QP: 100000,
	},
	{
		Materials: []model.Material{
			{
				ID:     7005,
				Name:   "Caster Piece",
				Icon:   "https://static.atlasacademy.io/JP/Items/7005.png",
				Amount: 12,
			},
			{
				ID:     7007,
				Name:   "Berserker Piece",
				Icon:   "https://static.atlasacademy.io/JP/Items/7007.png",
				Amount: 12,
			},
		},
		QP: 300000,
	},
	{
		Materials: []model.Material{
			{
				ID:     7105,
				Name:   "Caster Monument",
				Icon:   "https://static.atlasacademy.io/JP/Items/7105.png",
				Amount: 5,
			},
			{
				ID:     7107,
				Name:   "Berserker Monument",
				Icon:   "https://static.atlasacademy.io/JP/Items/7107.png",
				Amount: 5,
			},
		},
		QP: 1000000,
	},
	{
		Materials: []model.Material{
			{
				ID:     7105,
				Name:   "Caster Monument",
				Icon:   "https://static.atlasacademy.io/JP/Items/7105.png",
				Amount: 12,
			},
			{
				ID:     7107,
				Name:   "Berserker Monument",
				Icon:   "https://static.atlasacademy.io/JP/Items/7107.png",
				Amount: 12,
			},
		},
		QP: 3000000,
	},
}

func TestGetSortedKeysSortedInput(t *testing.T) {
	got := getSortedKeys(mapSortedKeys)
	want := []string{"f", "g", "o"}
	if !stringArrayEquality(got, want) {
		t.Errorf("Got %s, expected %s", got, want)
	}
}

func TestGetSortedKeysUnsortedInput(t *testing.T) {
	got := getSortedKeys(mapUnsortedKeys)
	want := []string{"f", "g", "o"}
	if !stringArrayEquality(got, want) {
		t.Errorf("Got %s, expected %s", got, want)
	}
}

func TestProcessMaterialListSortedInput(t *testing.T) {
	got := processMaterialList(arcAscMaterialsSorted)
	want := processedArcAscMaterials
	if !materialListArrayEquality(got, want) {
		t.Error("Actual and expected values differ")
	}
}

func TestProcessMaterialListUnsortedInput(t *testing.T) {
	got := processMaterialList(arcAscMaterialsUnsorted)
	want := processedArcAscMaterials
	if !materialListArrayEquality(got, want) {
		t.Error("Actual and expected values differ")
	}
}

func TestProcessAppends(t *testing.T) {
	got := processAppends(arcAppends)
	want := []model.Skill{
		{
			Name: "Extra Attack Boost",
			Icon: "https://static.atlasacademy.io/JP/SkillIcons/skill_00301.png",
		},
		{
			Name: "Load Magical Energy",
			Icon: "https://static.atlasacademy.io/JP/SkillIcons/skill_00601.png",
		},
		{
			Name: "Anti-Foreigner (ATK Up)",
			Icon: "https://static.atlasacademy.io/JP/SkillIcons/skill_00300.png",
		},
	}

	if !skillArrayEquality(got, want) {
		t.Errorf("Got %s, expected %s", got, want)
	}
}

func TestClassIconFilenameRarityInRange(t *testing.T) {
	got := classIconFilename(2, 1)
	want := "2_1"
	if got != want {
		t.Errorf("Got %s, expected %s", got, want)
	}
}

func TestClassIconFilenameRarityGreater(t *testing.T) {
	got := classIconFilename(9999, 1)
	want := "3_1"
	if got != want {
		t.Errorf("Got %s, expected %s", got, want)
	}
}

func TestClassIconFilenameRarityNegative(t *testing.T) {
	got := classIconFilename(-1, 1)
	want := "1_1"
	if got != want {
		t.Errorf("Got %s, expected %s", got, want)
	}
}

func stringArrayEquality(a []string, b []string) bool {
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}

	return true
}

func materialListArrayEquality(a []model.MaterialList, b []model.MaterialList) bool {
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if len(v.Materials) != len(b[i].Materials) {
			return false
		}

		for j, m := range v.Materials {
			if m != b[i].Materials[j] {
				return false
			}
		}

		if v.QP != b[i].QP {
			return false
		}
	}

	return true
}

func skillArrayEquality(a []model.Skill, b []model.Skill) bool {
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if v.Name != b[i].Name || v.Icon != b[i].Icon {
			return false
		}
	}

	return true
}
