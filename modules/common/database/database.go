package database

import (
	"errors"
	"time"

	"github.com/santoshanand/at/modules/common/config"
	"github.com/santoshanand/at/modules/common/database/entities"
	"github.com/santoshanand/at/modules/common/utils"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type database struct {
	log *zap.SugaredLogger
	cfg *config.Config
}

// NewDatabase - create new instance of database
func NewDatabase(log *zap.SugaredLogger, cfg *config.Config) (*gorm.DB, error) {
	db := &database{log: log, cfg: cfg}
	return db.openPostgresDB()
}

func (d *database) openPostgresDB() (*gorm.DB, error) {
	url := d.cfg.PostgresDBURL
	if utils.IsEmpty(url) {
		return nil, errors.New("url is empty")
	}
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{
		// Logger: logger.Default.LogMode(logger.Info),
		NowFunc: func() time.Time {
			return utils.CurrentTime()
		},
	})
	if err != nil {
		panic(err)
	}
	d.migrate(db)
	d.log.Debug("db connected")
	return db, err
}

func (d *database) migrate(db *gorm.DB) {
	err := db.AutoMigrate(
		&entities.User{},
		&entities.Setting{},
	)
	if err != nil {
		d.log.Debug("migration error: ", err.Error())
		return
	}
}

// Module - database module
var Module = fx.Options(
	fx.Provide(NewDatabase),
)
