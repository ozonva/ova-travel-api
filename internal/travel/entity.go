package travel

import "fmt"

type Trip struct {
	UserID       int
	FromLocation string
	DestLocation string
}

func (trip Trip) String() string {
	return fmt.Sprintf("[%d, %s, %s]", trip.UserID, trip.FromLocation, trip.DestLocation)
}
