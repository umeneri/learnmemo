package model

import (
	"time"
)

type Task struct {
	Id             int64     `xorm:"pk autoincr BIGINT(20)"`
	UserId         int64     `xorm:"not null BIGINT(20)"`
	Title          string    `xorm:"not null index VARCHAR(60)"`
	ProgressMinute int64     `xorm:"not null BIGINT(20)"`
	Status         int       `xorm:"not null TINYINT(4)"`
	CreatedAt      time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
	UpdatedAt      time.Time `xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
}
