package main

import (
	"context"
	"log"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/wendehals/bricks-mcp/tools/apitools"
)

func main() {
	server := mcp.NewServer(&mcp.Implementation{Name: "bricks-mcp", Version: "v0.0.5"}, nil)
	mcp.AddTool(server, &mcp.Tool{Name: apitools.GetSetOverviewToolName, Description: apitools.GetSetOverviewToolDescription}, apitools.GetSetOverview)
	mcp.AddTool(server, &mcp.Tool{Name: apitools.GetSetPartsToolName, Description: apitools.GetSetPartsToolDescription}, apitools.GetSetParts)
	mcp.AddTool(server, &mcp.Tool{Name: apitools.GetSetIncludingPartsToolName, Description: apitools.GetSetIncludingPartsToolDescription}, apitools.GetSetIncludingParts)
	mcp.AddTool(server, &mcp.Tool{Name: apitools.GetAllUserPartsToolName, Description: apitools.GetAllUserPartsToolDescription}, apitools.GetAllUserParts)
	mcp.AddTool(server, &mcp.Tool{Name: apitools.GetAllUserSetsToolName, Description: apitools.GetAllUserSetsToolDescription}, apitools.GetAllUserSets)
	mcp.AddTool(server, &mcp.Tool{Name: apitools.GetUserSetListsToolName, Description: apitools.GetUserSetListsToolDescription}, apitools.GetUserSetLists)
	mcp.AddTool(server, &mcp.Tool{Name: apitools.GetUserSetListToolName, Description: apitools.GetUserSetListToolDescription}, apitools.GetUserSetList)
	mcp.AddTool(server, &mcp.Tool{Name: apitools.GetSetsOfUserSetListToolName, Description: apitools.GetSetsOfUserSetListToolDescription}, apitools.GetSetsOfUserSetList)
	mcp.AddTool(server, &mcp.Tool{Name: apitools.GetUsersPartListsToolName, Description: apitools.GetUsersPartListsToolDescription}, apitools.GetUsersPartLists)

	// Run the server over stdin/stdout, until the client disconnects.
	if err := server.Run(context.Background(), &mcp.StdioTransport{}); err != nil {
		log.Fatal(err)
	}
}
