package services

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/wendehals/bricks-cli/model"
	bricksservices "github.com/wendehals/bricks-cli/services"
)

const (
	MergeCollectionsToolName        = "merge-collections"
	MergeCollectionsToolDescription = "merge parts in a collection by color, alternates, molds, or prints (modifies the collection in place)"
)

type MergeCollectionsInput struct {
	CollectionJSON string `json:"collection_json" title:"Collection JSON" description:"The collection to merge, as JSON string" required:"true"`
	Mode           string `json:"mode" title:"Merge Mode" description:"Merge mode: 'c' (color), 'a' (alternates), 'm' (molds), 'p' (prints), or combination like 'amp'" example:"c" required:"true"`
}

func MergeCollections(ctx context.Context, req *mcp.CallToolRequest, input MergeCollectionsInput) (*mcp.CallToolResult, model.Collection, error) {
	var collection model.Collection
	if err := json.Unmarshal([]byte(input.CollectionJSON), &collection); err != nil {
		return nil, model.Collection{}, fmt.Errorf("failed to parse collection JSON: %w", err)
	}

	mode := bricksservices.ModeToUInt8(input.Mode)
	if mode == 0 {
		return nil, model.Collection{}, fmt.Errorf("invalid mode '%s': use 'c' (color), 'a' (alternates), 'm' (molds), 'p' (prints), or combination", input.Mode)
	}

	bricksservices.Merge(&collection, mode)

	return nil, collection, nil
}
