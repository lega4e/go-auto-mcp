// Package all imports all inbound auth validator sub-packages for side effects,
// registering all built-in strategies (jwt, introspection, apikey, lua, js_script).
package all

import (
	_ "github.com/lega4e/mcp-auto/internal/auth/inbound/apikeyvalidator"
	_ "github.com/lega4e/mcp-auto/internal/auth/inbound/introspectionvalidator"
	_ "github.com/lega4e/mcp-auto/internal/auth/inbound/jsvalidator"
	_ "github.com/lega4e/mcp-auto/internal/auth/inbound/jwtvalidator"
	_ "github.com/lega4e/mcp-auto/internal/auth/inbound/luavalidator"
)
