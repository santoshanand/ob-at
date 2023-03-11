package dao

import (
	"github.com/santoshanand/at/modules/common/config"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// IUser - user interface
type IUser interface {
}

type params struct {
	log *zap.SugaredLogger
	cfg *config.Config
	db  *gorm.DB
}

// NewUser - instance of user
func NewUser(param *dao) IUser {
	return &params{
		log: param.log,
		cfg: param.cfg,
		db:  param.db,
	}
}
