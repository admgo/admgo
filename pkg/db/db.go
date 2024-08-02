package db

import (
	"context"
	"fmt"
	"github.com/admgo/admgo/pkg/db/config"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DB struct {
	*gorm.DB
}

type ormLog struct {
	LogLevel logger.LogLevel
}

func (l *ormLog) LogMode(level logger.LogLevel) logger.Interface {
	l.LogLevel = level
	return l
}

func (l *ormLog) Info(ctx context.Context, format string, v ...interface{}) {
	if l.LogLevel < logger.Info {
		return
	}
	logx.WithContext(ctx).Infof(format, v...)
}

func (l *ormLog) Warn(ctx context.Context, fromat string, v ...interface{}) {
	if l.LogLevel < logger.Warn {
		return
	}
	logx.WithContext(ctx).Infof(fromat, v...)
}

func (l *ormLog) Error(ctx context.Context, format string, v ...interface{}) {
	if l.LogLevel < logger.Error {
		return
	}
	logx.WithContext(ctx).Errorf(format, v...)
}

func (l *ormLog) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	elapsed := time.Since(begin)
	sql, rows := fc()
	logx.WithContext(ctx).WithDuration(elapsed).Infof("[%.3fms] [rows:%v] %s", float64(elapsed.Nanoseconds())/1e6, rows, sql)
}

func NewMysql(conf *config.Config) (*DB, error) {
	if conf.MaxIdleConns == 0 {
		conf.MaxIdleConns = 10
	}
	if conf.MaxOpenConns == 0 {
		conf.MaxOpenConns = 100
	}
	if conf.MaxLifetime == 0 {
		conf.MaxLifetime = 3600
	}
	timeout := config.DBTimeout
	if conf.TimeOut > 0 {
		timeout = conf.TimeOut
	}

	writeTimeout := config.DBWriteTimeout
	if conf.WriteTimeOut > 0 {
		writeTimeout = conf.WriteTimeOut
	}

	readTimeout := config.DBReadTimeout
	if conf.ReadTimeOut > 0 {
		readTimeout = conf.ReadTimeOut
	}

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local&timeout=%dms&writeTimeout=%dms&readTimeout=%dms",
		conf.Username,
		conf.Password,
		conf.Host,
		conf.Port,
		conf.Database,
		conf.Charset,
		timeout,
		writeTimeout,
		readTimeout,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: &ormLog{},
	})
	if err != nil {
		return nil, err
	}
	sdb, err := db.DB()
	if err != nil {
		return nil, err
	}
	sdb.SetMaxIdleConns(conf.MaxIdleConns)
	sdb.SetMaxOpenConns(conf.MaxOpenConns)
	sdb.SetConnMaxLifetime(time.Second * time.Duration(conf.MaxLifetime))

	err = db.Use(NewCustomePlugin())
	if err != nil {
		return nil, err
	}

	return &DB{DB: db}, nil
}

func MustNewMysql(conf *config.Config) *DB {
	db, err := NewMysql(conf)
	if err != nil {
		panic(err)
	}

	return db
}
