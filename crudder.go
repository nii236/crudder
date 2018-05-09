package main

// Models need to handle these

type Item interface {
	CreateParams() []interface{}
	CreateQuery() string

	ReadParams() []interface{}
	ReadQuery() string

	UpdateManyParams() []interface{}
	UpdateParams() []interface{}
	UpdateQuery() string

	DeleteParams() []interface{}
	DeleteQuery() string
}

// DB needs to implement these
type Crudder interface {
	Create(item Item) error
	Read(item Item, id string) error
	Update(item Item, id string) error
	Delete(item Item, id string) error
}
