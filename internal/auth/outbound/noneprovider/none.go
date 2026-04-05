// Package noneprovider registers the "none" outbound auth strategy.
package noneprovider

import (
	"context"

	"github.com/lega4e/mcp-auto/internal/auth/outbound"
	"github.com/lega4e/mcp-auto/internal/config"
	"github.com/lega4e/mcp-auto/internal/runtime"
)

func init() {
	outbound.RegisterProvider("none", func(_ context.Context, _ *config.OutboundAuthConfig, _ *runtime.Registry) (outbound.TokenProvider, error) {
		return &NoneProvider{}, nil
	})
}

// NoneProvider is a no-op provider that adds no authentication headers.
type NoneProvider struct{}

// Token returns an empty token; no authentication is injected.
func (p *NoneProvider) Token(_ context.Context) (string, error) { return "", nil }

// RawHeaders returns nil; no authentication headers are injected.
func (p *NoneProvider) RawHeaders(_ context.Context) (map[string]string, error) { return nil, nil }
