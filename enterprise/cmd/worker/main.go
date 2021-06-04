package main

import (
	"context"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/sourcegraph/sourcegraph/cmd/worker/shared"
	eiauthz "github.com/sourcegraph/sourcegraph/enterprise/internal/authz"
	"github.com/sourcegraph/sourcegraph/internal/authz"
	"github.com/sourcegraph/sourcegraph/internal/conf"
	"github.com/sourcegraph/sourcegraph/internal/database"
)

func main() {
	debug, _ := strconv.ParseBool(os.Getenv("DEBUG"))
	if debug {
		log.Println("enterprise edition")
	}

	go setAuthProviders()

	shared.Start(map[string]shared.Task{
		// Empty for now
	})
}

// TODO - rfc
// TODO - move
// TODO - document
func setAuthProviders() {
	_, err := shared.InitDatabase()
	if err != nil {
		return
	}

	ctx := context.Background()

	for range time.NewTicker(5 * time.Second).C {
		allowAccessByDefault, authzProviders, _, _ := eiauthz.ProvidersFromConfig(ctx, conf.Get(), database.GlobalExternalServices)
		authz.SetProviders(allowAccessByDefault, authzProviders)
	}
}
