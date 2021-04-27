/*
Package gigasecond points one gigasecond after the specified date

The syntax is accepted as
*/
package gigasecond

import (
	"time"
)

// gs is a constant symbols a gigasecond
const gs = time.Duration(1e9) * time.Second

// AddGigasecond adds one gigasecond to specified time
func AddGigasecond(t time.Time) time.Time {
	t = t.Add(gs)

	return t
}
