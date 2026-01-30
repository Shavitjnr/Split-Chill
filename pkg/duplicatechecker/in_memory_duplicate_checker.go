package duplicatechecker

import (
	"fmt"
	"sync"
	"time"

	"github.com/patrickmn/go-cache"

	"github.com/Shavitjnr/split-chill-ai/pkg/settings"
)


type InMemoryDuplicateChecker struct {
	cache *cache.Cache

	mutex sync.Mutex
}


func NewInMemoryDuplicateChecker(config *settings.Config) (*InMemoryDuplicateChecker, error) {
	checker := &InMemoryDuplicateChecker{
		cache: cache.New(config.DuplicateSubmissionsIntervalDuration, config.InMemoryDuplicateCheckerCleanupIntervalDuration),
	}

	return checker, nil
}


func (c *InMemoryDuplicateChecker) GetSubmissionRemark(checkerType DuplicateCheckerType, uid int64, identification string) (bool, string) {
	existedRemark, found := c.cache.Get(c.getCacheKey(checkerType, uid, identification))

	if found {
		return true, existedRemark.(string)
	}

	return false, ""
}


func (c *InMemoryDuplicateChecker) SetSubmissionRemark(checkerType DuplicateCheckerType, uid int64, identification string, remark string) {
	c.cache.Set(c.getCacheKey(checkerType, uid, identification), remark, cache.DefaultExpiration)
}


func (c *InMemoryDuplicateChecker) SetSubmissionRemarkWithCustomExpiration(checkerType DuplicateCheckerType, uid int64, identification string, remark string, expiration time.Duration) {
	c.cache.Set(c.getCacheKey(checkerType, uid, identification), remark, expiration)
}


func (c *InMemoryDuplicateChecker) RemoveSubmissionRemark(checkerType DuplicateCheckerType, uid int64, identification string) {
	c.cache.Delete(c.getCacheKey(checkerType, uid, identification))
}


func (c *InMemoryDuplicateChecker) GetOrSetCronJobRunningInfo(jobName string, runningInfo string, runningInterval time.Duration) (bool, string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	existedRunningInfo, found := c.cache.Get(c.getCacheKey(DUPLICATE_CHECKER_TYPE_BACKGROUND_CRON_JOB, 0, jobName))

	if found {
		return true, existedRunningInfo.(string)
	}

	expiration := runningInterval

	if expiration > 1*time.Second {
		expiration = expiration - 1*time.Second
	}

	c.cache.Set(c.getCacheKey(DUPLICATE_CHECKER_TYPE_BACKGROUND_CRON_JOB, 0, jobName), runningInfo, expiration)

	return false, ""
}


func (c *InMemoryDuplicateChecker) RemoveCronJobRunningInfo(jobName string) {
	c.cache.Delete(c.getCacheKey(DUPLICATE_CHECKER_TYPE_BACKGROUND_CRON_JOB, 0, jobName))
}


func (c *InMemoryDuplicateChecker) GetFailureCount(failureKey string) uint32 {
	existedFailureCount, found := c.cache.Get(c.getCacheKey(DUPLICATE_CHECKER_TYPE_FAILURE_CHECK, 0, failureKey))

	if found {
		return existedFailureCount.(uint32)
	}

	return 0
}


func (c *InMemoryDuplicateChecker) IncreaseFailureCount(failureKey string) uint32 {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	cacheKey := c.getCacheKey(DUPLICATE_CHECKER_TYPE_FAILURE_CHECK, 0, failureKey)
	_, found := c.cache.Get(cacheKey)

	if found {
		failureCount, _ := c.cache.IncrementUint32(cacheKey, uint32(1))
		return failureCount
	} else {
		c.cache.Set(cacheKey, uint32(1), 1*time.Minute)
		return 1
	}
}

func (c *InMemoryDuplicateChecker) getCacheKey(checkerType DuplicateCheckerType, uid int64, identification string) string {
	return fmt.Sprintf("%d|%d|%s", checkerType, uid, identification)
}
