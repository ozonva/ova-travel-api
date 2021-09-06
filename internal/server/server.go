package server

import (
	"context"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
	"github.com/ozonva/ova-travel-api/internal/metrics"

	"github.com/rs/zerolog"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/ozonva/ova-travel-api/internal/message.producer"
	"github.com/ozonva/ova-travel-api/internal/repo"
	"github.com/ozonva/ova-travel-api/internal/travel"
	"github.com/ozonva/ova-travel-api/internal/utils"
	api "github.com/ozonva/ova-travel-api/pkg/ova-travel-api"
)

type GRPCServer struct {
	api.UnimplementedTravelRpcServer
	logger          *zerolog.Logger
	repo            repo.Repo
	messageProducer message_producer.TravelMsgProducer
	metrics         *metrics.Metrics
}

func NewTravelServer(logger *zerolog.Logger, rep repo.Repo) api.TravelRpcServer {
	return &GRPCServer{
		UnimplementedTravelRpcServer: api.UnimplementedTravelRpcServer{},
		logger:                       logger,
		repo:                         rep,
		messageProducer:              message_producer.NewTravelMsgProducer(),
		metrics:                      metrics.NewMetrics(),
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

	g.messageProducer.TravelSaved()
	g.metrics.CreateTravelCounterInc()

	return resp, nil
}

func SplitTripsByBatch(arr []*api.Travel, batch int) [][]*api.Travel {
	batchSlice := make([][]*api.Travel, 0)
	for i := 0; i < len(arr); i += batch {
		batchSlice = append(batchSlice, arr[i:utils.MinInt(i+batch, len(arr))])
	}

	return batchSlice
}

func (g *GRPCServer) MultipleCreateTravel(ctx context.Context, request *api.MultipleCreateTravelRequest) (*emptypb.Empty, error) {
	g.logger.Info().Msg("Request: MultipleCreateTravel")

	span, ctx := opentracing.StartSpanFromContext(ctx, "operation_name")
	defer span.Finish()
	span.LogFields(log.String("MultipleCreateTravel", "start"))

	trips := request.GetItems()
	span.LogFields(log.Int("MultipleCreateTravel_input_len", len(trips)))

	const batchSize = 4
	batchedTrips := SplitTripsByBatch(trips, batchSize)
	for _, batch := range batchedTrips {
		sp := opentracing.StartSpan(
			"MultipleCreateTravel_batch_process",
			opentracing.ChildOf(span.Context()))
		defer sp.Finish()

		newTrips := make([]travel.Trip, 0)
		for _, trip := range batch {
			newTrip := travel.Trip{
				UserID:       0,
				FromLocation: trip.From,
				DestLocation: trip.Dest,
			}

			newTrips = append(newTrips, newTrip)
		}

		if err := g.repo.AddEntities(newTrips); err != nil {
			return nil, status.Error(codes.Unavailable, "database store failed")
		}

		sp.LogFields(log.String("batch_processing", "processed batch"))
	}

	g.metrics.MultiCreateTravelCounterInc()
	return new(emptypb.Empty), nil
}

func getTrip(travelMsg *api.Travel) travel.Trip {
	return travel.Trip{
		UserID:       travelMsg.Id,
		FromLocation: travelMsg.From,
		DestLocation: travelMsg.Dest,
	}
}

func (g *GRPCServer) UpdateTravel(ctx context.Context, request *api.UpdateTravelRequest) (*emptypb.Empty, error) {
	g.logger.Info().Msg("Request: UpdateTravel")

	newTrip := getTrip(request.GetTravel())
	if err := g.repo.UpdateEntity(newTrip.UserID, &newTrip); err != nil {
		return nil, status.Error(codes.Unavailable, "database update failed")
	}

	g.messageProducer.TravelUpdated()
	g.metrics.UpdateTravelCounterInc()
	return new(emptypb.Empty), nil
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

	g.metrics.DescribeTravelCounterInc()
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

	g.metrics.ListTravelCounterInc()
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

	g.messageProducer.TravelDeleted()
	g.metrics.RemoveTravelCounterInc()
	return new(emptypb.Empty), status.Error(codes.OK, "successfully removed")
}
