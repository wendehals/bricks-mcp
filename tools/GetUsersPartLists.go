package tools

import (
	"context"
	"fmt"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/wendehals/bricks-cli/model"
	"github.com/wendehals/bricks-mcp/utils"
)

const (
	GetUsersPartListsToolName        = "get-users-part-lists"
	GetUsersPartListsToolDescription = "get all part lists of the authenticated user's collection at Rebrickable©"
)

type GetUsersPartListsInput struct {
	// No input parameters needed
}

// GetUsersPartLists returns the part lists from the authenticated user
func GetUsersPartLists(ctx context.Context, req *mcp.CallToolRequest, input GetUsersPartListsInput) (*mcp.CallToolResult, *model.PartLists, error) {
	usersAPI, err := utils.GetUsersAPI()
	if err != nil {
		return nil, nil, err
	}

	partLists := usersAPI.GetPartLists()
	if partLists == nil {
		return nil, nil, fmt.Errorf("failed to retrieve user's part lists")
	}

	return nil, partLists, nil
}
