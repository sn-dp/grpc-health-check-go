package service

import (
	"context"
	"fmt"
	"github.com/sn-dp/grpc-health-check-go/proto/google.golang.org/grpc/health/grpc_health_v1"
)

type HealthCheckService struct {
	isHealthy bool
}

func (hs *HealthCheckService) Check(context context.Context, request *grpc_health_v1.HealthCheckRequest) (*grpc_health_v1.HealthCheckResponse, error) {
	if hs.isHealthy {
		return &grpc_health_v1.HealthCheckResponse{
			Status: grpc_health_v1.HealthCheckResponse_SERVING,
		}, nil
	} else {
		fmt.Println("Health Check failed")
		return &grpc_health_v1.HealthCheckResponse{
			Status: grpc_health_v1.HealthCheckResponse_NOT_SERVING,
		}, nil
	}
}

func (hs *HealthCheckService) StopServing() {
	fmt.Println("health check status changed to NOT_SERVING")
	hs.isHealthy = false
}

func SetupHealthCheck() *HealthCheckService {
	return &HealthCheckService{isHealthy:true}
}
