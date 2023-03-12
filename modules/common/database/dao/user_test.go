package dao

import (
	"testing"

	"github.com/santoshanand/at/modules/app/dto"
	"github.com/santoshanand/at/modules/common/config"
	"github.com/santoshanand/at/modules/common/database"
	"github.com/santoshanand/at/modules/common/database/entities"
	"github.com/santoshanand/at/modules/common/logger"
	"github.com/stretchr/testify/assert"
)

func getTestDAO() *params {
	log := logger.NewLogger()
	cfg := config.NewConfig()
	db, err := database.NewDatabase(log, cfg)
	if err != nil {
		return nil
	}
	return &params{log: log, cfg: cfg, db: db}
}

func TestUpsert(t *testing.T) {
	t.Parallel()
	d := getTestDAO()
	user := NewUser(d)

	u := &entities.User{
		UserID:      "DEMO1",
		Broker:      "zerodha",
		Blocked:     false,
		StopTrading: false,
		Name:        "Test",
		AccessToken: "sss1",
		AvatarURL:   "",
	}

	res, err := user.Upsert(u)
	assert.NotNil(t, res)
	assert.Nil(t, err)
}

func TestGetLogInUser(t *testing.T) {
	t.Parallel()
	d := getTestDAO()
	user := NewUser(d)

	login := dto.LoginDTO{
		Token:  "sss1",
		UserID: "DEMO1",
		Broker: "zerodha",
	}
	res, err := user.GetLogInUser(login)
	assert.NotNil(t, res)
	assert.Nil(t, err)
}
