package main

import (
	"context"
	"log"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/wendehals/bricks-mcp/tools"
)

func main() {
	server := mcp.NewServer(&mcp.Implementation{Name: "bricks-mcp", Version: "v0.0.3"}, nil)
	mcp.AddTool(server, &mcp.Tool{Name: tools.GetSetOverviewToolName, Description: tools.GetSetOverviewToolDescription}, tools.GetSetOverview)
	mcp.AddTool(server, &mcp.Tool{Name: tools.GetSetPartsToolName, Description: tools.GetSetPartsToolDescription}, tools.GetSetParts)
	mcp.AddTool(server, &mcp.Tool{Name: tools.GetSetIncludingPartsToolName, Description: tools.GetSetIncludingPartsToolDescription}, tools.GetSetIncludingParts)

	// Run the server over stdin/stdout, until the client disconnects.
	if err := server.Run(context.Background(), &mcp.StdioTransport{}); err != nil {
		log.Fatal(err)
	}
}
