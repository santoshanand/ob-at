package dao

import (
	"github.com/santoshanand/at/modules/common/config"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// IDao - interface of dao
type IDao interface {
	NewUserDao() IUser
}

type dao struct {
	log *zap.SugaredLogger
	cfg *config.Config
	db  *gorm.DB
}

// NewUserDao implements IDao
func (d *dao) NewUserDao() IUser {
	return NewUser(d.getParams())
}

func (d *dao) getParams() *dao {
	return d
}

// NewDao -
func NewDao(log *zap.SugaredLogger, cfg *config.Config, db *gorm.DB) IDao {
	return &dao{
		log: log,
		cfg: cfg,
		db:  db,
	}
}

// Module - database module
var Module = fx.Options(
	fx.Provide(NewDao),
)
