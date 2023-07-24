package common

import (
	"github.com/Kora-Dance/koradance-backend/internal/common/router"
)

// APIPathProvider is an interface for all API path providers
type APIPathProvider interface {
	// RegisterPath registers all API paths
	RegisterPath(router router.KoraRouter)
}
