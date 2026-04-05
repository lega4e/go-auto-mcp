// Package all imports all upstream builder sub-packages for side effects,
// registering all built-in upstream types (http, command, script).
package all

import (
	_ "github.com/lega4e/mcp-auto/internal/upstream/commandbuilder"
	_ "github.com/lega4e/mcp-auto/internal/upstream/http"
	_ "github.com/lega4e/mcp-auto/internal/upstream/scriptbuilder"
)
