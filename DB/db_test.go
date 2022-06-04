package DB

import (
	"testing"
	"fmt"
)

func TestDBSetValue(t *testing.T) {
	db := NewDB()

	value := "test"
	key := "x"

	db.AddRecord(key, value)
	if db.GetRecord(key) != value {
		t.Error(fmt.Sprintf("The parameter %s doesn't match with %s", "test", "test"))
	}
}

func TestDBNumeToEqual(t *testing.T) {
	db := NewDB()

	value := "test"
	keys := []string{"x", "y", "z"}

	for _, key := range keys {
		db.AddRecord(key, value)
	}
	
	if db.NumeToEqual(value) != len(keys) {
		t.Error(fmt.Sprintf("The expected result was: %v, but I got %v", len(keys), db.NumeToEqual(value)))
	}
}

func TestDBUnSetValue(t *testing.T) {
	db := NewDB()

	value := "test"
	key := "x"

	db.AddRecord(key, value)
	if db.GetRecord(key) != value {
		t.Error(fmt.Sprintf("The parameter %s doesn't match with %s", "test", "test"))
	}

	db.UnSet(key)
	if db.GetRecord(key) == value {
		t.Error(fmt.Sprintf("Unexpeted parameter %s", "test"))
	}
}