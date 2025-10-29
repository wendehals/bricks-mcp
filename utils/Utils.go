package utils

import (
	"net/http"
	"sync"
	"time"

	"github.com/wendehals/bricks-cli/api"
	brickscliutils "github.com/wendehals/bricks-cli/utils"
)

var (
	// Cached credentials and its mutex
	cachedCredentials *api.Credentials
	credentialsMutex  sync.RWMutex

	// Cached Bricks API client and its mutex
	cachedBricksAPI *api.BricksAPI
	bricksAPIMutex  sync.RWMutex

	// Test hook: if non-nil, this function will be called instead of the
	// normal GetBricksAPI implementation. Tests may set this to inject a
	// fake BricksAPI or to simulate errors without touching disk/network.
	GetBricksAPIOverride func() (*api.BricksAPI, error)

	// Cached Users API client and its mutex
	cachedUsersAPI *api.UsersAPI
	usersAPIMutex  sync.RWMutex

	// Test hook: if non-nil, this function will be called instead of the
	// normal GetUsersAPI implementation. Tests may set this to inject a
	// fake UsersAPI or to simulate errors without touching disk/network.
	GetUsersAPIOverride func() (*api.UsersAPI, error)
)

// GetCredentials returns cached credentials, loading them from disk if necessary.
// It is safe for concurrent use.
func GetCredentials() (*api.Credentials, error) {
	// Fast path: return cached credentials if present
	credentialsMutex.RLock()
	if cachedCredentials != nil {
		defer credentialsMutex.RUnlock()
		return cachedCredentials, nil
	}
	credentialsMutex.RUnlock()

	// Slow path: load credentials and cache them
	credentialsMutex.Lock()
	defer credentialsMutex.Unlock()

	// Double-check in case another goroutine loaded while we waited for lock
	if cachedCredentials != nil {
		return cachedCredentials, nil
	}

	credentialsFile := brickscliutils.CredentialsDefaultPath()
	credentials, err := api.ImportCredentials(credentialsFile)
	if err != nil {
		return nil, err
	}

	cachedCredentials = credentials
	return cachedCredentials, nil
}

// GetBricksAPI returns a cached BricksAPI client, creating it if necessary
// It is safe for concurrent use.
func GetBricksAPI() (*api.BricksAPI, error) {
	// If tests have installed an override, use that. This lets tests avoid
	// creating real HTTP clients or reading credentials from disk.
	if GetBricksAPIOverride != nil {
		return GetBricksAPIOverride()
	}

	// Try to return cached API client first
	bricksAPIMutex.RLock()
	if cachedBricksAPI != nil {
		defer bricksAPIMutex.RUnlock()
		return cachedBricksAPI, nil
	}
	bricksAPIMutex.RUnlock()

	// Need to create API client for the first time
	bricksAPIMutex.Lock()
	defer bricksAPIMutex.Unlock()

	// Double-check in case another goroutine created while we waited for lock
	if cachedBricksAPI != nil {
		return cachedBricksAPI, nil
	}

	// Create HTTP client with timeout
	client := &http.Client{
		Timeout: time.Second * 5,
	}

	// Ensure credentials are available (GetCredentials handles caching)
	credentials, err := GetCredentials()
	if err != nil {
		return nil, err
	}

	// Create and cache API client
	cachedBricksAPI = api.NewBricksAPI(client, credentials.APIKey, false)
	return cachedBricksAPI, nil
}

// GetUsersAPI returns a cached UsersAPI client, creating it if necessary
// It is safe for concurrent use.
func GetUsersAPI() (*api.UsersAPI, error) {
	// If tests have installed an override, use that. This lets tests avoid
	// creating real HTTP clients or reading credentials from disk.
	if GetUsersAPIOverride != nil {
		return GetUsersAPIOverride()
	}
	// Fast path: return cached users API if present
	usersAPIMutex.RLock()
	if cachedUsersAPI != nil {
		defer usersAPIMutex.RUnlock()
		return cachedUsersAPI, nil
	}
	usersAPIMutex.RUnlock()

	// Slow path: create and cache the users API
	usersAPIMutex.Lock()
	defer usersAPIMutex.Unlock()

	// Double-check in case another goroutine created while we waited for lock
	if cachedUsersAPI != nil {
		return cachedUsersAPI, nil
	}

	// Create HTTP client with timeout
	client := &http.Client{
		Timeout: time.Second * 5,
	}

	// Ensure credentials are available (GetCredentials handles caching)
	credentials, err := GetCredentials()
	if err != nil {
		return nil, err
	}

	// Create and cache Users API client
	cachedUsersAPI = api.NewUsersAPI(client, credentials, false)
	return cachedUsersAPI, nil
}
