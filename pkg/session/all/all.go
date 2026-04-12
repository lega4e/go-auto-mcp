// Package all imports all built-in session store sub-packages,
// registering them via their init() functions.
// Import this package with a blank import to make all session store providers available:
//
//	import _ "github.com/lega4e/mcp-auto/pkg/session/all"
package all

import (
	_ "github.com/lega4e/mcp-auto/pkg/session/memory"
	_ "github.com/lega4e/mcp-auto/pkg/session/postgres"
	_ "github.com/lega4e/mcp-auto/pkg/session/redis"
)
