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
	BuildToolName        = "build"
	BuildToolDescription = "build a collection by matching needed parts with provided parts, allowing substitutions based on mode"
)

type BuildInput struct {
	NeededCollectionJSON   string `json:"needed_collection_json" title:"Needed Collection JSON" description:"The collection of parts needed for the build, as JSON string" required:"true"`
	ProvidedCollectionJSON string `json:"provided_collection_json" title:"Provided Collection JSON" description:"The collection of parts available to use, as JSON string" required:"true"`
	Mode                   string `json:"mode" title:"Build Mode" description:"Build mode: 'c' (allow different colors), 'a' (alternates), 'm' (molds), 'p' (prints), or combination like 'camp'" example:"" required:"false"`
}

func Build(ctx context.Context, req *mcp.CallToolRequest, input BuildInput) (*mcp.CallToolResult, model.BuildCollection, error) {
	var neededCollection model.Collection
	if err := json.Unmarshal([]byte(input.NeededCollectionJSON), &neededCollection); err != nil {
		return nil, model.BuildCollection{}, fmt.Errorf("failed to parse needed collection JSON: %w", err)
	}

	var providedCollection model.Collection
	if err := json.Unmarshal([]byte(input.ProvidedCollectionJSON), &providedCollection); err != nil {
		return nil, model.BuildCollection{}, fmt.Errorf("failed to parse provided collection JSON: %w", err)
	}

	mode := bricksservices.ModeToUInt8(input.Mode)

	buildCollection := bricksservices.Build(&neededCollection, &providedCollection, mode)

	return nil, *buildCollection, nil
}
