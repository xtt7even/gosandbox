package bundle

import (
	"time"
)

// aka file
type Bundle struct {
	Id          string
	Source      string
	Destination string
	CreatedAt   time.Time
	TTL         time.Duration
	Payload     []byte
}
