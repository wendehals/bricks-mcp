package apitools

import (
	"context"
	"fmt"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/wendehals/bricks-cli/model"
	"github.com/wendehals/bricks-mcp/utils"
)

const (
	GetUserSetListToolName        = "get-user-set-list"
	GetUserSetListToolDescription = "get a specific set list from the authenticated user's collection at Rebrickable©"
)

type GetUserSetListInput struct {
	SetListID uint `json:"set_list_id" title:"Set List ID" description:"The ID of the set list to retrieve" example:"12345" required:"true"`
}

// GetUserSetList returns a specific set list from the authenticated user
func GetUserSetList(ctx context.Context, req *mcp.CallToolRequest, input GetUserSetListInput) (*mcp.CallToolResult, *model.SetList, error) {
	usersAPI, err := utils.GetUsersAPI()
	if err != nil {
		return nil, nil, err
	}

	setList := usersAPI.GetSetList(input.SetListID)
	if setList == nil {
		return nil, nil, fmt.Errorf("failed to retrieve set list with ID %d", input.SetListID)
	}

	return nil, setList, nil
}
