package server

import (
	"context"
	"fmt"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func New() *server.MCPServer {
	s := server.NewMCPServer(
		"caesar Cipher",
		"1.0.0",
	)

	tool := mcp.NewTool("ceaser_rotate",
		mcp.WithDescription("Rotate a string by a given number of positions. It is used to encrypt or decrypt text of Ceaser Cipher."),
		mcp.WithNumber("shift",
			mcp.Description("Number of positions to rotate")),
		mcp.WithString("text",
			mcp.Description("Text to rotate")),
	)
	s.AddTool(tool, rotateHandler)

	return s
}

func rotateHandler(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	shift, ok := request.Params.Arguments["shift"].(float64)
	if !ok {
		return nil, fmt.Errorf("shift must be an integer")
	}
	text, ok := request.Params.Arguments["text"].(string)
	if !ok {
		return nil, fmt.Errorf("text must be a string")
	}
	return mcp.NewToolResultText(ceaser.RotN(text, int(shift))), nil
}
