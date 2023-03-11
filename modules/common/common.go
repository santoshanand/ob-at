package common

import (
	"github.com/santoshanand/at/modules/common/config"
	"github.com/santoshanand/at/modules/common/database"
	"github.com/santoshanand/at/modules/common/job"
	"github.com/santoshanand/at/modules/common/logger"
	"go.uber.org/fx"
)

// Module - common modules
var Module = fx.Options(
	config.Module,
	database.Module,
	job.Module,
	logger.Module,
)
