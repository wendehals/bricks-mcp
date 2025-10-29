package utils

import (
	"net/http"
	"sync"
	"time"

	"github.com/wendehals/bricks-cli/api"
	brickscliutils "github.com/wendehals/bricks-cli/utils"
)

var (
	// Cached API client and credentials to avoid recreating/reloading on every request
	cachedBricksAPI   *api.BricksAPI
	cachedCredentials *api.Credentials

	// Mutexes for API client and credentials
	apiMutex   sync.RWMutex
	credsMutex sync.RWMutex
)

// GetCredentials returns cached credentials, loading them from disk if necessary.
// It is safe for concurrent use.
func GetCredentials() (*api.Credentials, error) {
	// Fast path: return cached credentials if present
	credsMutex.RLock()
	if cachedCredentials != nil {
		defer credsMutex.RUnlock()
		return cachedCredentials, nil
	}
	credsMutex.RUnlock()

	// Slow path: load credentials and cache them
	credsMutex.Lock()
	defer credsMutex.Unlock()

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
	// Try to return cached API client first
	apiMutex.RLock()
	if cachedBricksAPI != nil {
		defer apiMutex.RUnlock()
		return cachedBricksAPI, nil
	}
	apiMutex.RUnlock()

	// Need to create API client for the first time
	apiMutex.Lock()
	defer apiMutex.Unlock()

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
