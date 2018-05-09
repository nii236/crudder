package main

// Collections need to implement these

type Collection interface {
	ListQuery() string
	ReferenceQuery(table string, column string) string
	GetManyQuery() string
	UpdateManyQuery() string
	DeleteManyQuery() string
}

// DB needs to implement these
type Collectioner interface {
	List(collection Collection) error
	Reference(collection Collection, table string, column string, ID string) error
	GetMany(collection Collection, IDs []string) error
	UpdateMany(collection Collection, item Item, IDs []string) error
	DeleteMany(collection Collection, IDs []string) error
}
