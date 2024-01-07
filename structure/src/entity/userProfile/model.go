package userProfile

import (
	"github.com/dalikewara/ayapingping-go/v3/structure/src/library/timeLib"
	"time"
)

type Model struct {
	ID        uint64    `json:"id" db:"id"`
	UserID    uint64    `json:"user_id" db:"user_id"`
	Name      string    `json:"name" db:"name"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

// SetCreatedAtNow sets created at to time now
func (model *Model) SetCreatedAtNow() {
	model.CreatedAt = timeLib.NowUTC()
}
