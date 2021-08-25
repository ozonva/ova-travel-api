package flusher

import (
	"github.com/ozonva/ova-travel-api/internal/repo"
	"github.com/ozonva/ova-travel-api/internal/travel"
	"github.com/ozonva/ova-travel-api/internal/utils"
)

type Flusher interface {
	Flush(entities []travel.Trip) []travel.Trip
}

func NewFlusher(chunkSize int,
	entityRepo repo.Repo) Flusher {
	return &flusher{
		chunkSize:  chunkSize,
		entityRepo: entityRepo,
	}
}

type flusher struct {
	chunkSize  int
	entityRepo repo.Repo
}

func (a flusher) Flush(entities []travel.Trip) []travel.Trip {
	unsaved := make([]travel.Trip, 0)
	for i := 0; i < len(entities); i += a.chunkSize {
		maxIdx := utils.MinInt(i+a.chunkSize, len(entities))
		if err := a.entityRepo.AddEntities(entities[i:maxIdx]); err != nil {
			unsaved = append(unsaved, entities[i:maxIdx]...)
		}
	}

	return unsaved
}
