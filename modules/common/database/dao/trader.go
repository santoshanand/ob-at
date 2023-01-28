package dao

import (
	"errors"

	"github.com/santoshanand/at/modules/common/config"
	"github.com/santoshanand/at/modules/common/database/entities"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// ITrader -
type ITrader interface {
	Create(*entities.Trader) error
	Upsert(*entities.Trader) (*entities.Trader, error)
	Update(*entities.Trader) error
	StopTrading(userID, token string, isBlocked bool) error
	GetTraderByTokenAndUserID(userID, token string) (*entities.Trader, error)
	GetTraderByToken(token string) (*entities.Trader, error)
	Get(userID string) (*entities.Trader, error)
	GetAdminToken() (string, error)
	GetByToken(token string) (*entities.Trader, error)
	GetAll() ([]*entities.Trader, error)
	GetAllValidTrader() ([]*entities.Trader, error)
	Delete(userID string) error
}

type trader struct {
	log *zap.SugaredLogger
	cfg *config.Config
	db  *gorm.DB
}

func (t *trader) GetAdminToken() (string, error) {
	user, err := t.Get(t.cfg.Username)
	if err != nil {
		return "", err
	}
	return user.Token, nil
}

// GetTraderByToken implements ITrader
func (t *trader) GetTraderByToken(token string) (*entities.Trader, error) {
	r := &entities.Trader{}
	if res := t.db.Where("token = ?", token).First(r); res.Error != nil {
		if gorm.ErrRecordNotFound == res.Error {
			return nil, errors.New("unauthorized access")
		}
		return nil, res.Error
	}
	return r, nil
}

// GetTraderByTokenAndUserID implements ITrader
func (t *trader) GetTraderByTokenAndUserID(userID string, token string) (*entities.Trader, error) {
	eTrader := &entities.Trader{}
	if res := t.db.Where("user_id = ? and token = ? and is_blocked = ? ", userID, token, false).First(&eTrader); res.Error != nil {
		if gorm.ErrRecordNotFound == res.Error {
			return nil, errors.New("unauthorized access")
		}
		return nil, res.Error
	}
	return eTrader, nil
}

// StopTrading implements ITrader
func (t *trader) StopTrading(userID, token string, value bool) error {

	//&entities.Trader{IsBlocked: blocked, IsValidToken: validToken}
	updateValues := map[string]interface{}{"is_stop_trading": value}
	if res := t.db.Table("traders").Where("user_id = ? and token = ?", userID, token).Updates(updateValues); res.Error != nil {
		return res.Error
	}
	return nil
}

// GetAll implements ITrader
func (t *trader) GetAll() ([]*entities.Trader, error) {
	var traders []*entities.Trader
	if res := t.db.Where("is_blocked = ?", false).Find(&traders); res.Error != nil {
		return nil, res.Error
	}
	return traders, nil
}

// GetAllValidTrader implements ITrader
func (t *trader) GetAllValidTrader() ([]*entities.Trader, error) {
	var traders []*entities.Trader
	if res := t.db.Where("is_blocked = ? and is_valid_token = ? and is_stop_trading = ?", false, true, false).Find(&traders); res.Error != nil {
		return nil, res.Error
	}
	return traders, nil
}

// Delete implements ITrader
func (t *trader) Delete(userID string) error {
	r := &entities.Trader{}
	if res := t.db.Where("user_id = ?", userID).Delete(r); res.Error != nil {
		return res.Error
	}
	return nil
}

// Get implements ITrader
func (t *trader) Get(userID string) (*entities.Trader, error) {
	r := &entities.Trader{}
	if res := t.db.Where("user_id = ? and is_valid_token = ?", userID, true).First(r); res.Error != nil {
		if gorm.ErrRecordNotFound == res.Error {
			return nil, errors.New("unauthorized access")
		}
		return nil, res.Error
	}
	return r, nil
}

// getByUserID implements ITrader
func (t *trader) getByUserID(userID string) (*entities.Trader, error) {
	r := &entities.Trader{}
	if res := t.db.Where("user_id = ?", userID).First(r); res.Error != nil {
		return nil, res.Error
	}
	return r, nil
}

// GetByToken implements ITrader
func (t *trader) GetByToken(token string) (*entities.Trader, error) {
	r := &entities.Trader{}
	if res := t.db.Where("token = ?", token).First(r); res.Error != nil {
		return nil, res.Error
	}
	return r, nil
}

// Update implements ITrader
func (t *trader) Update(e *entities.Trader) error {
	if res := t.db.Where("user_id = ?", e.UserID).Updates(e); res.Error != nil {
		return res.Error
	}
	return nil
}

// Create implements ITrader
func (t *trader) Create(e *entities.Trader) error {
	if res := t.db.Create(e); res.Error != nil {
		return res.Error
	}
	return nil
}

// Upsert implements ITrader
func (t *trader) Upsert(e *entities.Trader) (*entities.Trader, error) {
	traderRes, err := t.getByUserID(e.UserID)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		if res := t.db.Create(e); res.Error != nil {
			return nil, res.Error
		}
		return e, nil
	}
	traderRes.Token = e.Token
	traderRes.IsValidToken = e.IsValidToken
	traderRes.IsStopTrading = false

	err = t.Update(traderRes)
	if err != nil {
		return nil, err
	}
	traderRes, err = t.getByUserID(e.UserID)
	return traderRes, err
}

// NewTrader - new trader instance
func NewTrader(param *dao) ITrader {
	return &trader{
		log: param.log,
		cfg: param.cfg,
		db:  param.db,
	}
}
