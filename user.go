package main

type User struct {
	ID       int
	Name     string
	Archived bool
}

// CreateParams implements the Cruddable interface
func (u *User) CreateParams() []interface{} {
	return []interface{}{u.Name}
}

// CreateQuery implements the Cruddable interface
func (u *User) CreateQuery() string {
	return "INSERT INTO users (name) VALUES ($1) RETURNING id, name, archived"
}

// ReadParams implements the Cruddable interface
func (u *User) ReadParams() []interface{} {
	return []interface{}{}
}

// ReadQuery implements the Cruddable interface
func (u *User) ReadQuery() string {
	return "SELECT id, name, archived FROM users WHERE id=$1 AND archived = false"
}

// UpdateManyParams implements the Cruddable interface
func (u *User) UpdateManyParams() []interface{} {
	return []interface{}{u.Name}
}

// UpdateParams implements the Cruddable interface
func (u *User) UpdateParams() []interface{} {
	return []interface{}{u.Name}
}

// UpdateQuery implements the Cruddable interface
func (u *User) UpdateQuery() string {
	return "UPDATE users SET name = $1 WHERE id=$2 AND archived = false RETURNING id, name, archived"
}

// DeleteParams implements the Cruddable interface
func (u *User) DeleteParams() []interface{} {
	return []interface{}{}
}

// DeleteQuery implements the Cruddable interface
func (u *User) DeleteQuery() string {
	return "UPDATE users SET archived=true WHERE id=$1 RETURNING id, name, archived"
}
