package server

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/inconshreveable/log15"
	"github.com/pkg/errors"

	"github.com/sourcegraph/sourcegraph/internal/conf/reposource"
	"github.com/sourcegraph/sourcegraph/internal/extsvc/jvmpackages/coursier"
	"github.com/sourcegraph/sourcegraph/internal/vcs"
	"github.com/sourcegraph/sourcegraph/schema"
)

type JvmPackagesArtifactSyncer struct {
	Config *schema.JvmPackagesConnection
}

var _ VCSSyncer = &JvmPackagesArtifactSyncer{}

func (s JvmPackagesArtifactSyncer) Type() string {
	return "jvm_packages"
}

// IsCloneable checks to see if the VCS remote URL is cloneable. Any non-nil
// error indicates there is a problem.
func (s JvmPackagesArtifactSyncer) IsCloneable(ctx context.Context, remoteURL *vcs.URL) error {
	dependencies, err := s.TrackedDependencies(remoteURL.Path)
	if err != nil {
		return err
	}

	for _, dependency := range dependencies {
		log15.Info("Maven.IsCloneable", "dependency", dependency, "url", remoteURL.Path)
		sources, err := coursier.FetchSources(ctx, s.Config, dependency)
		if err != nil {
			return err
		}
		if len(sources) == 0 {
			return errors.Errorf("no sources.jar for dependency %s", dependency)
		}
	}
	return nil
}

// CloneCommand returns the command to be executed for cloning from remote.
func (s JvmPackagesArtifactSyncer) CloneCommand(ctx context.Context, remoteURL *vcs.URL, bareGitDirectory string) (*exec.Cmd, error) {
	dependencies, err := s.TrackedDependencies(remoteURL.Path)
	if err != nil {
		return nil, err
	}

	err = os.MkdirAll(bareGitDirectory, 0755)
	if err != nil {
		return nil, err
	}

	initBareCmd := exec.CommandContext(ctx, "git", "--bare", "init")
	initBareCmd.Dir = bareGitDirectory
	if output, err := runWith(ctx, initBareCmd, false, nil); err != nil {
		return nil, errors.Wrapf(err, "command %s failed with output %q", initBareCmd.Args, string(output))
	}

	for _, dependency := range dependencies {
		if err := s.gitPushDependencyTag(ctx, bareGitDirectory, dependency); err != nil {
			return nil, err
		}
	}

	return exec.CommandContext(ctx, "git", "--version"), nil
}

// Fetch does nothing for Maven packages because they are immutable and cannot be updated after publishing.
func (s JvmPackagesArtifactSyncer) Fetch(ctx context.Context, remoteURL *vcs.URL, dir GitDir) error {
	return nil
}

// RemoteShowCommand returns the command to be executed for showing remote.
func (s JvmPackagesArtifactSyncer) RemoteShowCommand(ctx context.Context, remoteURL *vcs.URL) (cmd *exec.Cmd, err error) {
	return exec.CommandContext(ctx, "git", "remote", "show", "./"), nil
}

func (s JvmPackagesArtifactSyncer) commitJar(ctx context.Context, dependency reposource.Dependency, workingDirectory, path string) error {
	cmd := exec.CommandContext(ctx, "unzip", path, "-d", "./")
	cmd.Dir = workingDirectory
	if output, err := runWith(ctx, cmd, false, nil); err != nil {
		return errors.Wrapf(err, "failed to unzip with output %q", string(output))
	}

	file, err := os.Create(filepath.Join(workingDirectory, "lsif-java.json"))
	if err != nil {
		return err
	}
	defer file.Close()

	jsonContents, err := json.Marshal(&lsifJavaJson{
		Kind:         "maven",
		Jvm:          "8",
		Dependencies: []string{dependency.CoursierSyntax()},
	})
	if err != nil {
		return err
	}

	_, err = file.Write(jsonContents)
	if err != nil {
		return err
	}

	cmd = exec.CommandContext(ctx, "git", "add", ".")
	cmd.Dir = workingDirectory
	if output, err := runWith(ctx, cmd, false, nil); err != nil {
		return errors.Wrapf(err, "failed to git add with output %q", string(output))
	}

	cmd = exec.CommandContext(ctx, "git", "commit", "-m", dependency.CoursierSyntax())
	cmd.Dir = workingDirectory
	if output, err := runWith(ctx, cmd, false, nil); err != nil {
		return errors.Wrapf(err, "failed to git commit with output %q", string(output))
	}

	return nil
}

func (s JvmPackagesArtifactSyncer) TrackedDependencies(path string) (dependencies []reposource.Dependency, err error) {
	module, err := reposource.ParseMavenModule(path)
	if err != nil {
		return nil, err
	}
	for _, dependency := range s.Config.Maven.Artifacts {
		if module.MatchesDependencyString(dependency) {
			dependencies = append(dependencies, reposource.ParseMavenDependency(module, dependency))
		}
	}
	if len(dependencies) == 0 {
		return nil, errors.Errorf("no tracked dependencies for URL path %s", path)
	}
	return dependencies, nil
}

func (s JvmPackagesArtifactSyncer) gitPushDependencyTag(ctx context.Context, bareGitDirectory string, dependency reposource.Dependency) error {
	workingDirectory, err := ioutil.TempDir("", "maven")
	if err != nil {
		return err
	}

	gitTag := "v" + string(reposource.DefaultRepositoryPathPattern)

	paths, err := coursier.FetchSources(ctx, s.Config, dependency)
	if err != nil {
		return err
	}

	if len(paths) == 0 {
		return errors.Errorf("no sources.jar for dependency %s", dependency)
	}

	path := paths[0]

	initCmd := exec.CommandContext(ctx, "git", "init", "--initial-branch", gitTag)
	initCmd.Dir = workingDirectory
	log15.Info("CloneCommand", "tmpPath", workingDirectory, "cwd", initCmd.Dir)
	if output, err := runWith(ctx, initCmd, false, nil); err != nil {
		return errors.Wrapf(err, "command %s failed with output %q", initCmd.Args, string(output))
	}

	err = s.commitJar(ctx, dependency, workingDirectory, path)
	if err != nil {
		return err
	}

	remoteAddCmd := exec.CommandContext(ctx, "git", "remote", "add", "origin", bareGitDirectory)
	remoteAddCmd.Dir = workingDirectory
	if output, err := runWith(ctx, remoteAddCmd, false, nil); err != nil {
		return errors.Wrapf(err, "command %s failed with output %q", remoteAddCmd.Args, string(output))
	}

	gitPushCmd := exec.CommandContext(ctx, "git", "push", "origin", gitTag)
	gitPushCmd.Dir = workingDirectory
	if output, err := runWith(ctx, gitPushCmd, false, nil); err != nil {
		return errors.Wrapf(err, "command %s failed with output %q", gitPushCmd.Args, string(output))
	}

	return nil
}

type lsifJavaJson struct {
	Kind         string   `json:"kind"`
	Jvm          string   `json:"jvm"`
	Dependencies []string `json:"dependencies"`
}
