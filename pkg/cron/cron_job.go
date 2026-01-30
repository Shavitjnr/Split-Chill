package cron

import (
	"fmt"
	"time"

	"github.com/Shavitjnr/split-chill-ai/pkg/core"
	"github.com/Shavitjnr/split-chill-ai/pkg/duplicatechecker"
	"github.com/Shavitjnr/split-chill-ai/pkg/log"
	"github.com/Shavitjnr/split-chill-ai/pkg/utils"
)


type CronJob struct {
	Name        string
	Description string
	Period      CronJobPeriod
	Run         func(*core.CronContext) error
}

func (j *CronJob) doRun() {
	start := time.Now()
	c := core.NewCronJobContext(j.Name, j.Period.GetInterval())

	if duplicatechecker.Container.IsEnabled() {
		localAddr, err := utils.GetLocalIPAddressesString()

		if err != nil {
			log.Warnf(c, "[cron_job.doRun] job \"%s\" cannot get local ipv4 address, because %s", j.Name, err.Error())
			return
		}

		currentInfo := fmt.Sprintf("ip: %s, startTime: %d", localAddr, time.Now().Unix())
		found, runningInfo := duplicatechecker.Container.GetOrSetCronJobRunningInfo(j.Name, currentInfo, j.Period.GetInterval())

		if found {
			log.Warnf(c, "[cron_job.doRun] job \"%s\" is already running (%s)", j.Name, runningInfo)
			return
		}
	}

	err := j.Run(c)

	now := time.Now()

	if err != nil {
		log.Errorf(c, "[cron_job.doRun] failed to run job \"%s\", because %s", j.Name, err.Error())
		return
	}

	cost := now.Sub(start).Nanoseconds() / 1e6
	log.Infof(c, "[cron_job.doRun] run job \"%s\" successfully, cost %dms", j.Name, cost)
}
