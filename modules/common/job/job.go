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

	err := p.cron.AddFunc("0 30 8 * * 1-5", p.saveOptionStocks)
	if err != nil {
		log.Debug("error to add refresh instrument cron:", err.Error())
		return err
	}
	//err = p.cron.AddFunc("6 */5 9-15 * * 1-5", p.fiveMinuteHandler) // It will run every five minutes monday to friday
	//if err != nil {
	//	log.Debug("error to schedule five minute cron: ", err.Error())
	//	return err
	//}

	err = p.cron.AddFunc("9 */5 9-15 * * 1-5", p.niftyHandler) // It will run every five minutes monday to friday
	if err != nil {
		log.Debug("error to schedule five minute cron: ", err.Error())
		return err
	}

	err = p.cron.AddFunc("0 32 15 * * 1-5", p.saveTransactions) // It will run every five minutes monday to friday
	if err != nil {
		log.Debug("error to schedule save transaction cron: ", err.Error())
		return err
	}

	log.Debug("total scheduled job: ", len(p.cron.Entries()))
	return nil
}

func (p *params) saveOptionStocks() {
	p.log.Debug("start save option stocks cron")
	defer p.log.Debug("end save save option stocks cron")
}

func (p *params) fiveMinuteHandler() {
	log := p.log
	log.Debug("start 5 minute cron")
	defer log.Debug("end 5 minute cron")

}

func (p *params) fifteenMinuteStrategy() {
	log := p.log
	log.Debug("start 15 minutes cron")
	defer log.Debug("end 15 minutes cron")
}

func (p *params) crudCron() {
	log := p.log
	log.Debug("start test cron")
	defer log.Debug("end test cron")
}

func (p *params) saveTransactions() {
	p.log.Debug("start save transactions cron")
	defer p.log.Debug("end save transactions cron")
}

func (p *params) niftyHandler() {
	p.log.Debug("start nifty cron")
	defer p.log.Debug("end nifty cron")

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
	// p.saveOptionStocks()
	return cronJob
}

// Module provided to fx
var Module = fx.Options(
	fx.Provide(NewCron),
)
