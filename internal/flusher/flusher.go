package flusher

import (
	"github.com/ozonva/ova-travel-api/internal/repo"
	"github.com/ozonva/ova-travel-api/internal/travel"
	"github.com/ozonva/ova-travel-api/internal/utils"
)

type FlusherProvider interface {
	Flush(entities []travel.Trip) error
}

func NewFlusher(chunkSize int,
	entityRepo repo.Repo) FlusherProvider {
	return &flusher{
		chunkSize:  chunkSize,
		entityRepo: entityRepo,
	}
}

type flusher struct {
	chunkSize  int
	entityRepo repo.Repo
}

func (a flusher) Flush(entities []travel.Trip) error {
	for _, values := range utils.SplitByBatch(entities, a.chunkSize) {
		if err := a.entityRepo.AddEntities(values); err != nil {
			return err
		}
	}

	return nil
}
