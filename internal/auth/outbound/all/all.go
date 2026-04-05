// Package all imports all outbound auth provider sub-packages for side effects,
// registering all built-in strategies (bearer, api_key, oauth2, lua, js, none).
package all

import (
	_ "github.com/lega4e/mcp-auto/internal/auth/outbound/apikeyprovider"
	_ "github.com/lega4e/mcp-auto/internal/auth/outbound/bearerprovider"
	_ "github.com/lega4e/mcp-auto/internal/auth/outbound/jsprovider"
	_ "github.com/lega4e/mcp-auto/internal/auth/outbound/luaprovider"
	_ "github.com/lega4e/mcp-auto/internal/auth/outbound/noneprovider"
	_ "github.com/lega4e/mcp-auto/internal/auth/outbound/oauth2provider"
)
