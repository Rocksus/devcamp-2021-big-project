package tracer

import (
	"context"
	"github.com/opentracing/opentracing-go"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	"github.com/uber/jaeger-lib/metrics"
	"io"
	"log"
)

func Init(appName string) (io.Closer, error) {
	cfg := jaegercfg.Configuration{
		Sampler: &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
		Reporter: &jaegercfg.ReporterConfig{
			LogSpans:           true,
			LocalAgentHostPort: "jaeger:6831",
		},
	}

	jLogger := jaeger.NullLogger
	jMetricsFactory := metrics.NullFactory

	closer, err := cfg.InitGlobalTracer(
		appName,
		jaegercfg.Logger(jLogger),
		jaegercfg.Metrics(jMetricsFactory),
	)
	if err != nil {
		log.Println("[Tracer] could not initialize jaeger tracer, err:", err.Error())
		return nil, err
	}

	return closer, nil
}

func StartSpanFromContext(ctx context.Context, operationName string) (opentracing.Span, context.Context) {
	return opentracing.StartSpanFromContext(ctx, operationName)
}
