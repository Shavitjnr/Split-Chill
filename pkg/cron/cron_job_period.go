package cron

import (
	"fmt"
	"time"

	"github.com/go-co-op/gocron/v2"
)


type CronJobPeriod interface {
	GetInterval() time.Duration
	ToJobDefinition() gocron.JobDefinition
}


type CronJobIntervalPeriod struct {
	Interval time.Duration
}


type CronJobFixedHourPeriod struct {
	Hour uint32
}


type CronJobEvery15MinutesPeriod struct {
	Second uint32
}


type CronJobFixedTimePeriod struct {
	Time time.Time
}


func (p CronJobIntervalPeriod) GetInterval() time.Duration {
	return p.Interval
}


func (p CronJobIntervalPeriod) ToJobDefinition() gocron.JobDefinition {
	return gocron.DurationJob(p.Interval)
}


func (p CronJobFixedHourPeriod) GetInterval() time.Duration {
	return 24 * time.Hour
}


func (p CronJobFixedHourPeriod) ToJobDefinition() gocron.JobDefinition {
	return gocron.DailyJob(
		1,
		gocron.NewAtTimes(
			gocron.NewAtTime(uint(p.Hour), 0, 0),
		),
	)
}


func (p CronJobEvery15MinutesPeriod) GetInterval() time.Duration {
	return 15 * time.Minute
}


func (p CronJobEvery15MinutesPeriod) ToJobDefinition() gocron.JobDefinition {
	return gocron.CronJob(fmt.Sprintf("%d */15 * * * *", p.Second), true)
}


func (p CronJobFixedTimePeriod) GetInterval() time.Duration {
	return 0
}


func (p CronJobFixedTimePeriod) ToJobDefinition() gocron.JobDefinition {
	return gocron.OneTimeJob(gocron.OneTimeJobStartDateTime(p.Time))
}
