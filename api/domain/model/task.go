package model

import (
	"time"
)

type Task struct {
	Id          int64     `json:"id" xorm:"pk autoincr BIGINT(20)"`
	UserId      int64     `json:"userId" xorm:"not null BIGINT(20)"`
	Title       string    `json:"title" xorm:"not null index VARCHAR(60)"`
	ElapsedTime int64     `json:"elapsedTime" xorm:"not null BIGINT(20)"`
	Status      int       `json:"status" xorm:"not null TINYINT(4)"`
	CreatedAt   time.Time `json:"createdAt" xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
	UpdatedAt   time.Time `json:"updatedAt" xorm:"not null default 'CURRENT_TIMESTAMP' TIMESTAMP"`
}
