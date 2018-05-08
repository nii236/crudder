package main

// Collections need to implement these

type Collection interface {
	ListQuery() string
	ReferenceQuery(table string, column string) string
	GetManyQuery() string
	UpdateManyQuery() string
	DeleteManyQuery() string
}
