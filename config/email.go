package config

import "github.com/diy0663/go_project_packages/config"

func init() {
	config.Add("email", config.StrMap{
		"host":     config.Env("EMAIL_HOST", ""),
		"port":     config.Env("EMAIL_PORT", ""),
		"username": config.Env("EMAIL_USERNAME", ""),
		"password": config.Env("EMAIL_PASSWORD", ""),
		"from":     config.Env("EMAIL_FROM", ""),
		"to":       config.Env("EMAIL_TO", ""),
	})
}
