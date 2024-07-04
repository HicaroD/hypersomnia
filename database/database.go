package database

import (
	"database/sql"
	"fmt"
	models "github.com/HicaroD/hypersomnia/models"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var SCHEMA string = `
CREATE TABLE IF NOT EXISTS endpoint (
  id INTEGER PRIMARY KEY AUTOINCREMENT,

  method VARCHAR(10) NOT NULL,
  url VARCHAR(255) NOT NULL,
  query_params TEXT,
  headers TEXT,
  request_body_type VARCHAR(50),
  request_body TEXT,

  response_body_type VARCHAR(50),
  response_body TEXT,
  status_code INTEGER
)
`

type Database struct {
	logFile *os.File
	conn    *sql.DB
}

func New(dbPath string, logFile *os.File) (*Database, error) {
	conn, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, fmt.Errorf("unable to open SQLite3 database: %s\n", err)
	}
	if _, err := conn.Exec(SCHEMA); err != nil {
		return nil, err
	}
	return &Database{conn: conn, logFile: logFile}, nil
}

func (db *Database) ListEndpoints() ([]*models.Endpoint, error) {
	rows, err := db.conn.Query("SELECT * FROM endpoint")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var endpoints []*models.Endpoint

	for rows.Next() {
		endpoint, err := db.getEndpointFromQueryRow(rows)
		if err != nil {
			return nil, err
		}
		endpoints = append(endpoints, endpoint)
	}

	return endpoints, nil
}

func (db *Database) getEndpointFromQueryRow(rows *sql.Rows) (*models.Endpoint, error) {
	var endpoint models.Endpoint
	if err := rows.Scan(&endpoint.Id, &endpoint.Method, &endpoint.Url,
		&endpoint.RequestQueryParams, &endpoint.RequestHeaders, &endpoint.RequestBodyType,
		&endpoint.RequestBody, &endpoint.ResponseBodyType, &endpoint.ResponseBody,
		&endpoint.StatusCode); err != nil {
		return nil, err
	}
	return &endpoint, nil
}

func (db *Database) Close() error {
	return db.conn.Close()
}
