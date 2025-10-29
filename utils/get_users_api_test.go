package utils

import (
	"fmt"
	"testing"

	"github.com/wendehals/bricks-cli/api"
)

func TestGetUsersAPIOverrideUsed(t *testing.T) {
	// Save/restore any existing override
	prev := GetUsersAPIOverride
	defer func() { GetUsersAPIOverride = prev }()

	// Ensure cachedUsersAPI is cleared so we exercise the override path.
	usersAPIMutex.Lock()
	cachedUsersAPI = nil
	usersAPIMutex.Unlock()

	GetUsersAPIOverride = func() (*api.UsersAPI, error) {
		return nil, fmt.Errorf("injected-err")
	}

	_, err := GetUsersAPI()
	if err == nil || err.Error() != "injected-err" {
		t.Fatalf("expected injected error, got: %v", err)
	}
}
