package monitoring

import "github.com/prometheus/client_golang/prometheus"

func Init() error {
	if err := prometheus.Register(totalRequests); err != nil {
		return err
	}
	if err := prometheus.Register(responseStatus); err != nil {
		return err
	}
	if err := prometheus.Register(latency); err != nil {
		return err
	}
	return nil
}
