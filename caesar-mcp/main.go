package main

import (
	"fmt"

	"github.com/mark3labs/mcp-go/server"
	caesarserver "github.com/syumai/go-mcp-hands-on/caesar-mcp/server"
)

func main() {
	s := caesarserver.New()
	if err := server.ServeStdio(s); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}
