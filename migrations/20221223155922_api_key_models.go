package main

import (
	"database/sql"
	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(upApiKeyModels, downApiKeyModels)
}

func upApiKeyModels(tx *sql.Tx) error {
	_, err := tx.Exec(`CREATE TABLE api_key_models (
		api_key varchar(50) NOT NULL UNIQUE,
		is_valid int NOT NULL DEFAULT 1,
		usage_count bigint NOT NULL DEFAULT 0,
		PRIMARY KEY (api_key)
    );`)
	if err != nil {
		return err
	}
	return nil
}

func downApiKeyModels(tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	_, err := tx.Exec(`DROP TABLE IF EXISTS api_key_models;`)
	if err != nil {
		return err
	}
	return nil
}
