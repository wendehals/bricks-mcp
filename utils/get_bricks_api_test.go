package utils

import (
	"fmt"
	"testing"

	"github.com/wendehals/bricks-cli/api"
)

func TestGetBricksAPIOverrideUsed(t *testing.T) {
	// Save/restore any existing override
	prev := GetBricksAPIOverride
	defer func() { GetBricksAPIOverride = prev }()

	// Set override to return a controlled error
	GetBricksAPIOverride = func() (*api.BricksAPI, error) {
		return nil, fmt.Errorf("injected-err")
	}

	// Call GetBricksAPI and verify it returns the injected error
	_, err := GetBricksAPI()
	if err == nil || err.Error() != "injected-err" {
		t.Fatalf("expected injected error, got: %v", err)
	}
}
