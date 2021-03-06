package grpc

import (
	"github.com/sergionunezgo/go-reuse/v2/pkg/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/keepalive"
)

func NewGRPCService() (*grpc.Server, error) {
	gServer := grpc.NewServer(
		grpc.KeepaliveEnforcementPolicy(keepalive.EnforcementPolicy{PermitWithoutStream: true}),
		grpc.KeepaliveParams(keepalive.ServerParameters{}),
	)

	health := NewHealth()
	grpc_health_v1.RegisterHealthServer(gServer, health)
	logger.Log.Info("registered health service")

	return gServer, nil
}
