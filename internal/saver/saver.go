package saver

import (
	"errors"
	"github.com/ozonva/ova-travel-api/internal/flusher"
	"github.com/ozonva/ova-travel-api/internal/travel"
	"time"
)

type Saver interface {
	Save(entity travel.Trip)
	Init() error
	Close()
}

func NewSaver(
	capacity uint,
	timeLimitInSec uint64,
	flusher flusher.Flusher) (Saver, error) {
	if capacity == 0 {
		return nil, errors.New("saver capacity can't be 0")
	}
	if timeLimitInSec == 0 {
		return nil, errors.New("saver capacity can't be 0")
	}

	return &saver{
		capacity:       capacity,
		timeLimitInSec: timeLimitInSec,
		flusher:        flusher,
		entities:       make([]travel.Trip, 0, capacity),
	}, nil
}

type saver struct {
	capacity       uint
	timeLimitInSec uint64
	flusher        flusher.Flusher
	entities       []travel.Trip

	ticker *time.Ticker
}

func (s *saver) Save(entity travel.Trip) {
	s.entities = append(s.entities, entity)
	if uint(len(s.entities)) == s.capacity {
		s.flush()
	}
}

func (s *saver) restartTimer() {
	s.ticker = time.NewTicker(time.Duration(s.timeLimitInSec) * time.Second)

	go func() {
		<-s.ticker.C

		s.flush()
		s.restartTimer()
	}()
}

func (s *saver) Init() error {
	s.entities = make([]travel.Trip, 0)
	s.restartTimer()
	return nil
}

func (s *saver) Close() {
	s.ticker.Stop()
	s.flush()
}

func (s *saver) flush() {
	if len(s.entities) == 0 {
		return
	}

	err := s.flusher.Flush(s.entities)
	if err == nil {
		s.entities = make([]travel.Trip, 0)
	} else {
		// Log error
	}
}
