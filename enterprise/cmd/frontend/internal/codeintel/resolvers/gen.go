package resolvers

//go:generate env GOBIN=$PWD/.bin GO111MODULE=on go install github.com/derision-test/go-mockgen/cmd/go-mockgen
//go:generate $PWD/.bin/go-mockgen -f github.com/sourcegraph/sourcegraph/enterprise/cmd/frontend/internal/codeintel/resolvers -i GitserverClient -i DBStore -i LSIFStore -i IndexEnqueuer -i RepoUpdaterClient -i EnqueuerDBStore -i EnqueuerGitserverClient -o mock_iface_test.go
//go:generate $PWD/.bin/go-mockgen -f github.com/sourcegraph/sourcegraph/enterprise/cmd/frontend/internal/codeintel/resolvers -i PositionAdjuster -o mock_position_adjuster_test.go
