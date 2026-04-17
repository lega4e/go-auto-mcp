// Package memory registers the in-memory rate limit store via init().
// Import this package with a blank identifier to enable the memory store:
//
//	import _ "github.com/lega4e/mcp-auto/pkg/ratelimit/memory"
package memory

import (
	"context"

	"github.com/ulule/limiter/v3"
	memstore "github.com/ulule/limiter/v3/drivers/store/memory"

	"github.com/lega4e/mcp-auto/pkg/config"
	pkgratelimit "github.com/lega4e/mcp-auto/pkg/ratelimit"
)

func init() {
	pkgratelimit.Register("memory", func(_ context.Context, _ *config.ProxySpec) (limiter.Store, error) {
		return memstore.NewStore(), nil
	})
}
