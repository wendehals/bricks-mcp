package apitools

import (
	"context"
	"fmt"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/wendehals/bricks-cli/model"
	"github.com/wendehals/bricks-mcp/utils"
)

const (
	GetUserSetListsToolName        = "get-user-set-lists"
	GetUserSetListsToolDescription = "get the user's set lists from Rebrickable©"
)

type GetUserSetListsInput struct {
	// No input parameters needed
}

// GetUserSetLists returns all set lists from the authenticated user
func GetUserSetLists(ctx context.Context, req *mcp.CallToolRequest, input GetUserSetListsInput) (*mcp.CallToolResult, *model.SetLists, error) {
	usersAPI, err := utils.GetUsersAPI()
	if err != nil {
		return nil, nil, err
	}

	setLists := usersAPI.GetSetLists()
	if setLists == nil {
		return nil, nil, fmt.Errorf("failed to retrieve user's set lists")
	}

	return nil, setLists, nil
}
