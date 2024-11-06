package persistence

import (
	"database/sql"
	"gorm.io/gorm"
	"time"
)

type SyncLog struct {
}

type User struct {
	Self                 string
	Uid                  int64 `gorm:"primaryKey"`
	Login                string
	TrackerUid           int64
	PassportUid          int
	CloudUid             string
	FirstName            string
	LastName             string
	Display              string
	Email                string
	External             bool
	HasLicense           bool
	Dismissed            bool
	UseNewFilters        bool
	DisableNotifications bool
	FirstLoginDate       string
	LastLoginDate        string
}

type Issue struct {
	Key        string `gorm:"primaryKey"`
	Type       string
	Status     string
	Summary    string
	Complexity string
	Priority   string
	Spent      string
	Queue      string
	CreatedAt  time.Time
}

type IssueLog struct {
	IssueKey   string    `gorm:"primaryKey"`
	UpdatedAt  time.Time `gorm:"primaryKey"`
	UpdatedBy  int64
	ChangeType string
	Field      string `gorm:"primaryKey"`
	FromValue  sql.NullString
	ToValue    sql.NullString
}

type Comment struct {
	Id        uint32 `gorm:"primaryKey"`
	IssueKey  string
	ChangedAt time.Time
	CreatedBy int64
	Text      string
}

type Database struct {
	gorm.DB
}
