package ports

import "time"

// Clock is an interface that represents the clock.
type Clock interface {
	Now() time.Time
}
