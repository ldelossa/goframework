package etcd

import (
	"context"
	"os"
	"strings"
	"testing"

	etcd "go.etcd.io/etcd/clientv3"
)

const (
	defaultAddress = "localhost:2379"
	envVar         = "GOFRAMEWORK_ETCD_ADDRS"
)

// Setup returns an etcd.Client configured for our local dev environment. used for testing purposes
// providing nil for addrs uses default "localhost:2379"
// also looks for GOFRAMEWORK_ETCD_ADDRS env variable containing a comma separated list of '[host]:[port]' elements
func Setup(t *testing.T, addrs []string) (*etcd.Client, func()) {
	endpoints := resolve(addrs)

	client, err := etcd.New(
		etcd.Config{
			Endpoints: endpoints,
		},
	)
	if err != nil {
		t.Fatalf("failed to contact etcd: %v", err)
	}

	tearDown := func() {
		client.Delete(context.Background(), "", etcd.WithFromKey())
		client.Close()
	}

	return client, tearDown
}

func resolve(addrs []string) []string {
	env := os.Getenv(envVar)
	switch {
	case env != "":
		s := strings.Split(env, ",")
		return s
	case addrs == nil:
		return []string{defaultAddress}
	default:
		return addrs
	}
}
