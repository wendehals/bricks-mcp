package prompts

import (
	"context"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

const (
	GetUserSetPromptName        = "get-user-set"
	GetUserSetPromptTitle       = "Get User Set"
	GetUserSetPromptDescription = "Example Bricks script to retrieve a user's set and print it to stdout"
)

func GetUserSetPrompt() *mcp.Prompt {
	return &mcp.Prompt{
		Name:        GetUserSetPromptName,
		Title:       GetUserSetPromptTitle,
		Description: GetUserSetPromptDescription,
		Arguments: []*mcp.PromptArgument{
			{
				Name:        "set_number",
				Title:       "Set Number",
				Description: "The LEGO set number to retrieve (e.g., 918-1)",
				Required:    false,
			},
		},
	}
}

func GetUserSetHandler(ctx context.Context, req *mcp.GetPromptRequest) (*mcp.GetPromptResult, error) {
	setNumber := "918-1"
	if req.Params.Arguments != nil {
		if setNum, ok := req.Params.Arguments["set_number"]; ok && setNum != "" {
			setNumber = setNum
		}
	}

	script := `print(userSet("` + setNumber + `"))`

	return &mcp.GetPromptResult{
		Description: "Bricks script that retrieves user set " + setNumber + " and prints it to stdout",
		Messages: []*mcp.PromptMessage{
			{
				Role: "user",
				Content: &mcp.TextContent{
					Text: "Use the run-script tool with this Bricks script to retrieve set " + setNumber + " from the user's collection:\n\n" + script,
				},
			},
		},
	}, nil
}
