package services

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/wendehals/bricks-cli/model"
)

// Helper function to create a test collection with parts
func createTestCollection(parts []model.Part) string {
	coll := model.Collection{
		Parts: parts,
		Sets:  []model.Set{},
	}
	jsonBytes, _ := json.Marshal(coll)
	return string(jsonBytes)
}

// Helper function to create test parts
func createTestPart(partNum string, colorID int, colorName string, quantity int) model.Part {
	return model.Part{
		Shape: model.Shape{
			Number: partNum,
			Name:   "Test Part " + partNum,
		},
		Color: model.Color{
			ID:   colorID,
			Name: colorName,
		},
		Quantity: quantity,
		IsSpare:  false,
	}
}

func TestMergeCollections_ColorMode(t *testing.T) {
	// Create a collection with same parts in different colors
	parts := []model.Part{
		createTestPart("3001", 1, "Blue", 5),
		createTestPart("3001", 4, "Red", 3),
		createTestPart("3002", 1, "Blue", 2),
	}
	collJSON := createTestCollection(parts)

	input := MergeCollectionsInput{
		CollectionJSON: collJSON,
		Mode:           "c", // Color mode
	}

	_, result, err := MergeCollections(context.Background(), nil, input)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// After merging by color, we should have 2 unique part numbers (3001 and 3002)
	if len(result.Parts) != 2 {
		t.Fatalf("expected 2 merged parts, got %d", len(result.Parts))
	}

	// Find the merged 3001 part
	var merged3001 *model.Part
	for i := range result.Parts {
		if result.Parts[i].Shape.Number == "3001" {
			merged3001 = &result.Parts[i]
			break
		}
	}

	if merged3001 == nil {
		t.Fatal("expected to find merged part 3001")
	}

	// Should have combined quantity: 5 + 3 = 8
	if merged3001.Quantity != 8 {
		t.Fatalf("expected quantity 8 for merged 3001, got %d", merged3001.Quantity)
	}

	// Comment should indicate merge by color
	if result.Comment != "Merged by color" {
		t.Fatalf("expected comment 'Merged by color', got '%s'", result.Comment)
	}

	// Sets should be cleared
	if len(result.Sets) != 0 {
		t.Fatalf("expected sets to be cleared, got %d", len(result.Sets))
	}
}

func TestMergeCollections_InvalidMode(t *testing.T) {
	parts := []model.Part{
		createTestPart("3001", 1, "Blue", 5),
	}
	collJSON := createTestCollection(parts)

	input := MergeCollectionsInput{
		CollectionJSON: collJSON,
		Mode:           "xyz", // Invalid mode
	}

	_, _, err := MergeCollections(context.Background(), nil, input)
	if err == nil {
		t.Fatal("expected error for invalid mode, got nil")
	}
}

func TestMergeCollections_InvalidJSON(t *testing.T) {
	input := MergeCollectionsInput{
		CollectionJSON: "invalid json {{{",
		Mode:           "c",
	}

	_, _, err := MergeCollections(context.Background(), nil, input)
	if err == nil {
		t.Fatal("expected error for invalid JSON, got nil")
	}
}

func TestMergeAllCollections(t *testing.T) {
	// Create two collections
	coll1 := model.Collection{
		Parts: []model.Part{
			createTestPart("3001", 1, "Blue", 5),
			createTestPart("3002", 4, "Red", 3),
		},
		Sets: []model.Set{},
	}

	coll2 := model.Collection{
		Parts: []model.Part{
			createTestPart("3001", 1, "Blue", 2),
			createTestPart("3003", 7, "Gray", 4),
		},
		Sets: []model.Set{},
	}

	collections := []model.Collection{coll1, coll2}
	collectionsJSON, _ := json.Marshal(collections)

	input := MergeAllCollectionsInput{
		CollectionsJSON: string(collectionsJSON),
	}

	_, result, err := MergeAllCollections(context.Background(), nil, input)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Should have 3 unique parts: 3001 (from both), 3002, 3003
	if len(result.Parts) != 3 {
		t.Fatalf("expected 3 parts in merged collection, got %d", len(result.Parts))
	}

	// Find the 3001 part and verify quantity was combined
	var part3001 *model.Part
	for i := range result.Parts {
		if result.Parts[i].Shape.Number == "3001" {
			part3001 = &result.Parts[i]
			break
		}
	}

	if part3001 == nil {
		t.Fatal("expected to find part 3001 in merged collection")
	}

	// Should have combined quantity: 5 + 2 = 7
	if part3001.Quantity != 7 {
		t.Fatalf("expected quantity 7 for part 3001, got %d", part3001.Quantity)
	}
}

func TestMergeAllCollections_EmptyArray(t *testing.T) {
	collectionsJSON := "[]"

	input := MergeAllCollectionsInput{
		CollectionsJSON: collectionsJSON,
	}

	_, _, err := MergeAllCollections(context.Background(), nil, input)
	if err == nil {
		t.Fatal("expected error for empty collections array, got nil")
	}
}

func TestMergeAllCollections_InvalidJSON(t *testing.T) {
	input := MergeAllCollectionsInput{
		CollectionsJSON: "not valid json [[[",
	}

	_, _, err := MergeAllCollections(context.Background(), nil, input)
	if err == nil {
		t.Fatal("expected error for invalid JSON, got nil")
	}
}

func TestBuild_ExactMatch(t *testing.T) {
	// Needed collection: need 5 blue 3001 bricks
	neededParts := []model.Part{
		createTestPart("3001", 1, "Blue", 5),
	}
	neededJSON := createTestCollection(neededParts)

	// Provided collection: have 10 blue 3001 bricks
	providedParts := []model.Part{
		createTestPart("3001", 1, "Blue", 10),
	}
	providedJSON := createTestCollection(providedParts)

	input := BuildInput{
		NeededCollectionJSON:   neededJSON,
		ProvidedCollectionJSON: providedJSON,
		Mode:                   "", // No substitutions
	}

	_, result, err := Build(context.Background(), nil, input)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Should have 1 mapping
	if len(result.Mapping) != 1 {
		t.Fatalf("expected 1 mapping, got %d", len(result.Mapping))
	}

	mapping := result.Mapping[0]

	// Original should be the needed part
	if mapping.Original.Shape.Number != "3001" {
		t.Fatalf("expected original part 3001, got %s", mapping.Original.Shape.Number)
	}

	// Quantity should be 0 (all matched)
	if mapping.Quantity != 0 {
		t.Fatalf("expected quantity 0 (fully matched), got %d", mapping.Quantity)
	}

	// Should have 1 substitute (exact match)
	if len(mapping.Substitutes) != 1 {
		t.Fatalf("expected 1 substitute, got %d", len(mapping.Substitutes))
	}

	// Substitute should have quantity 5
	if mapping.Substitutes[0].Quantity != 5 {
		t.Fatalf("expected substitute quantity 5, got %d", mapping.Substitutes[0].Quantity)
	}
}

func TestBuild_PartialMatch(t *testing.T) {
	// Needed collection: need 10 blue 3001 bricks
	neededParts := []model.Part{
		createTestPart("3001", 1, "Blue", 10),
	}
	neededJSON := createTestCollection(neededParts)

	// Provided collection: have only 6 blue 3001 bricks
	providedParts := []model.Part{
		createTestPart("3001", 1, "Blue", 6),
	}
	providedJSON := createTestCollection(providedParts)

	input := BuildInput{
		NeededCollectionJSON:   neededJSON,
		ProvidedCollectionJSON: providedJSON,
		Mode:                   "",
	}

	_, result, err := Build(context.Background(), nil, input)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	mapping := result.Mapping[0]

	// Quantity should be 4 (still need 4 more)
	if mapping.Quantity != 4 {
		t.Fatalf("expected quantity 4 (missing), got %d", mapping.Quantity)
	}

	// Should have 1 substitute with quantity 6
	if len(mapping.Substitutes) != 1 || mapping.Substitutes[0].Quantity != 6 {
		t.Fatalf("expected 1 substitute with quantity 6")
	}
}

func TestBuild_ColorMode(t *testing.T) {
	// Needed collection: need 5 blue 3001 bricks
	neededParts := []model.Part{
		createTestPart("3001", 1, "Blue", 5),
	}
	neededJSON := createTestCollection(neededParts)

	// Provided collection: have red 3001 bricks (different color)
	providedParts := []model.Part{
		createTestPart("3001", 4, "Red", 10),
	}
	providedJSON := createTestCollection(providedParts)

	// Without color mode, should not match
	inputNoColor := BuildInput{
		NeededCollectionJSON:   neededJSON,
		ProvidedCollectionJSON: providedJSON,
		Mode:                   "",
	}

	_, resultNoColor, err := Build(context.Background(), nil, inputNoColor)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Should have quantity 5 unmatched
	if resultNoColor.Mapping[0].Quantity != 5 {
		t.Fatalf("expected quantity 5 unmatched without color mode, got %d", resultNoColor.Mapping[0].Quantity)
	}

	// With color mode, should match red as substitute
	inputWithColor := BuildInput{
		NeededCollectionJSON:   neededJSON,
		ProvidedCollectionJSON: providedJSON,
		Mode:                   "c", // Allow different colors
	}

	_, resultWithColor, err := Build(context.Background(), nil, inputWithColor)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	// Should have quantity 0 (all matched with red substitute)
	if resultWithColor.Mapping[0].Quantity != 0 {
		t.Fatalf("expected quantity 0 with color mode, got %d", resultWithColor.Mapping[0].Quantity)
	}

	// Should have substitutes
	if len(resultWithColor.Mapping[0].Substitutes) == 0 {
		t.Fatal("expected substitutes with color mode")
	}
}

func TestBuild_InvalidNeededJSON(t *testing.T) {
	input := BuildInput{
		NeededCollectionJSON:   "invalid json",
		ProvidedCollectionJSON: "{}",
		Mode:                   "",
	}

	_, _, err := Build(context.Background(), nil, input)
	if err == nil {
		t.Fatal("expected error for invalid needed collection JSON, got nil")
	}
}

func TestBuild_InvalidProvidedJSON(t *testing.T) {
	neededJSON := createTestCollection([]model.Part{})
	input := BuildInput{
		NeededCollectionJSON:   neededJSON,
		ProvidedCollectionJSON: "invalid json",
		Mode:                   "",
	}

	_, _, err := Build(context.Background(), nil, input)
	if err == nil {
		t.Fatal("expected error for invalid provided collection JSON, got nil")
	}
}
