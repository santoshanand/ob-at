package angleone

import (
	"github.com/angel-one/smartapigo"
	"github.com/santoshanand/at/modules/common/config"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// IAngelOne - interface of dao
type IAngelOne interface {
}

type angleone struct {
	log *zap.SugaredLogger
	cfg *config.Config
	db  *gorm.DB
}

// NewAngelOne - instance of angel one
func NewAngelOne(log *zap.SugaredLogger, cfg *config.Config) IAngelOne {
	abClient := smartapigo.New("S50542509", "9945", "Qe4fayZl")
	session, err := abClient.GenerateSession("181179")

	if err != nil {
		log.Debug("error to generate session, ", err)
	}

	log.Debug("session ", session.AccessToken)

	session.UserProfile, err = abClient.GetUserProfile()

	if err != nil {
		log.Debug(err.Error())
	}
	return &angleone{
		log: log,
		cfg: cfg,
	}
}

// // Module - database module
// var Module = fx.Options(
// 	fx.Provide(NewAngelOne),
// )
