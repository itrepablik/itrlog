// Package itrlog is the custom logger for Go using Zap and Lumberjack libraries.
package itrlog

import (
	"time"

	"github.com/itrepablik/lumberjack"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// LogTimeFormat formats the event timestamp.
const LogTimeFormat string = "Jan-02-2006 03:04:05 PM"

// InitLog initialize the zap and lumberjack logger library.
// Example log file: logs/app_name_2020-02-28.log
func InitLog(maxSizeInMB, maxAgeInDays int, logFolderName, logInitial string) *zap.SugaredLogger {
	logFile := logFolderName + "/" + logInitial + time.Now().Format("2006-01-02") + ".log"

	// lumberjack.Logger is already safe for concurrent use, so we don't need to lock it.
	w := zapcore.AddSync(&lumberjack.Logger{
		Filename:  logFile,      // log filename
		MaxSize:   maxSizeInMB,  // megabytes
		MaxAge:    maxAgeInDays, // days
		LocalTime: true,         // use the local machine's timezone
	})
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		w,
		zap.InfoLevel,
	)

	logger := zap.New(core)
	sugar := logger.Sugar()
	return sugar
}
