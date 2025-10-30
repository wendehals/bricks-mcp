package apitools

import (
	"context"
	"fmt"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/wendehals/bricks-cli/model"
	"github.com/wendehals/bricks-mcp/utils"
)

const (
	GetSetsOfUserSetListToolName        = "get-sets-of-user-set-list"
	GetSetsOfUserSetListToolDescription = "get the sets from a specific set list in the authenticated user's collection at Rebrickable©"
)

type GetSetsOfUserSetListInput struct {
	SetListID uint `json:"set_list_id" title:"Set List ID" description:"The ID of the set list to retrieve sets from" example:"12345" required:"true"`
}

// GetSetsOfUserSetList returns all sets from a specific set list of the authenticated user
func GetSetsOfUserSetList(ctx context.Context, req *mcp.CallToolRequest, input GetSetsOfUserSetListInput) (*mcp.CallToolResult, *model.UserSets, error) {
	usersAPI, err := utils.GetUsersAPI()
	if err != nil {
		return nil, nil, err
	}

	sets := usersAPI.GetSetListSets(input.SetListID)
	if sets == nil {
		return nil, nil, fmt.Errorf("failed to retrieve sets from set list with ID %d", input.SetListID)
	}

	return nil, sets, nil
}
