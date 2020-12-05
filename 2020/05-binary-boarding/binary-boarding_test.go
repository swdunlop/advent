package main

import "testing"

func TestFollowPath(t *testing.T) {
	path := `FBFBBFFRLR`
	row, col, err := followPath(path, 128, 8)
	if err != nil {
		t.Fatal(err)
	}
	if row != 44 {
		t.Errorf(`expected row 44, got %v following %q`, row, path)
	}
	if col != 5 {
		t.Errorf(`expected col 5, got %v following %q`, col, path)
	}
}

func TestLocateSeat(t *testing.T) {
	path := `FBFBBFFRLR`
	seat, err := locateSeat(path, 128, 8)
	if err != nil {
		t.Fatal(err)
	}
	if seat != 357 {
		t.Errorf(`expected seat 357, got %v following %q`, seat, path)
	}
}
