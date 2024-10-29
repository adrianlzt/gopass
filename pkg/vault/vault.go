package vault

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/vault/api"
)

// Client represents a Vault client.
type Client struct {
	client *api.Client
}

// NewClient creates a new Vault client.
func NewClient(addr string, token string) (*Client, error) {
	config := &api.Config{
		Address: addr,
	}
	client, err := api.NewClient(config)
	if err != nil {
		return nil, fmt.Errorf("failed to create Vault client: %w", err)
	}
	client.SetToken(token)
	return &Client{client: client}, nil
}

// ListKVEngines lists the KV engines available in Vault.
func (c *Client) ListKVEngines(ctx context.Context) ([]string, error) {
	mounts, err := c.client.Sys().ListMounts()
	if err != nil {
		return nil, fmt.Errorf("failed to list mounts: %w", err)
	}

	var kvEngines []string
	for path, mount := range mounts {
		if strings.HasPrefix(mount.Type, "kv") {
			kvEngines = append(kvEngines, path)
		}
	}

	return kvEngines, nil
}

// GetSecret retrieves a secret from Vault.
func (c *Client) GetSecret(ctx context.Context, path string) (map[string]interface{}, error) {
	secret, err := c.client.Logical().Read(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read secret: %w", err)
	}
	if secret == nil {
		return nil, fmt.Errorf("secret not found at path: %s", path)
	}
	return secret.Data, nil
}
