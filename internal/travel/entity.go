package travel

import "fmt"

type Trip struct {
	UserID       uint64
	FromLocation string
	DestLocation string
}

func (trip Trip) String() string {
	return fmt.Sprintf("[%d, %s, %s]", trip.UserID, trip.FromLocation, trip.DestLocation)
}
