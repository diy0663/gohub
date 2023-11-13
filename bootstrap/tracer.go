package bootstrap

import (
	"github.com/diy0663/go_project_packages/config"
	"github.com/diy0663/gohub/pkg/tracer"
)

func SetupTracer() error {

	jaegerTracer, _, err := tracer.NewJaegerTracer(config.GetString("app.name"), config.GetString("tracer.host")+":"+config.GetString("tracer.port"))
	if err != nil {
		return err
	}
	tracer.Tracer = jaegerTracer

	return nil
}
