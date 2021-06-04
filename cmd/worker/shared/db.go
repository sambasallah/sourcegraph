package shared

import (
	"database/sql"
	"fmt"

	"github.com/sourcegraph/sourcegraph/internal/conf/conftypes"
	"github.com/sourcegraph/sourcegraph/internal/database/dbconn"
)

// TODO - document
func InitDatabase() (*sql.DB, error) {
	conn, err := initDatabaseMemo.Init()
	return conn.(*sql.DB), err
}

var initDatabaseMemo = NewMemo(func() (interface{}, error) {
	postgresDSN := WatchServiceConnectionValue(func(serviceConnections conftypes.ServiceConnections) string {
		return serviceConnections.PostgresDSN
	})

	if err := dbconn.SetupGlobalConnection(postgresDSN); err != nil {
		return nil, fmt.Errorf("failed to connect to frontend database: %s", err)
	}

	return dbconn.Global, nil
})
