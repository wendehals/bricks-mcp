package services

import (
	"context"
	"testing"

	"github.com/modelcontextprotocol/go-sdk/mcp"
)

func TestRunScript(t *testing.T) {
	t.Run("with empty script content", func(t *testing.T) {
		req := &mcp.CallToolRequest{}
		input := RunScriptInput{
			Script:  "",
			Verbose: false,
		}

		_, _, err := RunScript(context.Background(), req, input)

		if err == nil {
			t.Fatalf("Expected error for empty script content, got nil")
		}

		if err.Error() != "script is required" {
			t.Errorf("Expected specific error message, got: %v", err)
		}
	})

	t.Run("with valid input parameters", func(t *testing.T) {
		req := &mcp.CallToolRequest{}
		input := RunScriptInput{
			Script:  "allParts",
			Verbose: true,
		}

		// Note: This test verifies the function doesn't crash with valid parameters
		// The actual script execution may fail, but that's handled by BricksScript internally
		_, result, err := RunScript(context.Background(), req, input)

		// We expect no error from our function wrapper, even if the script itself fails internally
		if err != nil {
			t.Fatalf("Expected no error from RunScript function wrapper, got: %v", err)
		}

		resultMap, ok := result.(map[string]interface{})
		if !ok {
			t.Fatalf("Expected result to be a map[string]interface{}, got: %T", result)
		}

		if resultMap["script"] != input.Script {
			t.Errorf("Expected script to be %s, got %v", input.Script, resultMap["script"])
		}

		if resultMap["verbose"] != input.Verbose {
			t.Errorf("Expected verbose to be %v, got %v", input.Verbose, resultMap["verbose"])
		}

		if resultMap["status"] != "executed" {
			t.Errorf("Expected status to be 'executed', got %v", resultMap["status"])
		}
	})

	t.Run("constants are properly defined", func(t *testing.T) {
		if RunScriptToolName == "" {
			t.Error("RunScriptToolName should not be empty")
		}

		if RunScriptToolDescription == "" {
			t.Error("RunScriptToolDescription should not be empty")
		}

		expectedName := "run-script"
		if RunScriptToolName != expectedName {
			t.Errorf("Expected tool name to be %s, got %s", expectedName, RunScriptToolName)
		}
	})
}
