// Package itrlog is the custom logger for Go using Zap and Lumberjack libraries.
package itrlog

import (
	"time"

	"github.com/itrepablik/lumberjack"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// ITRLogger contains the necessary information to initialize the required parameters for the itrlog package.
type ITRLogger struct {
	MaxSizeInMB, MaxAgeInDays int
	LogFolderName, LogInitial string
	Sugar                     *zap.SugaredLogger
}

// LogTimeFormat formats the event timestamp.
const LogTimeFormat string = "Jan-02-2006 03:04:05 PM"

// InitLog initialize the zap and lumberjack logger library.
// Example log file: logs/app_name_2020-02-28.log
func InitLog(maxSizeInMB, maxAgeInDays int, logFolderName, logInitial string) *ITRLogger {
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
	return &ITRLogger{MaxSizeInMB: maxSizeInMB, MaxAgeInDays: maxAgeInDays, LogFolderName: logFolderName, LogInitial: logInitial, Sugar: sugar}
}

// Debug uses fmt.Sprint to construct and log a message.
func (s *ITRLogger) Debug(args ...interface{}) {
	s.Sugar.Debug(args)
}

// Info uses fmt.Sprint to construct and log a message.
func (s *ITRLogger) Info(args ...interface{}) {
	s.Sugar.Info(args)
}

// Warn uses fmt.Sprint to construct and log a message.
func (s *ITRLogger) Warn(args ...interface{}) {
	s.Sugar.Warn(args)
}

// Error uses fmt.Sprint to construct and log a message.
func (s *ITRLogger) Error(args ...interface{}) {
	s.Sugar.Error(args)
}

// DPanic uses fmt.Sprint to construct and log a message. In development, the
// logger then panics. (See DPanicLevel for details.)
func (s *ITRLogger) DPanic(args ...interface{}) {
	s.Sugar.DPanic(args)
}

// Panic uses fmt.Sprint to construct and log a message, then panics.
func (s *ITRLogger) Panic(args ...interface{}) {
	s.Sugar.Panic(args)
}

// Fatal uses fmt.Sprint to construct and log a message, then calls os.Exit.
func (s *ITRLogger) Fatal(args ...interface{}) {
	s.Sugar.Fatal(args)
}

// Debugf uses fmt.Sprintf to log a templated message.
func (s *ITRLogger) Debugf(template string, args ...interface{}) {
	s.Sugar.Debugf(template, args)
}

// Infof uses fmt.Sprintf to log a templated message.
func (s *ITRLogger) Infof(template string, args ...interface{}) {
	s.Sugar.Infof(template, args)
}

// Warnf uses fmt.Sprintf to log a templated message.
func (s *ITRLogger) Warnf(template string, args ...interface{}) {
	s.Sugar.Infof(template, args)
}

// Errorf uses fmt.Sprintf to log a templated message.
func (s *ITRLogger) Errorf(template string, args ...interface{}) {
	s.Sugar.Errorf(template, args)
}

// DPanicf uses fmt.Sprintf to log a templated message. In development, the
// logger then panics. (See DPanicLevel for details.)
func (s *ITRLogger) DPanicf(template string, args ...interface{}) {
	s.Sugar.DPanicf(template, args)
}

// Panicf uses fmt.Sprintf to log a templated message, then panics.
func (s *ITRLogger) Panicf(template string, args ...interface{}) {
	s.Sugar.Panicf(template, args)
}

// Fatalf uses fmt.Sprintf to log a templated message, then calls os.Exit.
func (s *ITRLogger) Fatalf(template string, args ...interface{}) {
	s.Sugar.Fatalf(template, args)
}

// Debugw logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
//
// When debug-level logging is disabled, this is much faster than
//  s.With(keysAndValues).Debug(msg)
func (s *ITRLogger) Debugw(msg string, keysAndValues ...interface{}) {
	s.Sugar.Debugw(msg, keysAndValues)
}

// Infow logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
func (s *ITRLogger) Infow(msg string, keysAndValues ...interface{}) {
	s.Sugar.Infow(msg, keysAndValues)
}

// Warnw logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
func (s *ITRLogger) Warnw(msg string, keysAndValues ...interface{}) {
	s.Sugar.Warnw(msg, keysAndValues)
}

// Errorw logs a message with some additional context. The variadic key-value
// pairs are treated as they are in With.
func (s *ITRLogger) Errorw(msg string, keysAndValues ...interface{}) {
	s.Sugar.Errorw(msg, keysAndValues)
}

// DPanicw logs a message with some additional context. In development, the
// logger then panics. (See DPanicLevel for details.) The variadic key-value
// pairs are treated as they are in With.
func (s *ITRLogger) DPanicw(msg string, keysAndValues ...interface{}) {
	s.Sugar.DPanicw(msg, keysAndValues)
}

// Panicw logs a message with some additional context, then panics. The
// variadic key-value pairs are treated as they are in With.
func (s *ITRLogger) Panicw(msg string, keysAndValues ...interface{}) {
	s.Sugar.Panicw(msg, keysAndValues)
}

// Fatalw logs a message with some additional context, then calls os.Exit. The
// variadic key-value pairs are treated as they are in With.
func (s *ITRLogger) Fatalw(msg string, keysAndValues ...interface{}) {
	s.Sugar.Fatalw(msg, keysAndValues)
}
