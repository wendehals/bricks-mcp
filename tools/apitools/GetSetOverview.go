package apitools

import (
	"context"
	"fmt"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/wendehals/bricks-cli/model"
	"github.com/wendehals/bricks-mcp/utils"
)

const (
	GetSetOverviewToolName        = "get-set-overview"
	GetSetOverviewToolDescription = "get an overview of a set of bricks"
)

type GetSetOverviewInput struct {
	SetNumber string `json:"set_number" title:"Set Number" description:"The set number of the set to retrieve" example:"75192-1" required:"true"`
}

func GetSetOverview(ctx context.Context, req *mcp.CallToolRequest, input GetSetOverviewInput) (*mcp.CallToolResult, model.Set, error) {
	apiClient, err := utils.GetBricksAPI()
	if err != nil {
		return nil, model.Set{}, err
	}

	set := apiClient.GetSet(input.SetNumber)
	if set == nil {
		return nil, model.Set{}, fmt.Errorf("set number %s could not be found", input.SetNumber)
	}

	return nil, *set, nil
}
