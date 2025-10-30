package apitools

import (
	"context"
	"fmt"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/wendehals/bricks-cli/model"
	"github.com/wendehals/bricks-mcp/utils"
)

const (
	GetAllUserPartsToolName        = "get-all-user-parts"
	GetAllUserPartsToolDescription = "get all parts of the authenticated user's collection at Rebrickable©"
)

type GetAllUserPartsInput struct {
	// No input parameters needed
}

// GetAllUserParts returns the complete parts collection of the authenticated user
func GetAllUserParts(ctx context.Context, req *mcp.CallToolRequest, input GetAllUserPartsInput) (*mcp.CallToolResult, model.Collection, error) {
	usersAPI, err := utils.GetUsersAPI()
	if err != nil {
		return nil, model.Collection{}, err
	}

	collection := usersAPI.GetAllParts()
	if collection == nil {
		return nil, model.Collection{}, fmt.Errorf("failed to retrieve user's parts collection")
	}

	return nil, *collection, nil
}
