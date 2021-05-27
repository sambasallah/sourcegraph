package graphqlbackend

import (
	"context"

	"github.com/sourcegraph/sourcegraph/internal/featureflag"
)

func (r *schemaResolver) ViewerFeatureFlags(ctx context.Context) []*EvaluatedFeatureFlagResolver {
	f := featureflag.FromContext(ctx)
	return flagsToResolvers(f)
}

func flagsToResolvers(input map[string]bool) []*EvaluatedFeatureFlagResolver {
	res := make([]*EvaluatedFeatureFlagResolver, 0, len(input))
	for k, v := range input {
		res = append(res, &EvaluatedFeatureFlagResolver{name: k, value: v})
	}
	return res
}

type EvaluatedFeatureFlagResolver struct {
	name  string
	value bool
}

func (e *EvaluatedFeatureFlagResolver) Name() string {
	return e.name
}

func (e *EvaluatedFeatureFlagResolver) Value() bool {
	return e.value
}
