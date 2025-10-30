package apitools

import (
	"context"
	"fmt"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/wendehals/bricks-cli/model"
	"github.com/wendehals/bricks-mcp/utils"
)

const (
	GetAllUserSetsToolName        = "get-all-user-sets"
	GetAllUserSetsToolDescription = "get all sets from the authenticated user's collection at Rebrickable©"
)

type GetAllUserSetsInput struct {
	// No input parameters needed
}

// GetAllUserSets returns all sets from the authenticated user's collection
func GetAllUserSets(ctx context.Context, req *mcp.CallToolRequest, input GetAllUserSetsInput) (*mcp.CallToolResult, *model.UserSets, error) {
	usersAPI, err := utils.GetUsersAPI()
	if err != nil {
		return nil, nil, err
	}

	sets := usersAPI.GetUserSets()
	if sets == nil {
		return nil, nil, fmt.Errorf("failed to retrieve user's sets collection")
	}

	return nil, sets, nil
}
