package dao

import (
	"errors"

	"github.com/santoshanand/at/modules/app/dto"
	"github.com/santoshanand/at/modules/common/config"
	"github.com/santoshanand/at/modules/common/database/entities"
	"github.com/santoshanand/at/modules/common/utils"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// IUser - user interface
type IUser interface {
	Upsert(user *entities.User) (*entities.User, error)
	GetLogInUser(loginDTO dto.LoginDTO) (*entities.User, error)
	GetSuperToken() (*string, error)
	Logout(logoutDTO dto.LogoutDTO) error
}

type user struct {
	log *zap.SugaredLogger
	cfg *config.Config
	db  *gorm.DB
}

// GetSuperToken implements IUser
func (u *user) GetSuperToken() (*string, error) {
	user := viper.GetString("ADMIN_USER")
	r := &entities.User{}
	if res := u.db.Where("user_id=?", user).First(r); res.Error != nil {
		return nil, res.Error
	}
	return &r.AccessToken, nil
}

// GetLogInUser implements IUser
func (u *user) Logout(logoutDTO dto.LogoutDTO) error {
	user, err := u.getUser(logoutDTO.UserID, logoutDTO.Broker)
	if err != nil {
		return err
	}
	user.StopTrading = true
	user.LoginAt = utils.CurrentTime().AddDate(0, 0, -4)
	_, err = u.update(user)
	if err != nil {
		return err
	}
	return nil
}

// GetLogInUser implements IUser
func (u *user) GetLogInUser(loginDTO dto.LoginDTO) (*entities.User, error) {
	user, err := u.getUser(loginDTO.UserID, loginDTO.Broker)
	if err != nil {
		return nil, err
	}

	loginAt := utils.FormatTimeToYYYYMMDDDate(user.LoginAt)
	currentTime := utils.FormatTimeToYYYYMMDDDate(utils.CurrentTime())

	if loginAt != currentTime {
		return nil, errors.New("session expired")
	}
	return user, nil
}

// getUser implements IUser
func (u *user) getUser(userID, broker string) (*entities.User, error) {
	r := &entities.User{}
	if res := u.db.Where("user_id=? and broker=? and access_token!=?", userID, broker, "").First(r); res.Error != nil {
		return nil, res.Error
	}
	return r, nil
}

// update implements IUser
func (u *user) update(e *entities.User) (*entities.User, error) {
	if res := u.db.Where("user_id=? and broker=?", e.UserID, e.Broker).Updates(e); res.Error != nil {
		return nil, res.Error
	}
	return e, nil
}

// Upsert implements IUser
func (u *user) Upsert(e *entities.User) (*entities.User, error) {
	user, err := u.getUser(e.UserID, e.Broker)
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		e.LoginAt = utils.CurrentTime()
		if res := u.db.Save(e); res.Error != nil {
			return nil, res.Error
		}
		return e, nil
	}
	user.LoginAt = utils.CurrentTime()
	user.AccessToken = e.AccessToken
	if res := u.db.Where("user_id=? and broker=?", e.UserID, user.Broker).Updates(user); res.Error != nil {
		return nil, res.Error
	}
	return user, nil
}

// NewUser - instance of user
func NewUser(param *params) IUser {
	return &user{
		log: param.log,
		cfg: param.cfg,
		db:  param.db,
	}
}
