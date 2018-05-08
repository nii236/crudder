package main

import "github.com/jmoiron/sqlx"

type DB struct {
	conn *sqlx.DB
}

func (db *DB) List(item Listable) error {
	err := db.conn.Select(item, item.ListQuery(), item.ListParams()...)
	if err != nil {
		return err
	}

	return nil
}

func (db *DB) Create(item Createable) error {
	err := db.conn.QueryRow(item.CreateQuery(), item.CreateParams()...).Scan(item.CreateTargets()...)
	if err != nil {
		return err
	}
	return nil
}

func (db *DB) Read(item Readable) error {
	err := db.conn.QueryRow(item.ReadQuery(), item.ReadParams()...).Scan(item.ReadTargets()...)
	if err != nil {
		return err
	}
	return nil
}

func (db *DB) Update(item Updateable) error {
	err := db.conn.QueryRow(item.UpdateQuery(), item.UpdateParams()...).Scan(item.UpdateTargets()...)
	if err != nil {
		return err
	}
	return nil
}

func (db *DB) Delete(item Deleteable) error {
	err := db.conn.QueryRow(item.DeleteQuery(), item.DeleteParams()...).Scan(item.DeleteTargets()...)
	if err != nil {
		return err
	}
	return nil
}
