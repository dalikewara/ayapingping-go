package user

import (
	"github.com/dalikewara/ayapingping-go/v3/structure/src/entity/userProfile"
	"github.com/dalikewara/ayapingping-go/v3/structure/src/library/timeLib"
	"time"
)

type Model struct {
	ID        uint64    `json:"id" db:"id"`
	Username  string    `json:"username" db:"username"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

type ModelWithDetail struct {
	Model
	Profile *userProfile.Model `json:"profile"`
}

// SetCreatedAtNow sets created at to time now
func (model *Model) SetCreatedAtNow() {
	model.CreatedAt = timeLib.NowUTC()
}
