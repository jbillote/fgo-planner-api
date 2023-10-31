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

var arcSkills = []model.Skill{
	{
		Name: "Rainbow Mystic Eyes A",
		Icon: "https://static.atlasacademy.io/JP/SkillIcons/skill_00512.png",
	},
}

func TestGetSortedKeysSortedInput(t *testing.T) {
	got := getSortedKeys(mapSortedKeys)
	want := []string{"f", "g", "o"}
	if !arrayEquality(got, want) {
		t.Errorf("Got %s, expected %s", got, want)
	}
}

func TestGetSortedKeysUnsortedInput(t *testing.T) {
	got := getSortedKeys(mapUnsortedKeys)
	want := []string{"f", "g", "o"}
	if !arrayEquality(got, want) {
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

func arrayEquality(a []string, b []string) bool {
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
