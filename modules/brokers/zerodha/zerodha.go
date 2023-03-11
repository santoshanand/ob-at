package zerodha

import (
	"errors"

	"github.com/santoshanand/at-kite/kite"
	"github.com/santoshanand/at/modules/common/config"
	"github.com/santoshanand/at/modules/common/database/dao"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

// IZerodha -
type IZerodha interface {
	Login(req LoginDTO) (*kite.UserProfile, error)
}

type zerodha struct {
	kConnect *kite.Client
	log      *zap.SugaredLogger
	cfg      *config.Config
	dao      dao.IDao
}

// Login implements IZerodha
func (z *zerodha) Login(req LoginDTO) (*kite.UserProfile, error) {
	z.kConnect.SetAccessToken(req.Token)
	profile, err := z.kConnect.GetUserProfile()
	if err != nil {
		return nil, err
	}
	if profile.UserID != req.UserID {
		return nil, errors.New("user id is not correct")
	}
	return &profile, nil
}

// NewZerodha - new instance of zerodha
func newZerodha(log *zap.SugaredLogger, cfg *config.Config, dao dao.IDao) IZerodha {
	return &zerodha{
		kConnect: kite.New(""),
		log:      log,
		cfg:      cfg,
		dao:      dao,
	}
}

// Module provided to fx
var Module = fx.Options(
	fx.Provide(newZerodha),
)
