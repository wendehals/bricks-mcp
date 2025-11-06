package services

import (
	"context"
	"fmt"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/wendehals/bricks-cli/scripting"
	bricksutils "github.com/wendehals/bricks-mcp/utils"
)

const (
	RunScriptToolName        = "run-script"
	RunScriptToolDescription = "execute a Bricks script using the bricks-cli scripting engine"
)

type RunScriptInput struct {
	Script  string `json:"script" title:"Script" description:"The Bricks script to be executed" required:"true"`
	Verbose bool   `json:"verbose" title:"Verbose Output" description:"Enable verbose output during script execution" required:"false"`
}

func RunScript(ctx context.Context, req *mcp.CallToolRequest, input RunScriptInput) (*mcp.CallToolResult, interface{}, error) {
	credentials, err := bricksutils.GetCredentials()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get credentials: %w", err)
	}

	// Check if script content is provided
	if input.Script == "" {
		return nil, nil, fmt.Errorf("script is required")
	}

	bricksScript := scripting.NewBricksScript(credentials, input.Script, input.Verbose)

	// Execute the script - Note: This may log errors internally but doesn't return them
	bricksScript.Execute()

	result := map[string]interface{}{
		"script":  input.Script,
		"verbose": input.Verbose,
		"status":  "executed",
		"message": "Script execution completed (check logs for any execution errors)",
	}

	return nil, result, nil
}
