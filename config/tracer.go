package config

import (
	configPkg "github.com/diy0663/go_project_packages/config"
)

func init() {
	configPkg.Add("tracer", map[string]interface{}{
		"host": configPkg.Env("TRACE_HOST", "127.0.0.1"),
		"port": configPkg.Env("TRACE_PORT", "6831"),
	})
}
