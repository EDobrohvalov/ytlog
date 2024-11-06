package persistence

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"ytlog/internal/config"
)

func NewDatabase(cfg *config.Config) *Database {
	db, err := gorm.Open(postgres.Open(cfg.ConnectionString), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	return &Database{
		DB: *db,
	}
}

func (db *Database) CreateTables() error {
	err := db.AutoMigrate(&User{}, &Issue{}, &IssueLog{})
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