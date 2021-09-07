package main

import (
	"context"
	"database/sql"
	"github.com/ozonva/ova-travel-api/internal/repo"
	"github.com/ozonva/ova-travel-api/internal/tracer"
	"net"
	"os"
	"time"

	_ "github.com/jackc/pgx/stdlib"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"

	"github.com/ozonva/ova-travel-api/internal/server"
	api "github.com/ozonva/ova-travel-api/pkg/ova-travel-api"
)

func main() {
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
	log := zerolog.New(output).With().Timestamp().Logger()

	_, closer := tracer.InitGlobalTracer()
	defer func() {
		err := closer.Close()
		log.Fatal().Msgf("Can't init tracer %s", err)
	}()

	dsn := "user=kmolchan password=demo dbname=travel_demo sslmode=disable"
	db, err := sql.Open("pgx", dsn)
	defer func() {
		err = db.Close()
	}()

	if err != nil {
		log.Fatal().Msgf("cannot open database connection: %s", err)
		return
	}

	const port = ":8087"
	log.Info().Msgf("start serve port %s", port)

	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal().Msgf("failed to listen: %v", err)
	}

	err = db.PingContext(context.Background())
	if err != nil {
		log.Fatal().Msgf("failed to connect to db: %v", err)
		return
	}

	s := grpc.NewServer()
	r := repo.NewRepo(db)
	api.RegisterTravelRpcServer(s, server.NewTravelServer(&log, r))
	if err := s.Serve(listen); err != nil {
		log.Fatal().Msgf("failed to serve: %v", err)
	}
}
