package vault

import (
	"context"
	"fmt"

	"github.com/gopasspw/gopass/internal/backend"
	"github.com/gopasspw/gopass/pkg/debug"
)

const (
	name = "vault"
)

func init() {
	backend.StorageRegistry.Register(backend.Vault, name, &loader{})
}

type loader struct{}

func (l loader) New(ctx context.Context, path string) (backend.Storage, error) {
	debug.Log("Using Storage Backend: %s", name)

	return New(ctx, path)
}

func (l loader) Handles(ctx context.Context, path string) error {
	if path == ".vault" {
		return nil
	}

	return fmt.Errorf("not supported")
}

func (l loader) Priority() int {
	return 20
}

func (l loader) String() string {
	return name
}
