package server

import "github.com/mark3labs/mcp-go/server"

func New() *server.MCPServer {
	s := server.NewMCPServer(
		"caesar Cipher",
		"1.0.0",
	)
	// Add tool here
	// name: "caesar_rotate"
	// description: "Rotate a string by a given number of positions. It is used to encrypt or decrypt text of caesar Cipher."
	return s
}

func rotateHandler() {
	// TODO: implement
}
