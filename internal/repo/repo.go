package repo

import (
	"database/sql"
	"fmt"

	sq "github.com/Masterminds/squirrel"

	"github.com/ozonva/ova-travel-api/internal/travel"
)

const (
	tableName  = "trips"
	idColumn   = "id"
	fromColumn = "from"
	destColumn = "dest"
)

type Repo interface {
	AddEntities(entities []travel.Trip) error
	ListEntities(limit, offset uint64) ([]travel.Trip, error)
	UpdateEntity(entityID uint64, newTrip *travel.Trip) error
	DescribeEntity(entityId uint64) (*travel.Trip, error)
	RemoveEntity(entityId uint64) error
}

func NewRepo(db *sql.DB) Repo {
	return &repoImpl{db: db}
}

type repoImpl struct {
	db *sql.DB
}

func (r repoImpl) UpdateEntity(entityID uint64, newTrip *travel.Trip) error {
	result, err := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).
		Update(tableName).
		Set(fromColumn, newTrip.FromLocation).
		Set(destColumn, newTrip.DestLocation).
		Where(sq.Eq{idColumn: entityID}).
		RunWith(r.db).
		Exec()

	if err != nil {
		return fmt.Errorf("cannot run update query: %w", err)
	}

	_, err = result.RowsAffected()
	if err != nil {
		return fmt.Errorf("cannot get rows affected: %w", err)
	}

	return nil
}

func (r repoImpl) AddEntities(entities []travel.Trip) error {
	if len(entities) == 0 {
		return nil
	}

	sqlStmt, _, err := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).
		Insert(tableName).Columns(fromColumn, destColumn).
		Values("", "").ToSql()

	if err != nil {
		return fmt.Errorf("failed to build sql template: %w", err)
	}

	tx, err := r.db.Begin()
	if err != nil {
		return fmt.Errorf("failer to start transaction: %w", err)
	}

	stmt, err := tx.Prepare(sqlStmt)
	if err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to build sql template: %w", err)
	}
	defer stmt.Close()

	for i := 0; i < len(entities); i++ {
		if _, err := stmt.Exec(entities[i].FromLocation, entities[i].DestLocation); err != nil {
			tx.Rollback()
			return fmt.Errorf("cannot fill prepared statement: %w", err)
		}
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("transaction failed: %w", err)
	}

	return nil
}

func (r repoImpl) ListEntities(limit, offset uint64) ([]travel.Trip, error) {
	trips, err := sq.Select("*").
		From(tableName).
		OrderBy(idColumn).
		Limit(limit).
		Offset(offset).
		RunWith(r.db).Query()

	if err != nil {
		return nil, fmt.Errorf("cannot run list query: %w", err)
	}
	defer trips.Close()

	result := make([]travel.Trip, 0, limit)

	for trips.Next() {
		var trip travel.Trip
		if err := trips.Scan(&trip.UserID, &trip.FromLocation, &trip.DestLocation); err != nil {
			return nil, fmt.Errorf("cannot parse trip: %w", err)
		}
		result = append(result, trip)
	}
	if err := trips.Err(); err != nil {
		return nil, fmt.Errorf("error list query %w", err)
	}

	return result, nil
}

func (r repoImpl) DescribeEntity(entityId uint64) (*travel.Trip, error) {
	trips, err := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).
		Select("*").
		From(tableName).
		Where(sq.Eq{idColumn: entityId}).
		RunWith(r.db).
		Query()

	if err != nil {
		return nil, fmt.Errorf("cannot run describe query: %w", err)
	}
	defer trips.Close()

	if !trips.Next() {
		return nil, nil
	}

	var trip = &travel.Trip{}

	if err := trips.Scan(&trip.UserID, &trip.FromLocation, &trip.DestLocation); err != nil {
		return nil, fmt.Errorf("cannot parse trip: %w", err)
	}

	if err := trips.Err(); err != nil {
		return nil, fmt.Errorf("error list query %w", err)
	}

	return trip, nil
}

func (r repoImpl) RemoveEntity(entityId uint64) error {
	result, err := sq.StatementBuilder.PlaceholderFormat(sq.Dollar).
		Delete("").
		From(tableName).
		Where(sq.Eq{idColumn: entityId}).
		RunWith(r.db).
		Exec()

	if err != nil {
		return fmt.Errorf("cannot run delete query: %w", err)
	}

	_, err = result.RowsAffected()
	if err != nil {
		return fmt.Errorf("cannot get rows affected: %w", err)
	}

	return nil
}
