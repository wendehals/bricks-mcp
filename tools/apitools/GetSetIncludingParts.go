package apitools

import (
	"context"
	"fmt"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/wendehals/bricks-cli/api"
	"github.com/wendehals/bricks-cli/model"
	"github.com/wendehals/bricks-mcp/utils"
)

const (
	GetSetIncludingPartsToolName        = "get-set-including-parts"
	GetSetIncludingPartsToolDescription = "get an overview about a set of bricks and its complete parts inventory (with caching)"
)

type GetSetIncludingPartsInput struct {
	SetNumber       string `json:"set_number" title:"Set Number" description:"The set number of the set to retrieve" example:"75192-1" required:"true"`
	IncludeMinifigs bool   `json:"include_minifigs,omitempty" title:"Include Minifigs" description:"Whether to include minifigure parts in the returned collection" example:"false"`
}

// GetSetIncludingParts returns both the set details and its parts collection with caching
func GetSetIncludingParts(ctx context.Context, req *mcp.CallToolRequest, input GetSetIncludingPartsInput) (*mcp.CallToolResult, model.Collection, error) {
	apiClient, err := utils.GetBricksAPI()
	if err != nil {
		return nil, model.Collection{}, err
	}

	collection := api.RetrieveSetParts(apiClient, input.SetNumber, input.IncludeMinifigs)
	if collection == nil {
		return nil, model.Collection{}, fmt.Errorf("parts of set number %s could not be found", input.SetNumber)
	}

	return nil, *collection, nil
}
