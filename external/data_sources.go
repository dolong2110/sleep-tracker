package external

import (
	"context"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"sleep-tracker/pkg/zapx"
	"time"
)

type DataSources struct {
	PostgreSQL *gorm.DB
}

func InitDS(config *Configs) (*DataSources, error) {
	log.Printf("Initializing data sources\n")
	zapx.Info(context.TODO(), "initializing data sources.")

	pg := config.DS.PostgreSQL
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", pg.Host, pg.Port, pg.User, pg.Password, pg.DB, pg.SSL)
	log.Printf("Connecting to Postgresql\n")
	zapx.Info(context.TODO(), "connecting to Postgresql ...")

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: dsn,
	}), &gorm.Config{})
	if err != nil {
		zapx.Error(context.TODO(), "error opening database.", err)
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		zapx.Error(context.TODO(), "failed to connect to SQL database.", err)
		return nil, err
	}

	if err = sqlDB.Ping(); err != nil {
		zapx.Error(context.TODO(), "error connecting to database.", err)
		return nil, err
	}

	sqlDB.SetConnMaxLifetime(time.Duration(pg.Timeout) * time.Minute)

	return &DataSources{
		PostgreSQL: db,
	}, nil
}

// Close to be used in graceful server shutdown
func (d *DataSources) Close() error {
	db, err := d.PostgreSQL.DB()
	if err != nil {
		zapx.Error(context.TODO(), "failed to connect to SQL database.", err)
		return err
	}

	if err = db.Close(); err != nil {
		zapx.Error(context.TODO(), "error closing database.", err)
		return err
	}

	return nil
}
