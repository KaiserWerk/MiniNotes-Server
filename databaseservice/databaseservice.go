package databaseservice

import (
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"mininotes-server/config"
	"mininotes-server/entity"
	"sync"
)

type databaseService struct {
	db *gorm.DB
}

var (
	dbInit sync.Once
	service databaseService
)

// Get returns the database service connection
func Get() *databaseService {
	dbInit.Do(func() {
		cfg := config.GetConfig()

		var driver gorm.Dialector
		if cfg.Database.Driver == "sqlite" {
			driver = sqlite.Open(cfg.Database.DSN)
		} else if cfg.Database.Driver == "mysql" {
			driver = mysql.Open(cfg.Database.DSN)
		} else if cfg.Database.Driver == "postgres" {
			driver = postgres.Open(cfg.Database.DSN)
		}

		db, err := gorm.Open(driver, &gorm.Config{
			PrepareStmt: true,
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true,
				NoLowerCase:   false,
			},
		})
		if err != nil {
			panic("gorm connection error: " + err.Error())
		}

		service = databaseService{db: db}
	})


	return &service
}

// AutoMigrate makes sure the database tables exist, corresponding
// to the supplied structs
func (ds databaseService) AutoMigrate() error {
	err := ds.db.AutoMigrate(
		&entity.User{},
	)
	if err != nil {
		return err
	}

	return nil
}
