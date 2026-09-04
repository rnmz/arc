package config

import (
	"fmt"

	"github.com/valkey-io/valkey-go"
)

func InitValkey(cfg Valkey) (valkey.Client, error) {
	addr := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	client, err := valkey.NewClient(valkey.ClientOption{InitAddress: []string{addr}})
	return client, err
}
