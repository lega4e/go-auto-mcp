package upstream

import (
	nethttp "net/http"

	pkgoutbound "github.com/lega4e/mcp-auto/pkg/auth/outbound"
	"github.com/lega4e/mcp-auto/pkg/config"
	pkghttp "github.com/lega4e/mcp-auto/pkg/upstream/http"
)

// NewHTTPClient builds an *http.Client for an upstream.
// See pkg/upstream/http.NewHTTPClient.
func NewHTTPClient(cfg *config.UpstreamConfig, provider pkgoutbound.TokenProvider) (*nethttp.Client, error) {
	return pkghttp.NewHTTPClient(cfg, provider)
}
