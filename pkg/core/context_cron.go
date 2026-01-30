package core

import (
	"context"
	"strconv"
	"strings"
	"time"
)


type CronContext struct {
	context.Context
	contextId       string
	cronJobInterval time.Duration
}


func (c *CronContext) GetContextId() string {
	return c.contextId
}


func (c *CronContext) GetClientLocale() string {
	return ""
}


func (c *CronContext) GetInterval() time.Duration {
	return c.cronJobInterval
}


func NewCronJobContext(cronJobName string, cronJobInterval time.Duration) *CronContext {
	return &CronContext{
		Context:         context.Background(),
		contextId:       generateNewRandomCronContextId(cronJobName),
		cronJobInterval: cronJobInterval,
	}
}

func generateNewRandomCronContextId(cronJobName string) string {
	var ret strings.Builder
	ret.WriteString("cron-job-")
	ret.WriteString(strings.ToLower(cronJobName))
	ret.WriteRune('-')
	ret.WriteString(strconv.FormatInt(time.Now().Unix(), 10))

	return ret.String()
}
