package main_test

import (
	"testing"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/wendehals/bricks-mcp/tools"
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
	mcp.AddTool(server, &mcp.Tool{Name: tools.GetSetOverviewToolName, Description: tools.GetSetOverviewToolDescription}, tools.GetSetOverview)
	mcp.AddTool(server, &mcp.Tool{Name: tools.GetSetPartsToolName, Description: tools.GetSetPartsToolDescription}, tools.GetSetParts)
	mcp.AddTool(server, &mcp.Tool{Name: tools.GetSetIncludingPartsToolName, Description: tools.GetSetIncludingPartsToolDescription}, tools.GetSetIncludingParts)
	mcp.AddTool(server, &mcp.Tool{Name: tools.GetAllUserPartsToolName, Description: tools.GetAllUserPartsToolDescription}, tools.GetAllUserParts)
	mcp.AddTool(server, &mcp.Tool{Name: tools.GetAllUserSetsToolName, Description: tools.GetAllUserSetsToolDescription}, tools.GetAllUserSets)
	mcp.AddTool(server, &mcp.Tool{Name: tools.GetUserSetListsToolName, Description: tools.GetUserSetListsToolDescription}, tools.GetUserSetLists)
	mcp.AddTool(server, &mcp.Tool{Name: tools.GetUserSetListToolName, Description: tools.GetUserSetListToolDescription}, tools.GetUserSetList)
	mcp.AddTool(server, &mcp.Tool{Name: tools.GetSetsOfUserSetListToolName, Description: tools.GetSetsOfUserSetListToolDescription}, tools.GetSetsOfUserSetList)
	mcp.AddTool(server, &mcp.Tool{Name: tools.GetUsersPartListsToolName, Description: tools.GetUsersPartListsToolDescription}, tools.GetUsersPartLists)

	// If we reached this point the registration succeeded.
	_ = server
}
