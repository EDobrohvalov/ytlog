package persistence

import (
	"errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"
	"time"
	"ytlog/internal/config"
)

func NewDatabase(cfg *config.Config) *Database {
	db, err := gorm.Open(postgres.Open(cfg.ConnectionString),
		&gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				TablePrefix:   "ytlog_",
				SingularTable: false,
			}})
	if err != nil {
		panic(err.Error())
	}
	return &Database{
		DB: *db,
	}
}

func (db *Database) CreateTables() error {
	err := db.AutoMigrate(&User{}, &Issue{}, &IssueLog{}, &SyncLog{})
	return err
}

func (db *Database) SaveUsers(items *[]User) {
	db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "uid"}},
		DoNothing: true,
	}).Create(&items)
}

func (db *Database) SaveIssues(items *[]Issue) {
	db.Clauses(clause.OnConflict{
		Columns:   []clause.Column{{Name: "key"}},
		DoUpdates: clause.AssignmentColumns([]string{"status", "priority", "complexity", "summary", "type", "spent"}),
	}).Create(&items)
}

func (db *Database) SaveIssueLog(items *[]IssueLog) {
	db.Clauses(clause.OnConflict{
		DoNothing: true,
	}).Create(&items)
}

func (db *Database) GetLastSync() SyncLog {
	var syncTimestamp SyncLog
	result := db.Last(&syncTimestamp)
	errors.Is(result.Error, gorm.ErrRecordNotFound)
	return syncTimestamp
}

func (db *Database) UpdateSync() {
	db.Create(&SyncLog{UpdatedAt: time.Now()})
}
