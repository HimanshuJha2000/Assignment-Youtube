package main

import (
	"database/sql"
	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(upVideoDataModels, downVideoDataModels)
}

func upVideoDataModels(tx *sql.Tx) error {
	_, err := tx.Exec(`CREATE TABLE video_data_models (
 		title  varchar(200) NOT NULL UNIQUE,
  		description varchar(200) NOT NULL UNIQUE,
  		published_at bigint NOT NULL,
		channel_title varchar(200) NOT NULL,
		thumbnail_url varchar(200),
		PRIMARY KEY(Title, Description)
    );
	CREATE INDEX video_data_models_pkey ON video_data_models USING btree (title, description);`)
	if err != nil {
		return err
	}
	return nil
}

func downVideoDataModels(tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	_, err := tx.Exec(`DROP TABLE IF EXISTS video_data_models;`)
	if err != nil {
		return err
	}
	return nil
}
