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
