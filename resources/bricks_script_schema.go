package resources

import (
	"context"
	_ "embed"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

const (
	BricksScriptSchemaResourceURI         = "bricks://schema/bricks-script.json"
	BricksScriptSchemaResourceName        = "bricks-script-schema"
	BricksScriptSchemaResourceTitle       = "BricksScript JSON Schema"
	BricksScriptSchemaResourceDescription = "JSON Schema for the BricksScript scripting language used in bricks-cli"
	BricksScriptSchemaResourceMIMEType    = "application/schema+json"
)

//go:embed bricks-schema.json
var bricksScriptSchemaContent string

func BricksScriptSchemaResource() *mcp.Resource {
	return &mcp.Resource{
		URI:         BricksScriptSchemaResourceURI,
		Name:        BricksScriptSchemaResourceName,
		Title:       BricksScriptSchemaResourceTitle,
		Description: BricksScriptSchemaResourceDescription,
		MIMEType:    BricksScriptSchemaResourceMIMEType,
		Size:        int64(len(bricksScriptSchemaContent)),
	}
}

func BricksScriptSchemaHandler(ctx context.Context, req *mcp.ReadResourceRequest) (*mcp.ReadResourceResult, error) {
	return &mcp.ReadResourceResult{
		Contents: []*mcp.ResourceContents{
			{
				URI:      BricksScriptSchemaResourceURI,
				MIMEType: BricksScriptSchemaResourceMIMEType,
				Text:     bricksScriptSchemaContent,
			},
		},
	}, nil
}
