package tools

import (
	"context"
	"fmt"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/wendehals/bricks-cli/model"
	"github.com/wendehals/bricks-mcp/utils"
)

const (
	GetSetToolName        = "get-set"
	GetSetToolDescription = "get the details of a set of bricks"
)

type GetSetInput struct {
	SetNumber string `json:"set_number" title:"Set Number" description:"The set number of the set to retrieve" example:"75192-1" required:"true"`
}

func GetSet(ctx context.Context, req *mcp.CallToolRequest, input GetSetInput) (*mcp.CallToolResult, model.Set, error) {
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
