package cron

import (
	"fmt"
	"github.com/go-co-op/gocron/v2"
	"strings"

	"github.com/Shavitjnr/split-chill-ai/pkg/core"
	"github.com/Shavitjnr/split-chill-ai/pkg/log"
)


type GocronLoggerAdapter struct {
}


func (logger GocronLoggerAdapter) Debug(msg string, args ...any) {
	log.Debugf(core.NewNullContext(), "%s", logger.getFinalLog(msg, args...))
}


func (logger GocronLoggerAdapter) Info(msg string, args ...any) {
	log.Infof(core.NewNullContext(), "%s", logger.getFinalLog(msg, args...))
}


func (logger GocronLoggerAdapter) Warn(msg string, args ...any) {
	log.Warnf(core.NewNullContext(), "%s", logger.getFinalLog(msg, args...))
}


func (logger GocronLoggerAdapter) Error(msg string, args ...any) {
	log.Errorf(core.NewNullContext(), "%s", logger.getFinalLog(msg, args...))
}

func (logger GocronLoggerAdapter) getFinalLog(msg string, args ...any) string {
	var ret strings.Builder
	ret.WriteString(msg)

	for i := 0; i < len(args); i++ {
		if ret.Len() > 0 {
			ret.WriteRune(' ')
		}

		ret.WriteString(fmt.Sprint(args[i]))
	}

	return ret.String()
}


func NewGocronLoggerAdapter() gocron.Logger {
	return GocronLoggerAdapter{}
}
