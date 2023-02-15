package app

import "github.com/diy0663/gohub/pkg/config"

func IsLocal() bool {
	return config.Get("app.env") == "local"
}
func IsProduction() bool {
	return config.Get("app.env") == "production"
}

func IsTesting() bool {
	return config.Get("app.env") == "testing"
}
