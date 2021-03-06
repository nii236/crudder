package main

import (
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

var _ Collectioner = &DB{}
var _ Crudder = &DB{}

// DB is the conn
type DB struct {
	conn *sqlx.DB
}

// List implements the Collectioner interface
func (db *DB) List(collection Collection) error {
	return db.conn.Select(collection, collection.ListQuery())
}

// UpdateMany implements the Collectioner interface
func (db *DB) UpdateMany(collection Collection, source Item, IDs []string) error {
	params := source.UpdateManyParams()
	params = append(params, pq.Array(IDs))
	return db.conn.Select(collection, collection.UpdateManyQuery(), params...)
}

// DeleteMany implements the Collectioner interface
func (db *DB) DeleteMany(collection Collection, IDs []string) error {
	return db.conn.Select(collection, collection.DeleteManyQuery(), pq.Array(IDs))
}

// GetMany implements the Collectioner interface
func (db *DB) GetMany(collection Collection, IDs []string) error {
	return db.conn.Select(collection, collection.GetManyQuery(), pq.Array(IDs))
}

// Reference implements the Collectioner interface
func (db *DB) Reference(collection Collection, table string, column string, ID string) error {
	return db.conn.Select(collection, collection.ReferenceQuery(table, column), ID)
}

// Create implements the Crudder interface
func (db *DB) Create(target Item, source Item) error {
	params := source.CreateParams()
	log.Println(params)
	return db.conn.Get(target, source.CreateQuery(), params...)
}

// Read implements the Crudder interface
func (db *DB) Read(item Item, ID string) error {
	params := item.ReadParams()
	params = append(params, ID)
	return db.conn.Get(item, item.ReadQuery(), params...)
}

// Update implements the Crudder interface
func (db *DB) Update(target Item, source Item, ID string) error {
	params := target.UpdateParams()
	params = append(params, ID)
	return db.conn.Get(source, target.UpdateQuery(), params...)
}

// Delete implements the Crudder interface
func (db *DB) Delete(item Item, ID string) error {
	params := item.DeleteParams()
	params = append(params, ID)
	return db.conn.Get(item, item.DeleteQuery(), params...)
}
