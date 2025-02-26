package logger

import (
	"context"
	"errors"
	"time"

	"go.uber.org/zap"
	gormlogger "gorm.io/gorm/logger"
)

type GormLogger struct {
	ZapLogger     *zap.Logger
	SlowThreshold time.Duration
}

func NewGormLogger() *GormLogger {
	return &GormLogger{
		ZapLogger:     Logger,
		SlowThreshold: 200 * time.Millisecond,
	}
}

func (l GormLogger) LogMode(level gormlogger.LogLevel) gormlogger.Interface {
	return GormLogger{
		ZapLogger:     l.ZapLogger,
		SlowThreshold: l.SlowThreshold,
	}
}

func (l GormLogger) Error(ctx context.Context, str string, data ...interface{}) {
	l.ZapLogger.Sugar().Errorf(str, data...)
}
func (l GormLogger) Info(ctx context.Context, str string, data ...interface{}) {
	l.ZapLogger.Sugar().Debugf(str, data...)
}

func (l GormLogger) Warn(ctx context.Context, str string, data ...interface{}) {
	l.ZapLogger.Sugar().Warnf(str, data...)
}

func (l GormLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	elapsed := time.Since(begin)

	// 获取 SQL 请求和返回条数
	sql, rowsAffected := fc()

	// 通用字段
	logFields := []zap.Field{
		zap.String("sql", sql),
		zap.String("time", elapsed.String()),
		zap.Int64("rows", rowsAffected),
	}

	// Gorm 错误
	if err != nil {
		if errors.Is(err, gormlogger.ErrRecordNotFound) {
			l.ZapLogger.Warn("Database ErrRecordNotFound", logFields...)
		} else {
			logFields = append(logFields, zap.Error(err))
			l.ZapLogger.Error("Database Error", logFields...)
		}
	}

	// 慢查询日志
	if l.SlowThreshold != 0 && elapsed > l.SlowThreshold {
		l.ZapLogger.Warn("Database Slow Log", logFields...)
	}

	// 记录所有 SQL 请求
	l.ZapLogger.Debug("Database Query", logFields...)
}
