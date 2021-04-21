package pgconf

import (
	"fmt"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Postgres struct {
	Addr     string
	Port     int
	User     string
	Password string
	DbName   string
	sslMode  bool
}

// Conn 连接数据库
func (pg *Postgres) Conn() *gorm.DB {

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s TimeZone=Asia/Shanghai",
		pg.Addr,
		pg.User,
		pg.Password,
		pg.DbName,
		pg.Port,
		pg.sslmode(),
	)
	pgDial := postgres.Open(dsn)
	db, err := gorm.Open(pgDial, &gorm.Config{})

	if err != nil {
		logrus.Errorf("Open Database failed: %v", err)
	}
	return db
}

func (pg *Postgres) sslmode() string {
	if pg.sslMode {
		return "enable"
	}
	return "disable"
}
