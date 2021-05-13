package reposource

import (
	"fmt"
	"net/url"
	"strings"

	"github.com/sourcegraph/sourcegraph/internal/api"
)

type Module struct {
	GroupId    string
	ArtifactId string
}

func (m Module) MatchesDependencyString(dependency string) bool {
	return strings.HasPrefix(dependency, fmt.Sprintf("%s:%s:", m.GroupId, m.ArtifactId))
}

func (m Module) RepoName() api.RepoName {
	return api.RepoName(fmt.Sprintf("maven/%s/%s", m.GroupId, m.ArtifactId))
}

func (m Module) CloneURL() string {
	cloneURL := url.URL{Path: string(m.RepoName())}
	return cloneURL.String()
}

type Dependency struct {
	Module
	Version string
}

func (d Dependency) CoursierSyntax() string {
	return fmt.Sprintf("%s:%s:%s", d.Module.GroupId, d.Module.ArtifactId, d.Version)
}

func ParseMavenDependencyString(dependency string) Dependency {
	parts := strings.Split(dependency, ":")
	return Dependency{
		Module: Module{
			GroupId:    parts[0],
			ArtifactId: parts[1],
		},
		Version: parts[2],
	}
}
func ParseMavenDependency(module Module, dependency string) Dependency {
	colonIndex := strings.LastIndex(dependency, ":") + 1
	return Dependency{
		Module:  module,
		Version: dependency[colonIndex:],
	}
}

func ParseMavenModule(path string) (Module, error) {
	parts := strings.SplitN(strings.TrimPrefix(path, "maven/"), "/", 2)
	if len(parts) != 2 {
		return Module{}, fmt.Errorf("failed to parse a maven module from the path %s", path)
	}

	return Module{
		GroupId:    parts[0],
		ArtifactId: parts[1],
	}, nil
}
