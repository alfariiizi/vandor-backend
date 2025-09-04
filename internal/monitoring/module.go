package monitoring

import (
	"github.com/prometheus/client_golang/prometheus"
	"go.uber.org/fx"
)

var Module = fx.Module(
	"monitoring",
	fx.Provide(
		func() *prometheus.Registry {
			reg := prometheus.NewRegistry()
			reg.MustRegister(prometheus.NewGoCollector())
			reg.MustRegister(prometheus.NewProcessCollector(prometheus.ProcessCollectorOpts{}))
			return reg
		},
	),
)
