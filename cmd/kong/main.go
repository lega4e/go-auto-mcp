// Package main is the entry point for the mcp-auto Kong Go PDK plugin server.
// Build with: go build ./cmd/kong
//
// The binary serves dual roles:
//   - Kong plugin schema introspection: mcp-auto-kong --dump
//   - Kong plugin server:               mcp-auto-kong (started by Kong automatically)
package main

import (
	"log/slog"
	"os"
	_ "time/tzdata"

	"github.com/Kong/go-pdk/server"

	"github.com/lega4e/mcp-auto/pkg/kong"
)

func main() {
	if err := server.StartServer(kong.New, kong.Version, kong.Priority); err != nil {
		slog.Error("kong plugin server stopped", "error", err)
		os.Exit(1)
	}
}
