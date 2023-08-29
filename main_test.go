package env

import (
	"slices"
	"testing"
)

func TestEnv(t *testing.T) {
	Init()

	gotStr := GetString("STR")
	wantStr := "text"

	if gotStr != wantStr {
		t.Errorf("got %q, wanted %q", gotStr, wantStr)
	}

	gotInt := GetInt("INT")
	wantInt := 100

	if gotInt != wantInt {
		t.Errorf("got %q, wanted %q", gotInt, wantInt)
	}

	gotSlice := GetSlice("SLICE")
	wantSlice := []string{"el1", "el2", "el3"}

	if slices.Compare(gotSlice, wantSlice) != 0 {
		t.Errorf("got %q, wanted %q", gotSlice, wantSlice)
	}
}
