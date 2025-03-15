package bootstrap

import (
	"assignment/infrastructures/sql/database"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log/slog"
	"strings"
	"time"
)

type Database struct {
	Query *database.Queries
}

type mysqlDataSource struct {
	username string
	password string
	host     string
	port     string
	socket   string
	dbname   string
	options  options
}

type options map[string]string

func (o options) toString() string {
	if len(o) == 0 {
		return ""
	}
	var values []string
	for k, v := range o {
		values = append(values, fmt.Sprintf("%s=%s", k, v))
	}
	return "?" + strings.Join(values, "&")
}

func (ds *mysqlDataSource) address() string {
	switch {
	case ds.socket != "":
		return fmt.Sprintf("unix(%s)", ds.socket)
	default:
		return fmt.Sprintf("tcp(%s:%s)", ds.hostOrDefault(), ds.portOrDefault())
	}
}

func (ds *mysqlDataSource) hostOrDefault() string {
	if ds.host == "" {
		return "localhost"
	}
	return ds.host
}

func (ds *mysqlDataSource) portOrDefault() string {
	if ds.port == "" {
		return "3306"
	}
	return ds.port
}

func NewMySQLDatabase(env *Env) *Database {
	dbUser := env.DBUser
	dbPassword := env.DBPassword
	dbHost := env.DBHost
	dbPort := env.DBPort
	dbDatabase := env.DBDatabase
	opts := options{}

	dataSource := &mysqlDataSource{
		username: dbUser,
		password: dbPassword,
		host:     dbHost,
		port:     dbPort,
		dbname:   dbDatabase,
		options:  opts,
	}

	dbUrl := fmt.Sprintf("%s:%s@%s/%s%s?parseTime=true", dataSource.username, dataSource.password, dataSource.address(), dataSource.dbname, dataSource.options.toString())
	db, err := sql.Open("mysql", dbUrl)
	if err != nil {
		slog.Error("Failed to connect to database", slog.String("reason", err.Error()))
		return nil
	}
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)
	err = db.Ping()
	if err != nil {
		slog.Error("Failed to ping database", slog.String("reason", err.Error()))
		return nil
	}

	return &Database{
		Query: database.New(db),
	}
}
