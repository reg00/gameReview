package config

import (
	"time"
)

type Configuration struct {
	Http    HTTP
	Storage Storage
	Igdb    IGDB
	Cache   Cache
}

type HTTP struct {
	Port    string
	Timeout int
}

func (h *HTTP) GetTimeout() time.Duration {
	return time.Duration(h.Timeout) * time.Second
}

type Storage struct {
	Provider string
	Options  StorageOptions
}

type StorageOptions struct {
	Driver   string
	Host     string
	User     string
	Password string
	Dbname   string
	Port     string
}

type IGDB struct {
	ClientId     string
	ClientSecret string
}

type Cache struct {
	Addr string
}
