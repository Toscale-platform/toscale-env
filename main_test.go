package env

import (
	"slices"
	"testing"
)

func TestGetString(t *testing.T) {
	Init()

	gotStr := GetString("STR")
	wantStr := "text"

	if gotStr != wantStr {
		t.Errorf("got %q, wanted %q", gotStr, wantStr)
	}
}

func TestGetInt(t *testing.T) {
	Init()

	gotInt := GetInt("INT")
	wantInt := 100

	if gotInt != wantInt {
		t.Errorf("got %q, wanted %q", gotInt, wantInt)
	}
}

func TestGetSlice(t *testing.T) {
	Init()

	gotSlice := GetSlice("SLICE")
	wantSlice := []string{"el1", "el2", "el3"}

	if slices.Compare(gotSlice, wantSlice) != 0 {
		t.Errorf("got %q, wanted %q", gotSlice, wantSlice)
	}
}

func TestGetBool(t *testing.T) {
	Init()

	gotBool := GetBool("BOOL")

	if gotBool != true {
		t.Errorf("got %t, wanted %t", gotBool, true)
	}
}

func TestEmptyString(t *testing.T) {
	Init()

	gotEmptyString := GetString("NOT EXIST KEY")
	wantEmptyString := ""

	if gotEmptyString != wantEmptyString {
		t.Errorf("got %q, wanted %q", gotEmptyString, wantEmptyString)
	}
}

func TestEmptyInt(t *testing.T) {
	Init()

	gotInt := GetInt("NOT EXIST KEY")
	wantInt := 0

	if gotInt != wantInt {
		t.Errorf("got %q, wanted %q", gotInt, wantInt)
	}
}

func TestEmptySlice(t *testing.T) {
	Init()

	gotSlice := GetSlice("NOT EXIST KEY")
	var wantSlice []string

	if slices.Compare(gotSlice, wantSlice) != 0 {
		t.Errorf("got %q, wanted %q", gotSlice, wantSlice)
	}
}

func TestEmptyBool(t *testing.T) {
	Init()

	gotBool := GetBool("NOT EXIST KEY")

	if gotBool != false {
		t.Errorf("got %t, wanted %t", gotBool, false)
	}
}
