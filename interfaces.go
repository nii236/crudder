package main

type Listable interface {
	ListParams() []interface{}
	ListTargets() []interface{}
	ListQuery() string
}

type Createable interface {
	CreateParams() []interface{}
	CreateTargets() []interface{}
	CreateQuery() string
}

type Readable interface {
	ReadParams() []interface{}
	ReadTargets() []interface{}
	ReadQuery() string
}

type Updateable interface {
	UpdateParams() []interface{}
	UpdateTargets() []interface{}
	UpdateQuery() string
}

type Deleteable interface {
	DeleteParams() []interface{}
	DeleteTargets() []interface{}
	DeleteQuery() string
}

type Lister interface {
	List(item Listable) error
}

type Creater interface {
	Create(item Createable) error
}

type Reader interface {
	Read(item Createable) error
}

type Updater interface {
	Update(item Createable) error
}

type Deleter interface {
	Delete(item Createable) error
}
