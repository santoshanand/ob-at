package database

import (
	"errors"
	"time"

	"github.com/santoshanand/at/modules/common/config"
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
	// d.migrate(db)
	d.log.Debug("pgsql connected")
	return db, err
}

// Module - database module
var Module = fx.Options(
	fx.Provide(NewDatabase),
)
