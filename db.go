package main

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

// DB is the conn
type DB struct {
	conn *sqlx.DB
}

// List implements the Collectioner interface
func (db *DB) List(collection Collection) error {
	return db.conn.Select(collection, collection.ListQuery())
}

// UpdateMany implements the Collectioner interface
func (db *DB) UpdateMany(collection Collection, item Item, IDs ...interface{}) error {
	params := item.UpdateManyParams()
	args := []interface{}{pq.Array(IDs)}
	for _, v := range params {
		args = append(args, v)
	}
	return db.conn.Select(collection, collection.UpdateManyQuery(), args...)
}

// DeleteMany implements the Collectioner interface
func (db *DB) DeleteMany(collection Collection, IDs ...interface{}) error {
	return db.conn.Select(collection, collection.DeleteManyQuery(), pq.Array(IDs))
}

// GetMany implements the Collectioner interface
func (db *DB) GetMany(collection Collection, IDs ...interface{}) error {
	return db.conn.Select(collection, collection.GetManyQuery(), pq.Array(IDs))
}

// GetManyReference implements the Collectioner interface
func (db *DB) GetManyReference(collection Collection, table string, FKColumn string, PK interface{}) error {
	fmt.Println(collection.ReferenceQuery(table, FKColumn))
	return db.conn.Select(collection, collection.ReferenceQuery(table, FKColumn), PK)
}

// Create implements the Crudder interface
func (db *DB) Create(item Item) error {
	return db.conn.Get(item, item.CreateQuery(), item.CreateParams()...)
}

// Read implements the Crudder interface
func (db *DB) Read(item Item) error {
	return db.conn.Get(item, item.ReadQuery(), item.ReadParams()...)
}

// Update implements the Crudder interface
func (db *DB) Update(item Item) error {
	return db.conn.Get(item, item.UpdateQuery(), item.UpdateParams()...)
}

// Delete implements the Crudder interface
func (db *DB) Delete(item Item) error {
	return db.conn.Get(item, item.DeleteQuery(), item.DeleteParams())
}
