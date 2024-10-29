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
	backend.CryptoRegistry.Register(backend.Vault, name, &loader{})
}

type loader struct{}

func (l loader) New(ctx context.Context) (backend.Crypto, error) {
	debug.Log("Using Crypto Backend: %s", name)

	return New(ctx)
}

func (l loader) Handles(ctx context.Context, s backend.Storage) error {
	if s.Exists(ctx, ".vault-id") {
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
