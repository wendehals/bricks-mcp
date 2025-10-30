package main

import (
	"context"
	"log"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/wendehals/bricks-mcp/tools/apitools"
	"github.com/wendehals/bricks-mcp/tools/services"
)

func main() {
	server := mcp.NewServer(&mcp.Implementation{Name: "bricks-mcp", Version: "v0.0.6"}, nil)
	mcp.AddTool(server, &mcp.Tool{Name: apitools.GetSetOverviewToolName, Description: apitools.GetSetOverviewToolDescription}, apitools.GetSetOverview)
	mcp.AddTool(server, &mcp.Tool{Name: apitools.GetSetPartsToolName, Description: apitools.GetSetPartsToolDescription}, apitools.GetSetParts)
	mcp.AddTool(server, &mcp.Tool{Name: apitools.GetSetIncludingPartsToolName, Description: apitools.GetSetIncludingPartsToolDescription}, apitools.GetSetIncludingParts)
	mcp.AddTool(server, &mcp.Tool{Name: apitools.GetAllUserPartsToolName, Description: apitools.GetAllUserPartsToolDescription}, apitools.GetAllUserParts)
	mcp.AddTool(server, &mcp.Tool{Name: apitools.GetAllUserSetsToolName, Description: apitools.GetAllUserSetsToolDescription}, apitools.GetAllUserSets)
	mcp.AddTool(server, &mcp.Tool{Name: apitools.GetUserSetListsToolName, Description: apitools.GetUserSetListsToolDescription}, apitools.GetUserSetLists)
	mcp.AddTool(server, &mcp.Tool{Name: apitools.GetUserSetListToolName, Description: apitools.GetUserSetListToolDescription}, apitools.GetUserSetList)
	mcp.AddTool(server, &mcp.Tool{Name: apitools.GetSetsOfUserSetListToolName, Description: apitools.GetSetsOfUserSetListToolDescription}, apitools.GetSetsOfUserSetList)
	mcp.AddTool(server, &mcp.Tool{Name: apitools.GetUsersPartListsToolName, Description: apitools.GetUsersPartListsToolDescription}, apitools.GetUsersPartLists)

	mcp.AddTool(server, &mcp.Tool{Name: services.MergeCollectionsToolName, Description: services.MergeCollectionsToolDescription}, services.MergeCollections)
	mcp.AddTool(server, &mcp.Tool{Name: services.MergeAllCollectionsToolName, Description: services.MergeAllCollectionsToolDescription}, services.MergeAllCollections)
	mcp.AddTool(server, &mcp.Tool{Name: services.BuildToolName, Description: services.BuildToolDescription}, services.Build)

	// Run the server over stdin/stdout, until the client disconnects.
	if err := server.Run(context.Background(), &mcp.StdioTransport{}); err != nil {
		log.Fatal(err)
	}
}
