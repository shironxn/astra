package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/shironxn/astra/internal/config"
)

type Databse struct {
	User string
	Pass string
	Host string
	Port string
	Name string
}

func NewDatabase(conf config.DatabaseConfig) *Databse {
	return &Databse{
		User: conf.User,
		Pass: conf.Pass,
		Host: conf.Host,
		Port: conf.Port,
		Name: conf.Name,
	}
}

func (d *Databse) Connection() (*sql.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		d.User,
		d.Pass,
		d.Host,
		d.Port,
		d.Name,
	)
	return sql.Open("mysql", dsn)
}
