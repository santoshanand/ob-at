package dao

import (
	"testing"

	"github.com/santoshanand/at/modules/common/database/entities"
	"github.com/stretchr/testify/assert"
)

func TestCreateTrader(t *testing.T) {
	t.Parallel()
	dao := getTestDAO()
	trader := &entities.Trader{
		UserID:     "XY200",
		ProfileRaw: "{}",
		Token:      "B17jpUacUWmecuaWoS87afqMKJNdXyNeIGpbulsV70dSruSpC8mlOyRAYaHEPdyrItuzANCJzYDn6FnETQLs6EATn83AL1faAdDxiz5QR1cn0Z0cUrIFiA==",
		Name:       "test",
		AvatarURL:  "test.png",
		IsBlocked:  false,
	}
	err := dao.NewTraderDao().Create(trader)
	assert.Nil(t, err)
}

func TestUpdateTrader(t *testing.T) {
	t.Parallel()
	dao := getTestDAO()
	trader := &entities.Trader{
		UserID:     "XY200",
		ProfileRaw: "",
		Token:      "B17jpUacUWmecuaWoS87afqMKJNdXyNeIGpbulsV70dSruSpC8mlOyRAYaHEPdyrItuzANCJzYDn6FnETQLs6EATn83AL1faAdDxiz5QR1cn0Z0cUrIFiA==",
		Name:       "testsss",
		AvatarURL:  "",
		IsBlocked:  false,
	}
	err := dao.NewTraderDao().Update(trader)
	assert.Nil(t, err)
}

func TestGetTrader_Error(t *testing.T) {
	t.Parallel()
	dao := getTestDAO()
	userID := "XY200s"
	trader, err := dao.NewTraderDao().Get(userID)
	assert.NotNil(t, err)
	assert.Nil(t, trader)
}

func TestGetTrader_Success(t *testing.T) {
	t.Parallel()
	dao := getTestDAO()
	userID := "YA9556"
	trader, err := dao.NewTraderDao().Get(userID)
	assert.Nil(t, err)
	assert.NotNil(t, trader)
	assert.Equal(t, trader.UserID, userID)
}

func TestDeleteTrader(t *testing.T) {
	t.Parallel()
	dao := getTestDAO()
	userID := "XY200"
	err := dao.NewTraderDao().Delete(userID)
	assert.Nil(t, err)
}

func TestGetAllTrader(t *testing.T) {
	t.Parallel()
	dao := getTestDAO()
	traders, err := dao.NewTraderDao().GetAll()
	assert.Nil(t, err)
	assert.NotNil(t, traders)
}

func TestStopTrading(t *testing.T) {
	t.Parallel()
	dao := getTestDAO()
	userID := "XY200"
	err := dao.NewTraderDao().StopTrading(userID, "", true)
	assert.Nil(t, err)
}
