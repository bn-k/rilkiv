package exchange

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	Orm
	Email        string `json:"email" gorm:"uniqueIndex"`
	Firstname    string `json:"firstname"`
	Lastname     string `json:"lastname"`
	Auth
}

type Orm struct {
	ID        uuid.UUID      `json:"id" gorm:"primarykey;type:uuid;default:gen_random_uuid()"`
	CreatedAt time.Time      `json:"created_at" gorm:"type:timestamp"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"type:timestamp"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index,type:timestamp"`
}

type Auth struct {
	Password     string `json:"password"`
	ConfirmToken string `json:"confirm_token"`
	Confirmed bool `json:"confirmed"`
}
