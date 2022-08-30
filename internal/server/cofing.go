package server

import (
	"github.com/yudgxe/simple-server/internal/storage"
	"github.com/yudgxe/simple-server/internal/subscriber"
)

type Config struct {
	BindAddr   string `toml:"bind_addr"`
	Storage    *storage.Config
	Subscriber *subscriber.Config
}

func NewConfig() *Config {
	return &Config{
		BindAddr:   ":8080",
		Storage:    storage.NewConfig(),
		Subscriber: subscriber.NewConfig(),
	}
}
