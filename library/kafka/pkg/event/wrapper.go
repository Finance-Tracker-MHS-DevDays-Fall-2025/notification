package event

import (
	"time"

	"github.com/google/uuid"
)

type Type = string

type Wrapper struct {
	RawEvent  []byte    `json:"raw_event"`
	RequestID uuid.UUID `json:"request_id"`
	Timestamp time.Time `json:"timestamp"`
}
