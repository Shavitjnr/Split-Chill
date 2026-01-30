package duplicatechecker

import (
	"time"

	"github.com/Shavitjnr/split-chill-ai/pkg/errs"
	"github.com/Shavitjnr/split-chill-ai/pkg/settings"
)


type DuplicateCheckerContainer struct {
	current DuplicateChecker
}


var (
	Container = &DuplicateCheckerContainer{}
)


func InitializeDuplicateChecker(config *settings.Config) error {
	if config.DuplicateCheckerType == settings.InMemoryDuplicateCheckerType {
		checker, err := NewInMemoryDuplicateChecker(config)
		Container.current = checker

		return err
	}

	return errs.ErrInvalidDuplicateCheckerType
}


func SetDuplicateChecker(checker DuplicateChecker) {
	Container.current = checker
}


func (c *DuplicateCheckerContainer) IsEnabled() bool {
	return c.current != nil
}


func (c *DuplicateCheckerContainer) GetSubmissionRemark(checkerType DuplicateCheckerType, uid int64, identification string) (bool, string) {
	if c.current == nil {
		return false, ""
	}

	return c.current.GetSubmissionRemark(checkerType, uid, identification)
}


func (c *DuplicateCheckerContainer) SetSubmissionRemark(checkerType DuplicateCheckerType, uid int64, identification string, remark string) {
	if c.current == nil {
		return
	}

	c.current.SetSubmissionRemark(checkerType, uid, identification, remark)
}


func (c *DuplicateCheckerContainer) SetSubmissionRemarkWithCustomExpiration(checkerType DuplicateCheckerType, uid int64, identification string, remark string, expiration time.Duration) {
	if c.current == nil {
		return
	}

	c.current.SetSubmissionRemarkWithCustomExpiration(checkerType, uid, identification, remark, expiration)
}


func (c *DuplicateCheckerContainer) RemoveSubmissionRemark(checkerType DuplicateCheckerType, uid int64, identification string) {
	if c.current == nil {
		return
	}

	c.current.RemoveSubmissionRemark(checkerType, uid, identification)
}


func (c *DuplicateCheckerContainer) GetOrSetCronJobRunningInfo(jobName string, runningInfo string, runningInterval time.Duration) (bool, string) {
	if c.current == nil {
		return false, ""
	}

	return c.current.GetOrSetCronJobRunningInfo(jobName, runningInfo, runningInterval)
}


func (c *DuplicateCheckerContainer) RemoveCronJobRunningInfo(jobName string) {
	if c.current == nil {
		return
	}

	c.current.RemoveCronJobRunningInfo(jobName)
}


func (c *DuplicateCheckerContainer) GetFailureCount(failureKey string) uint32 {
	if c.current == nil {
		return 0
	}

	return c.current.GetFailureCount(failureKey)
}


func (c *DuplicateCheckerContainer) IncreaseFailureCount(failureKey string) uint32 {
	if c.current == nil {
		return 0
	}

	return c.current.IncreaseFailureCount(failureKey)
}
