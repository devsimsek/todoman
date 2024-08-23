package utils_test

import (
	"testing"

	"go.smsk.dev/todoman/utils"
)

func TestContains(t *testing.T) {
	slice := []string{"a", "b", "c", "d", "e"}

	if !utils.Contains(slice, "a") {
		t.Errorf("Expected true, got false")
	}

	if utils.Contains(slice, "f") {
		t.Errorf("Expected false, got true")
	}
}
