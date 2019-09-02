package etcd

import (
	"fmt"
	"testing"

	etcd "go.etcd.io/etcd/clientv3"
)

const (
	DefaultAddress = "localhost:12379"
)

// Setup returns an etcd.Client configured for our local dev environment. used for testing purposes
func Setup(t *testing.T) (*etcd.Client, func()) {
	e := fmt.Sprintf(DefaultAddress)
	endpoints := []string{
		e,
	}
	client, err := etcd.New(
		etcd.Config{
			Endpoints: endpoints,
		},
	)
	if err != nil {
		t.Fatalf("failed to contact etcd: %v", err)
	}

	tearDown := func() {
		client.Close()
	}

	return client, tearDown
}
