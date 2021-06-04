package shared

import (
	"log"

	"github.com/sourcegraph/sourcegraph/internal/conf"
	"github.com/sourcegraph/sourcegraph/internal/conf/conftypes"
)

// TODO - document
func WatchServiceConnectionValue(f func(serviceConnections conftypes.ServiceConnections) string) string {
	value := f(conf.Get().ServiceConnections)
	conf.Watch(func() {
		if newValue := f(conf.Get().ServiceConnections); value != newValue {
			log.Fatalf("Detected settings change change, restarting to take effect: %s", newValue)
		}
	})

	return value
}
