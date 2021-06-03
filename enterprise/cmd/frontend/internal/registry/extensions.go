package registry

import (
	"context"

	"github.com/sourcegraph/sourcegraph/cmd/frontend/envvar"
	"github.com/sourcegraph/sourcegraph/cmd/frontend/graphqlbackend"
	registry "github.com/sourcegraph/sourcegraph/cmd/frontend/registry/api"
	"github.com/sourcegraph/sourcegraph/internal/conf"
	"github.com/sourcegraph/sourcegraph/internal/database/dbutil"
)

func init() {
	conf.DefaultRemoteRegistry = "https://sourcegraph.com/.api/registry"
	registry.GetLocalExtensionByExtensionID = func(ctx context.Context, db dbutil.DB, extensionIDWithoutPrefix string) (graphqlbackend.RegistryExtension, error) {
		x, err := dbExtensions{}.GetByExtensionID(ctx, extensionIDWithoutPrefix)
		if err != nil {
			return nil, err
		}
		if err := prefixLocalExtensionID(x); err != nil {
			return nil, err
		}
		return &extensionDBResolver{db: db, v: x}, nil
	}

	if envvar.SourcegraphDotComMode() {
		registry.GetLocalFeaturedExtensions = func(ctx context.Context, db dbutil.DB) ([]graphqlbackend.RegistryExtension, error) {
			dbExtensions, err := dbExtensions{}.GetFeaturedExtensions(ctx)
			if err != nil {
				return nil, err
			}
			registryExtensions := make([]graphqlbackend.RegistryExtension, len(dbExtensions))
			for i, x := range dbExtensions {
				registryExtensions[i] = &extensionDBResolver{db: db, v: x}
			}
			return registryExtensions, nil
		}
	}
}

// prefixLocalExtensionID adds the local registry's extension ID prefix (from
// GetLocalRegistryExtensionIDPrefix) to all extensions' extension IDs in the list.
func prefixLocalExtensionID(xs ...*dbExtension) error {
	prefix := registry.GetLocalRegistryExtensionIDPrefix()
	if prefix == nil {
		return nil
	}
	for _, x := range xs {
		x.NonCanonicalExtensionID = *prefix + "/" + x.NonCanonicalExtensionID
		x.NonCanonicalRegistry = *prefix
	}
	return nil
}
