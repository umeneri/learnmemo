package model

import (
	"time"
)

type User struct {
	Id         int64     `xorm:"pk autoincr BIGINT(20)"`
	Email      string    `xorm:"not null index VARCHAR(255)"`
	Name       string    `xorm:"not null VARCHAR(255)"`
	ProviderId string    `xorm:"not null VARCHAR(100)"`
	AvaterUrl  string    `xorm:"not null VARCHAR(255)"`
	CreatedAt  time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
	UpdatedAt  time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
}
