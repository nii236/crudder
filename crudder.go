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
	Create(target Item, source Item) error
	Read(target Item, id string) error
	Update(target Item, source Item, id string) error
	Delete(target Item, id string) error
}
