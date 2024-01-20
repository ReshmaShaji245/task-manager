package postgreservice

import (
	"context"
	"sync"

	"github.com/jackc/pgx/v4"
)

type PostgresDBService struct {
	Config     *PostgresConfig
	dbinstance *pgx.Conn
	lock       *sync.Mutex
}

func (dbs *PostgresDBService) createNewDBInstance() (*pgx.Conn, error) {
	postgresConfig := GetPostgresConfig(dbs.Config)

	dbinst, err := pgx.ConnectConfig(context.Background(), postgresConfig.ConnConfig)
	if err != nil {
		return nil, err
	}

	err = dbinst.Ping(context.Background())
	if err != nil {
		dbinst.Close(context.Background())
		return nil, err
	}

	return dbinst, nil
}

func (dbs *PostgresDBService) GetDBInstance() (*pgx.Conn, error) {
	dbs.lock.Lock()
	defer dbs.lock.Unlock()

	if dbs.dbinstance != nil {
		return dbs.dbinstance, nil
	}

	dbinst, err := dbs.createNewDBInstance()
	if err != nil {
		return nil, err
	}
	dbs.dbinstance = dbinst
	return dbs.dbinstance, nil
}

func (dbs *PostgresDBService) ClearDBInstance() (bool, error) {
	dbs.lock.Lock()
	defer dbs.lock.Unlock()

	if dbs.dbinstance == nil {
		return true, nil
	}

	dbs.dbinstance.Close(context.Background())
	dbs.dbinstance = nil
	return true, nil
}

func NewPostgresDBService(postgresConfig *PostgresConfig) *PostgresDBService {
	return &PostgresDBService{
		Config:     postgresConfig,
		dbinstance: nil,
		lock:       &sync.Mutex{},
	}
}
