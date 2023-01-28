package dao

import (
	"testing"

	"github.com/santoshanand/at/modules/common/config"
	"github.com/santoshanand/at/modules/common/database"
	"github.com/santoshanand/at/modules/common/logger"

	"github.com/stretchr/testify/assert"
)

func getTestDAO() IDao {
	log := logger.NewLogger()
	cfg := config.NewConfig()
	db, err := database.NewDatabase(log, cfg)
	if err != nil {
		return nil
	}
	dao := NewDao(log, cfg, db)
	return dao
}

func TestNewDao(t *testing.T) {
	t.Parallel()

	dao := getTestDAO()
	assert.NotNil(t, dao)
}
