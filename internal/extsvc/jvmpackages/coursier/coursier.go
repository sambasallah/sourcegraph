package coursier

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"
	"strings"

	"github.com/inconshreveable/log15"
	"github.com/pkg/errors"

	"github.com/sourcegraph/sourcegraph/internal/conf/reposource"
	"github.com/sourcegraph/sourcegraph/schema"
)

func ListArtifactIDs(ctx context.Context, config *schema.JvmPackagesConnection, groupID string) ([]string, error) {
	return runCoursierCommand(ctx, config, "complete", groupID+":")
}

func ListVersions(ctx context.Context, config *schema.JvmPackagesConnection, groupID, artifactID string) ([]string, error) {
	return runCoursierCommand(ctx, config, "complete", groupID+":"+artifactID+":")
}

func FetchSources(ctx context.Context, config *schema.JvmPackagesConnection, dependency reposource.Dependency) ([]string, error) {
	return runCoursierCommand(
		ctx,
		config,
		"fetch", "--intransitive",
		dependency.CoursierSyntax(),
		"--classifier", "sources",
	)
}

func Exists(ctx context.Context, config *schema.JvmPackagesConnection, dependency reposource.Dependency) (bool, error) {
	versions, err := runCoursierCommand(
		ctx,
		config,
		"complete",
		dependency.CoursierSyntax(),
	)
	return len(versions) > 0, err
}

func runCoursierCommand(ctx context.Context, config *schema.JvmPackagesConnection, args ...string) ([]string, error) {
	log15.Info("runCoursierCommand", "args", args)
	cmd := exec.CommandContext(ctx, "coursier", args...)
	if config.Maven.Credentials != "" {
		cmd.Env = append(cmd.Env, fmt.Sprintf("COURSIER_CREDENTIALS=%v", config.Maven.Credentials))
	}
	if len(config.Maven.Repositories) > 0 {
		cmd.Env = append(
			cmd.Env,
			fmt.Sprintf("COURSIER_REPOSITORIES=%v", strings.Join(config.Maven.Repositories, "|")),
		)
	}
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		return nil, errors.Wrap(err, stderr.String())
	}

	return strings.Split(strings.Trim(stdout.String(), " \n"), "\n"), nil
}
