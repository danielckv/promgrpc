package promgrpc_test

import (
	"context"
	"testing"
	"time"

	"github.com/piotrkowalczuk/promgrpc/v4"
	"github.com/piotrkowalczuk/promgrpc/v4/pb/private/test"
	"github.com/prometheus/client_golang/prometheus"
	"google.golang.org/grpc"
)

func suite(t *testing.T) (test.TestServiceClient, *prometheus.Registry, func(*testing.T)) {
	lis := listener(t)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	ssh := promgrpc.ServerStatsHandler()
	csh := promgrpc.ClientStatsHandler()
	srv := grpc.NewServer(grpc.StatsHandler(ssh))

	test.RegisterTestServiceServer(srv, newDemoServer())

	reg := prometheus.NewRegistry()
	registerCollector(t, reg, ssh)
	registerCollector(t, reg, csh)

	go func() {
		if err := srv.Serve(lis); err != grpc.ErrServerStopped {
			if err != nil {
				t.Error(err)
			}
		}
	}()

	con, err := grpc.DialContext(ctx, lis.Addr().String(), grpc.WithInsecure(), grpc.WithBlock(), grpc.WithStatsHandler(csh))
	if err != nil {
		t.Fatal(err)
	}

	return test.NewTestServiceClient(con), reg, func(t *testing.T) {
		srv.GracefulStop()
	}
}
