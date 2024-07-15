package database

import (
	_ "embed"

	"database/sql"
	"fmt"
	models "github.com/HicaroD/hypersomnia/models"

	_ "github.com/mattn/go-sqlite3"
)

//go:embed schema.sql
var SQL_SCHEMA string

type Database struct {
	conn *sql.DB
}

func New(dbPath string) (*Database, error) {
	conn, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, fmt.Errorf("unable to open SQLite3 database: %s\n", err)
	}
	if _, err := conn.Exec(SQL_SCHEMA); err != nil {
		return nil, err
	}
	return &Database{conn: conn}, nil
}

func (db *Database) AddNewCollection(name string) error {
	// TODO: add new collection to database
	// See https://go.dev/doc/database/sql-injection for avoiding SQL injection
	return nil
}

func (db *Database) ListCollections() ([]*models.Collection, error) {
	rows, err := db.conn.Query("SELECT * FROM collection;")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var collections []*models.Collection
	for rows.Next() {
		collection, err := db.getCollectionFromQueryRow(rows)
		if err != nil {
			return nil, err
		}
		collections = append(collections, collection)
	}
	return collections, nil
}

func (db *Database) ListEndpoints() ([]*models.Endpoint, error) {
	rows, err := db.conn.Query("SELECT * FROM endpoint;")
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

func (db *Database) getCollectionFromQueryRow(rows *sql.Rows) (*models.Collection, error) {
	var collection models.Collection
	if err := rows.Scan(&collection.Id, &collection.Name); err != nil {
		return nil, err
	}
	return &collection, nil
}

func (db *Database) Close() error {
	return db.conn.Close()
}
