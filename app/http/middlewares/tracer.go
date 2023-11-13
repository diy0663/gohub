package middlewares

import (
	"context"

	"github.com/diy0663/gohub/pkg/tracer"
	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
)

func Tracing() gin.HandlerFunc {
	return func(c *gin.Context) {

		var newCtx context.Context
		var span opentracing.Span
		spanCtx, err := tracer.Tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(c.Request.Header))
		if err != nil {

			span, newCtx = opentracing.StartSpanFromContextWithTracer(c.Request.Context(), tracer.Tracer, c.Request.URL.Path)
		} else {

			span, newCtx = opentracing.StartSpanFromContextWithTracer(
				c.Request.Context(),
				tracer.Tracer,
				c.Request.URL.Path,
				opentracing.ChildOf(spanCtx),
				opentracing.Tags{
					"Key":   string(ext.Component),
					"Value": "HTTP",
				},
			)
		}
		defer span.Finish()
		c.Request = c.Request.WithContext(newCtx)
		c.Next()
	}
}
