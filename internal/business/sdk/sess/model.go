package sess

import (
	"time"
)

type Session struct {
	ID        string
	Data      string
	CreatedAt time.Time
	UpdatedAt time.Time
	ExpiresAt time.Time
}
