package tools

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/wendehals/bricks-cli/api"
	"github.com/wendehals/bricks-cli/model"
	"github.com/wendehals/bricks-cli/utils"
)

const (
	GetSetToolName        = "get-set"
	GetSetToolDescription = "get the details of a set of bricks"
)

type GetSetInput struct {
	SetNumber string `json:"set_number" title:"Set Number" description:"The set number of the set to retrieve" example:"75192-1" required:"true"`
}

func GetSet(ctx context.Context, req *mcp.CallToolRequest, input GetSetInput) (*mcp.CallToolResult, model.Set, error) {
	client := &http.Client{
		Timeout: time.Second * 5,
	}

	credentialsFile := utils.CredentialsDefaultPath()
	credentials, err := api.ImportCredentials(credentialsFile)
	if err != nil {
		log.Fatalf("no credentials file found: %v", err)
	}

	api := api.NewBricksAPI(client, credentials.APIKey, false)

	set := api.GetSet(input.SetNumber)
	if set == nil {
		log.Fatalf("set %s not found or API returned nil", input.SetNumber)
	}

	return nil, *set, nil
}
