package pkg

import (
	"fmt"
	"lcdb/config"
	"log"
	"sync"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBController struct {
	Config  config.DBConfig
	DB      *gorm.DB
	MutexDb sync.Mutex
}

func NewDbController(Conf config.DBConfig) *DBController {
	return &DBController{
		Config: Conf,
	}
}

func (dbc *DBController) Connect() (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		dbc.Config.Host, dbc.Config.User, dbc.Config.Password, dbc.Config.Name, dbc.Config.Port)
	log.Printf("Attemptin to connect to %s:%s", dbc.Config.Host, dbc.Config.Port)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	return db, err
}

func (dbc *DBController) Migrate(model interface{}) error {
	log.Printf("Migrating the table %v for db %s", model, dbc.Config.Name)
	err := dbc.DB.AutoMigrate(model)
	return err
}
