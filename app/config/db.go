package config

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

func DBConnection(cfgDB DB) (*sqlx.DB, error) {
	dsn := buildDSN(cfgDB)
	db, err := sqlx.Open(cfgDB.Driver, dsn)
	if err != nil {
		panic(err)
	}
	db.SetMaxOpenConns(32)
	db.SetMaxIdleConns(32)

	if err := db.Ping(); err != nil {
		panic(err)
	}

	if err := initTables(db); err != nil {
		return db, err
	}
	if err := prepareData(db); err != nil {
		return db, err
	}
	return db, nil
}

func initTables(db *sqlx.DB) error {
	tx, txErr := db.Begin()
	if txErr != nil {
		return txErr
	}
	tx.Exec()

	return nil
}

func prepareData(db *sqlx.DB) error {
	return nil
}

func buildDSN(cfg DB) string {
	sslMode := "disable"
	if cfg.SSLMode {
		sslMode = "require"
	}

	return fmt.Sprintf("%s://%s:%s@%s:%d?sslmode=%s",
		cfg.Driver,
		cfg.Login,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		sslMode,
	)
}
