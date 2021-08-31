package server

import (
	"context"

	"github.com/rs/zerolog"
	"google.golang.org/protobuf/types/known/emptypb"

	api "github.com/ozonva/ova-travel-api/pkg/ova-travel-api"
)

type GRPCServer struct {
	api.UnimplementedTravelRpcServer
	logger *zerolog.Logger
}

func NewTravelServer(logger *zerolog.Logger) api.TravelRpcServer {
	return &GRPCServer{
		UnimplementedTravelRpcServer: api.UnimplementedTravelRpcServer{},
		logger:                       logger,
	}
}

func (g *GRPCServer) CreateTravel(ctx context.Context, request *api.CreateTravelRequest) (*api.CreateTravelResponce, error) {
	g.logger.Info().Msg("Request: CreateTravel")
	return g.UnimplementedTravelRpcServer.CreateTravel(ctx, request)

}

func (g *GRPCServer) DescribeTravel(ctx context.Context, request *api.DescribeTravelRequest) (*api.DescribeTravelResponse, error) {
	g.logger.Info().Msg("Request: DescribeTravel")
	return g.UnimplementedTravelRpcServer.DescribeTravel(ctx, request)
}

func (g *GRPCServer) ListTravel(ctx context.Context, req *emptypb.Empty) (*api.ListTravelsResponse, error) {
	g.logger.Info().Msg("Request: ListTravel")
	return g.UnimplementedTravelRpcServer.ListTravels(ctx, req)
}

func (g *GRPCServer) RemoveTravel(ctx context.Context, req *api.RemoveTravelRequest) (*emptypb.Empty, error) {
	g.logger.Info().Msg("Request: RemoveTravel")
	return g.UnimplementedTravelRpcServer.RemoveTravel(ctx, req)
}
