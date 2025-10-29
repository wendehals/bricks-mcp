package main

import (
	"context"
	"log"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/wendehals/bricks-mcp/tools"
)

func main() {
	server := mcp.NewServer(&mcp.Implementation{Name: "bricks-mcp", Version: "v0.0.1"}, nil)
	mcp.AddTool(server, &mcp.Tool{Name: tools.GetSetToolName, Description: tools.GetSetToolDescription}, tools.GetSet)

	// Run the server over stdin/stdout, until the client disconnects.
	if err := server.Run(context.Background(), &mcp.StdioTransport{}); err != nil {
		log.Fatal(err)
	}
}
