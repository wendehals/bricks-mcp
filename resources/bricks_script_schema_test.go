package resources

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func TestBricksScriptSchemaResource(t *testing.T) {
	t.Run("resource_metadata", func(t *testing.T) {
		resource := BricksScriptSchemaResource()

		if resource.URI != BricksScriptSchemaResourceURI {
			t.Errorf("Expected URI %s, got %s", BricksScriptSchemaResourceURI, resource.URI)
		}

		if resource.Name != BricksScriptSchemaResourceName {
			t.Errorf("Expected name %s, got %s", BricksScriptSchemaResourceName, resource.Name)
		}

		if resource.MIMEType != BricksScriptSchemaResourceMIMEType {
			t.Errorf("Expected MIME type %s, got %s", BricksScriptSchemaResourceMIMEType, resource.MIMEType)
		}

		if resource.Size <= 0 {
			t.Errorf("Expected size > 0, got %d", resource.Size)
		}
	})

	t.Run("handler_returns_valid_schema", func(t *testing.T) {
		req := &mcp.ReadResourceRequest{}
		result, err := BricksScriptSchemaHandler(context.Background(), req)

		if err != nil {
			t.Fatalf("Expected no error, got: %v", err)
		}

		if len(result.Contents) != 1 {
			t.Fatalf("Expected 1 content item, got %d", len(result.Contents))
		}

		content := result.Contents[0]

		if content.URI != BricksScriptSchemaResourceURI {
			t.Errorf("Expected URI %s, got %s", BricksScriptSchemaResourceURI, content.URI)
		}

		if content.MIMEType != BricksScriptSchemaResourceMIMEType {
			t.Errorf("Expected MIME type %s, got %s", BricksScriptSchemaResourceMIMEType, content.MIMEType)
		}

		if content.Text == "" {
			t.Error("Expected non-empty schema text")
		}

		// Verify it's valid JSON
		var schemaData map[string]interface{}
		if err := json.Unmarshal([]byte(content.Text), &schemaData); err != nil {
			t.Errorf("Schema is not valid JSON: %v", err)
		}

		// Verify it has expected schema properties
		if schemaData["$schema"] == nil {
			t.Error("Expected $schema property in JSON schema")
		}

		if schemaData["title"] == nil {
			t.Error("Expected title property in JSON schema")
		}
	})

	t.Run("constants_properly_defined", func(t *testing.T) {
		if BricksScriptSchemaResourceURI == "" {
			t.Error("BricksScriptSchemaResourceURI should not be empty")
		}

		if BricksScriptSchemaResourceName == "" {
			t.Error("BricksScriptSchemaResourceName should not be empty")
		}

		if BricksScriptSchemaResourceMIMEType == "" {
			t.Error("BricksScriptSchemaResourceMIMEType should not be empty")
		}

		expectedMIME := "application/schema+json"
		if BricksScriptSchemaResourceMIMEType != expectedMIME {
			t.Errorf("Expected MIME type %s, got %s", expectedMIME, BricksScriptSchemaResourceMIMEType)
		}
	})
}
