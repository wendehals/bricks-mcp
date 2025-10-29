package tools

import (
	"context"
	"fmt"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/wendehals/bricks-cli/model"
	"github.com/wendehals/bricks-mcp/utils"
)

const (
	GetSetPartsToolName        = "get-set-parts"
	GetSetPartsToolDescription = "get the parts (inventory) of a set of bricks"
)

type GetSetPartsInput struct {
	SetNumber       string `json:"set_number" title:"Set Number" description:"The set number of the set to retrieve parts for" example:"75192-1" required:"true"`
	IncludeMinifigs bool   `json:"include_minifigs,omitempty" title:"Include Minifigs" description:"Whether to include minifigure parts in the returned collection" example:"false"`
}

// GetSetParts returns the parts collection for a set (calls bricks-cli API)
func GetSetParts(ctx context.Context, req *mcp.CallToolRequest, input GetSetPartsInput) (*mcp.CallToolResult, model.Collection, error) {
	apiClient, err := utils.GetBricksAPI()
	if err != nil {
		return nil, model.Collection{}, err
	}

	collection := apiClient.GetSetParts(input.SetNumber, input.IncludeMinifigs)
	if collection == nil {
		return nil, model.Collection{}, fmt.Errorf("parts of set number %s could not be found", input.SetNumber)
	}

	return nil, *collection, nil
}
