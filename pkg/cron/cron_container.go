package cron

import (
	"time"

	"github.com/go-co-op/gocron/v2"

	"github.com/Shavitjnr/split-chill-ai/pkg/core"
	"github.com/Shavitjnr/split-chill-ai/pkg/errs"
	"github.com/Shavitjnr/split-chill-ai/pkg/log"
	"github.com/Shavitjnr/split-chill-ai/pkg/settings"
)


type CronJobSchedulerContainer struct {
	scheduler        gocron.Scheduler
	allJobs          []*CronJob
	allJobsMap       map[string]*CronJob
	allGocronJobsMap map[string]gocron.Job
}


var (
	Container = &CronJobSchedulerContainer{
		allJobsMap:       make(map[string]*CronJob),
		allGocronJobsMap: make(map[string]gocron.Job),
	}
)


func InitializeCronJobSchedulerContainer(ctx core.Context, config *settings.Config, startScheduler bool) error {
	var err error

	Container.scheduler, err = gocron.NewScheduler(
		gocron.WithLocation(time.Local),
		gocron.WithLogger(NewGocronLoggerAdapter()),
	)

	if err != nil {
		return err
	}

	Container.registerAllJobs(ctx, config)

	if startScheduler {
		Container.scheduler.Start()
	}

	return nil
}


func (c *CronJobSchedulerContainer) GetAllJobs() []*CronJob {
	return c.allJobs
}


func (c *CronJobSchedulerContainer) SyncRunJobNow(jobName string) error {
	if jobName == "" {
		return errs.ErrCronJobNameIsEmpty
	}

	job := c.allJobsMap[jobName]

	if job == nil {
		return errs.ErrCronJobNotExistsOrNotEnabled
	}

	gocronJob := c.allGocronJobsMap[jobName]

	if gocronJob == nil {
		return errs.ErrCronJobNotExistsOrNotEnabled
	}

	job.doRun()
	return nil
}

func (c *CronJobSchedulerContainer) registerAllJobs(ctx core.Context, config *settings.Config) {
	if config.EnableRemoveExpiredTokens {
		Container.registerIntervalJob(ctx, RemoveExpiredTokensJob)
	}

	if config.EnableCreateScheduledTransaction {
		Container.registerIntervalJob(ctx, CreateScheduledTransactionJob)
	}
}

func (c *CronJobSchedulerContainer) registerIntervalJob(ctx core.Context, job *CronJob) {
	gocronJob, err := c.scheduler.NewJob(
		job.Period.ToJobDefinition(),
		gocron.NewTask(job.doRun),
		gocron.WithName(job.Name),
		gocron.WithSingletonMode(gocron.LimitModeReschedule),
	)

	if err == nil {
		c.allJobs = append(c.allJobs, job)
		c.allJobsMap[job.Name] = job
		c.allGocronJobsMap[job.Name] = gocronJob
		log.Infof(ctx, "[cron_container.registerJob] job \"%s\" has been registered", job.Name)
		log.Debugf(ctx, "[cron_container.registerJob] job \"%s\" gocron id is %s", job.Name, gocronJob.ID())
	} else {
		log.Errorf(ctx, "[cron_container.registerJob] job \"%s\" cannot be been registered, because %s", job.Name, err.Error())
	}
}
