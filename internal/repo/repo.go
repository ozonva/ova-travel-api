package repo

import (
	"github.com/ozonva/ova-travel-api/internal/travel"
)

type Repo interface {
	AddEntities(entities []travel.Trip) error
	ListEntities(limit, offset uint64) ([]travel.Trip, error)
	DescribeEntity(entityId uint64) (*travel.Trip, error)
}
