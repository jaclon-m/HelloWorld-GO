package healthz

import "httpserver-demo/handler"

func init() {
	health := Health{Indicators: map[string]HealthIndicator{"Default": &ServerHealth{}}}
	handler.DefaultDispatchHandler.AddHandler(&health)
}
