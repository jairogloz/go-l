package time

import (
	"time"

	"github.com/jairogloz/go-l/pkg/ports"
)

// Make sure Clock implements ports.Clock
// at compile time
var _ ports.Clock = &Clock{}

// Clock is a struct that represents the clock.
type Clock struct{}

// Now returns the current time.
func (c *Clock) Now() time.Time {
	return time.Now()
}
