// Package all imports all built-in inbound auth strategy sub-packages,
// registering them via their init() functions.
// Import this package with a blank import to make all strategies available:
//
//	import _ "github.com/lega4e/mcp-auto/pkg/auth/inbound/all"
package all

import (
	_ "github.com/lega4e/mcp-auto/pkg/auth/inbound/apikey"
	_ "github.com/lega4e/mcp-auto/pkg/auth/inbound/introspection"
	_ "github.com/lega4e/mcp-auto/pkg/auth/inbound/jwt"
	_ "github.com/lega4e/mcp-auto/pkg/scripting/js"
	_ "github.com/lega4e/mcp-auto/pkg/scripting/lua"
)
