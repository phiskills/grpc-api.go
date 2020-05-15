package grpc

import (
	"google.golang.org/grpc/health/grpc_health_v1"
)

type status int

const (
	Unknown status = iota
	Serving
	NotServing
)

func (s status) String() string {
	return map[status]string{
		Unknown:    "UNKNOWN",
		Serving:    "SERVING",
		NotServing: "NOT_SERVING",
	}[s]
}

func (s status) Value() grpc_health_v1.HealthCheckResponse_ServingStatus {
	return map[status]grpc_health_v1.HealthCheckResponse_ServingStatus{
		Unknown:    grpc_health_v1.HealthCheckResponse_UNKNOWN,
		Serving:    grpc_health_v1.HealthCheckResponse_SERVING,
		NotServing: grpc_health_v1.HealthCheckResponse_NOT_SERVING,
	}[s]
}
