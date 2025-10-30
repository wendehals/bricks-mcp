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
	MergeAllCollectionsToolName        = "merge-all-collections"
	MergeAllCollectionsToolDescription = "merge all given collections to a new single collection"
)

type MergeAllCollectionsInput struct {
	CollectionsJSON string `json:"collections_json" title:"Collections JSON" description:"Array of collections to merge, as JSON string" required:"true"`
}

func MergeAllCollections(ctx context.Context, req *mcp.CallToolRequest, input MergeAllCollectionsInput) (*mcp.CallToolResult, model.Collection, error) {
	var collections []model.Collection
	if err := json.Unmarshal([]byte(input.CollectionsJSON), &collections); err != nil {
		return nil, model.Collection{}, fmt.Errorf("failed to parse collections JSON: %w", err)
	}

	if len(collections) == 0 {
		return nil, model.Collection{}, fmt.Errorf("no collections provided to merge")
	}

	mergedCollection := bricksservices.MergeAllCollections(collections)

	return nil, *mergedCollection, nil
}
