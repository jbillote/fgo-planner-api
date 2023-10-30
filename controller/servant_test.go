package controller

import "testing"

func TestClassIconFilenameRarityInRange(t *testing.T) {
	want := classIconFilename(2, 1)
	got := "2_1"
	if got != want {
		t.Errorf("Got %s, expected %s", got, want)
	}
}

func TestClassIconFilenameRarityGreater(t *testing.T) {
	want := classIconFilename(9999, 1)
	got := "3_1"
	if got != want {
		t.Errorf("Got %s, expected %s", got, want)
	}
}

func TestClassIconFilenameRarityNegative(t *testing.T) {
	want := classIconFilename(-1, 1)
	got := "1_1"
	if got != want {
		t.Errorf("Got %s, expected %s", got, want)
	}
}
