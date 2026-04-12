// Package all imports all built-in cache store sub-packages,
// registering them via their init() functions.
// Import this package with a blank import to make all providers available:
//
//	import _ "github.com/lega4e/mcp-auto/pkg/cache/all"
package all

import (
	_ "github.com/lega4e/mcp-auto/pkg/cache/memory"
	_ "github.com/lega4e/mcp-auto/pkg/cache/redis"
)
