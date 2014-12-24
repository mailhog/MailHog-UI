package config

import (
	"flag"

	"github.com/ian-kent/envconf"
)

func DefaultConfig() *Config {
	return &Config{
		APIHost:      "",
		HTTPBindAddr: "0.0.0.0:8025",
	}
}

type Config struct {
	APIHost      string
	HTTPBindAddr string
}

var cfg = DefaultConfig()

func Configure() *Config {
	return cfg
}

func RegisterFlags() {
	flag.StringVar(&cfg.APIHost, "api-host", envconf.FromEnvP("MH_API_HOST", "").(string), "API URL for MailHog UI to connect to, e.g. http://some.host:1234")
	flag.StringVar(&cfg.HTTPBindAddr, "httpbindaddr", envconf.FromEnvP("MH_HTTP_BIND_ADDR", "0.0.0.0:8025").(string), "HTTP bind interface and port, e.g. 0.0.0.0:8025 or just :8025")
}
