package postgreservice

import (
	"fmt"
	"log"

	"github.com/jackc/pgx/v4/pgxpool"
)

type PostgresConfig struct {
	Host        string `json:"host" validate:"required"`
	Port        int    `json:"port" validate:"required,numeric"`
	Database    string `json:"database" validate:"required"`
	User        string `json:"user" validate:"required"`
	Password    string `json:"password" validate:"required"`
	DSN         string `json:"dsn"`
	ConnMaxOpen int32  `json:"connmaxopen"`
}

func GetPostgresConfig(config *PostgresConfig) *pgxpool.Config {

	var maxconn int32 = 10
	log.Println("config: ", *config)
	if config.ConnMaxOpen != 0 {
		maxconn = config.ConnMaxOpen
	}
	//connString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=verify-ca&pool_max_conns=%s", config.User, config.Password, config.Host, config.Port, config.Database, maxconn)
	connString := fmt.Sprintf("user=%s password=%s host=%s port=%v dbname=%s sslmode=disable", config.User, config.Password, config.Host, config.Port, config.Database)
	pgxConfig, err := pgxpool.ParseConfig(connString)
	if err != nil {
		log.Println("error in parsing postgreconf", err)
		return nil
	}
	log.Println("pgxConfig: ", *pgxConfig)
	pgxConfig.MaxConns = maxconn
	log.Println(pgxConfig, err)
	if err != nil {
		return nil
	}

	return pgxConfig
}
