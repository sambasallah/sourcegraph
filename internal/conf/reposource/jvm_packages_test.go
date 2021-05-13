package reposource

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func assertEqual(t *testing.T, got, want interface{}) {
	t.Helper()

	if diff := cmp.Diff(got, want); diff != "" {
		t.Fatalf("(-want +got):\n%s", diff)
	}
}

func TestDecomposeMavenPath(t *testing.T) {
	obtained, _ := ParseMavenModule("//maven/junit/junit")
	assertEqual(t, obtained.RepoName(), "maven/junit/junit")
}
