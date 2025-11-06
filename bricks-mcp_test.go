package main_test

import (
	"testing"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/wendehals/bricks-mcp/tools/apitools"
	"github.com/wendehals/bricks-mcp/tools/services"
)

// TestToolRegistration is a lightweight wrapper test that verifies all tool
// functions can be registered with an MCP server. This ensures the tool
// wrappers are exported and their signatures match the MCP expectations.
//
// The test does NOT call external APIs or run the server transport; it only
// constructs a server and registers the tools to catch compile-time/signature
// issues.
func TestToolRegistration(t *testing.T) {
	server := mcp.NewServer(&mcp.Implementation{Name: "bricks-mcp-test", Version: "v0"}, nil)

	// Register all known tools. If any of the tool functions have the wrong
	// signature or are not exported, this test will fail to compile.
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

	mcp.AddTool(server, &mcp.Tool{Name: services.RunScriptToolName, Description: services.RunScriptToolDescription}, services.RunScript)

	// If we reached this point the registration succeeded.
	_ = server
}
