package main

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type DB struct {
	conn *sqlx.DB
}

// List implements the Collectioner interface
func (db *DB) List(collection Collection) error {
	return db.conn.Select(collection, collection.ListQuery())
}

// UpdateMany implements the Collectioner interface
func (db *DB) UpdateMany(collection Collection, item Item, IDs ...interface{}) error {
	params := item.UpdateParams()
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
func (db *DB) GetManyReference(collection Collection, table string, PKColumn string, PK interface{}) error {
	fmt.Println(collection.ReferenceQuery(table, PKColumn))
	return db.conn.Select(collection, collection.ReferenceQuery(table, PKColumn), PK)
}

// Create implements the Crudder interface
func (db *DB) Create(item Item, fkID interface{}) error {
	params := item.CreateParams()
	args := []interface{}{}
	if fkID != nil {
		args = append(args, fkID)
	}
	for _, v := range params {
		args = append(args, v)
	}
	return db.conn.Get(item, item.CreateQuery(), args...)
}

// Read implements the Crudder interface
func (db *DB) Read(item Item) error {
	return db.conn.Get(item, item.ReadQuery(), item.ReadParams()...)
}

// Update implements the Crudder interface
func (db *DB) Update(item Item, ID interface{}) error {
	params := item.UpdateParams()
	args := []interface{}{ID}
	for _, v := range params {
		args = append(args, v)
	}
	return db.conn.Get(item, item.UpdateQuery(), args...)
}

// Delete implements the Crudder interface
func (db *DB) Delete(item Item) error {
	return db.conn.Get(item, item.DeleteQuery(), item.DeleteParams()...)
}
