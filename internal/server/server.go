package server

import (
	"context"

	"github.com/rs/zerolog"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/ozonva/ova-travel-api/internal/repo"
	"github.com/ozonva/ova-travel-api/internal/travel"
	api "github.com/ozonva/ova-travel-api/pkg/ova-travel-api"
)

type GRPCServer struct {
	api.UnimplementedTravelRpcServer
	logger *zerolog.Logger
	repo   repo.Repo
}

func NewTravelServer(logger *zerolog.Logger, rep repo.Repo) api.TravelRpcServer {
	return &GRPCServer{
		UnimplementedTravelRpcServer: api.UnimplementedTravelRpcServer{},
		logger:                       logger,
		repo:                         rep,
	}
}

func (g *GRPCServer) CreateTravel(ctx context.Context, request *api.CreateTravelRequest) (*api.CreateTravelResponse, error) {
	g.logger.Info().Msg("Request: CreateTravel")

	trip := travel.Trip{
		UserID:       0,
		FromLocation: request.GetFrom(),
		DestLocation: request.GetDest(),
	}

	if err := g.repo.AddEntities([]travel.Trip{trip}); err != nil {
		return nil, status.Error(codes.Unavailable, "database store failed")
	}

	resp := &api.CreateTravelResponse{}

	return resp, nil
}

func (g *GRPCServer) DescribeTravel(ctx context.Context, request *api.DescribeTravelRequest) (*api.DescribeTravelResponse, error) {
	g.logger.Info().Msg("Request: DescribeTravel")

	id := request.Id
	trip, err := g.repo.DescribeEntity(id)

	if err != nil {
		return nil, status.Error(codes.Unavailable, "database fetch error")
	}

	if trip == nil {
		return nil, status.Error(codes.NotFound, "identity not found")
	}

	res := &api.DescribeTravelResponse{
		Travel: &api.Travel{
			Id:   trip.UserID,
			From: trip.FromLocation,
			Dest: trip.DestLocation,
		},
	}
	return res, status.Error(codes.OK, "")
}

func (g *GRPCServer) ListTravel(ctx context.Context, req *api.ListTravelsRequest) (*api.ListTravelsResponse, error) {
	g.logger.Info().Msg("Request: ListTravel")

	limit := req.GetLimit()
	offset := req.GetOffset()

	list, err := g.repo.ListEntities(limit, offset)
	if err != nil {
		return nil, status.Error(codes.Unavailable, "database list error")
	}

	if len(list) == 0 {
		return nil, status.Error(codes.NotFound, "entities not found")
	}

	idList := make([]*api.Travel, 0, len(list))
	for i := 0; i < len(list); i++ {
		idList = append(idList, &api.Travel{
			Id:   list[i].UserID,
			From: list[i].FromLocation,
			Dest: list[i].DestLocation,
		})
	}

	res := &api.ListTravelsResponse{
		Items: idList,
	}

	return res, status.Error(codes.OK, "")
}

func (g *GRPCServer) RemoveTravel(ctx context.Context, req *api.RemoveTravelRequest) (*emptypb.Empty, error) {
	g.logger.Info().Msg("Request: RemoveTravel")

	id := req.GetId()
	err := g.repo.RemoveEntity(id)

	if err != nil {
		g.logger.Info().Msg("error occurred while RemoveTrip")
		return new(emptypb.Empty), status.Error(codes.Unavailable, "database delete error")
	}

	return new(emptypb.Empty), status.Error(codes.OK, "successfully removed")
}
