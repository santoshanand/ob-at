package job

import (
	"time"

	"github.com/robfig/cron"
	"github.com/santoshanand/at/modules/common/config"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

const asiaTimezone = "Asia/Kolkata"

// Params - cron job struct
type params struct {
	cron *cron.Cron
	log  *zap.SugaredLogger
	cfg  *config.Config
}

// scheduleCronJob - used to create a single instance of the service
func (p *params) scheduleCronJob() error {
	log := p.log

	err := p.cron.AddFunc("6 18 9 * * 1-5", p.morningTrade)
	if err != nil {
		log.Debug("error to add refresh instrument cron:", err.Error())
		return err
	}
	log.Debug("total scheduled job: ", len(p.cron.Entries()))
	return nil
}

func (p *params) morningTrade() {
	p.log.Debug("start morningTrade cron")
	defer p.log.Debug("end morningTrade cron")
}

// NewCron - instance of cron job
func NewCron(log *zap.SugaredLogger, cfg *config.Config) *cron.Cron {
	loc, err := time.LoadLocation(asiaTimezone)
	if err != nil {
		panic(err)
	}
	cronJob := cron.NewWithLocation(loc)
	p := &params{
		cron: cronJob,
		log:  log,
		cfg:  cfg,
	}
	err = p.scheduleCronJob()
	if err != nil {
		log.Debug("unable to schedule cron job: ", err.Error())
		return nil
	}
	cronJob.Start()

	entries := p.cron.Entries()
	for _, entry := range entries {
		p.log.Debug("Cron : ", entry.Next)
	}
	return cronJob
}

// Module provided to fx
var Module = fx.Options(
	fx.Provide(NewCron),
)
