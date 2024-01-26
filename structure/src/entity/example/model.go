package example

import (
	"github.com/dalikewara/ayapingping-go/v3/structure/src/library/timeLib"
	"time"
)

type Model struct {
	ID        uint64    `json:"id" db:"id"`
	Name      string    `json:"name" db:"name"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

// SetCreatedAtNowUTC sets created at to time now in UTC
func (model *Model) SetCreatedAtNowUTC() {
	model.CreatedAt = timeLib.NowUTC()
}
