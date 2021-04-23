package config

import (
	"github.com/tangx/dnsx/pkg/pgconf"
	"gorm.io/gorm"
)

var DbConn *gorm.DB
var dbConfig = &pgconf.Postgres{
	User:     "postgres",
	Password: "Abc123123",
	DbName:   "dnsx",
}

func init() {
	DbConn = dbConfig.Conn()
}
