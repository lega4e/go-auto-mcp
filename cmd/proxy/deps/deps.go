// Package deps registers all built-in components for the mcp-auto proxy binary.
// It imports every sub-package that self-registers via init(), so that the proxy
// supports all cache backends, upstream types, auth strategies, rate-limit stores,
// embedding providers, and session stores out of the box.
//
// SDK users who embed only specific components should import individual sub-packages
// instead of this package.
package deps

import (
	// Cache backends.
	_ "github.com/lega4e/mcp-auto/pkg/cache/memory"
	_ "github.com/lega4e/mcp-auto/pkg/cache/redis"

	// Upstream builders.
	_ "github.com/lega4e/mcp-auto/pkg/upstream/command"
	_ "github.com/lega4e/mcp-auto/pkg/upstream/http"
	_ "github.com/lega4e/mcp-auto/pkg/upstream/http/withui"
	_ "github.com/lega4e/mcp-auto/pkg/upstream/script"

	// Inbound auth strategies.
	_ "github.com/lega4e/mcp-auto/pkg/auth/inbound/apikey"
	_ "github.com/lega4e/mcp-auto/pkg/auth/inbound/introspection"
	_ "github.com/lega4e/mcp-auto/pkg/auth/inbound/jwt"

	// Outbound auth strategies.
	_ "github.com/lega4e/mcp-auto/pkg/auth/outbound/apikey"
	_ "github.com/lega4e/mcp-auto/pkg/auth/outbound/bearer"
	_ "github.com/lega4e/mcp-auto/pkg/auth/outbound/none"
	_ "github.com/lega4e/mcp-auto/pkg/auth/outbound/oauth2"
	_ "github.com/lega4e/mcp-auto/pkg/auth/outbound/oauth2usersession"

	// Scripting runtimes (register both inbound and outbound JS/Lua strategies).
	_ "github.com/lega4e/mcp-auto/pkg/runtime/js"
	_ "github.com/lega4e/mcp-auto/pkg/runtime/lua"

	// Rate-limit stores.
	_ "github.com/lega4e/mcp-auto/pkg/ratelimit/memory"
	_ "github.com/lega4e/mcp-auto/pkg/ratelimit/redis"

	// Embedding providers.
	_ "github.com/lega4e/mcp-auto/pkg/embedding/hugot"

	// Session store backends.
	_ "github.com/lega4e/mcp-auto/pkg/session/memory"
	_ "github.com/lega4e/mcp-auto/pkg/session/postgres"
	_ "github.com/lega4e/mcp-auto/pkg/session/redis"
)
