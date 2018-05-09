package main

type Team struct {
	ID       int
	Name     string
	UserID   int
	Archived bool
}

// CreateParams implements the Cruddable interface
func (u *Team) CreateParams() []interface{} {
	return []interface{}{u.UserID, u.Name}
}

// CreateQuery implements the Cruddable interface
func (u *Team) CreateQuery() string {
	return "INSERT INTO teams (user_id, name) VALUES ($1, $2) RETURNING id, name, archived"
}

// ReadParams implements the Cruddable interface
func (u *Team) ReadParams() []interface{} {
	return []interface{}{}
}

// ReadQuery implements the Cruddable interface
func (u *Team) ReadQuery() string {
	return "SELECT id, name, archived FROM teams WHERE id=$1 AND archived = false"
}

// UpdateManyParams implements the Cruddable interface
func (u *Team) UpdateManyParams() []interface{} {
	return []interface{}{u.Name}
}

// UpdateParams implements the Cruddable interface
func (u *Team) UpdateParams() []interface{} {
	return []interface{}{u.Name}
}

// UpdateQuery implements the Cruddable interface
func (u *Team) UpdateQuery() string {
	return "UPDATE teams SET name = $1 WHERE id=$2 AND archived = false RETURNING id, name, archived"
}

// DeleteParams implements the Cruddable interface
func (u *Team) DeleteParams() []interface{} {
	return []interface{}{}
}

// DeleteQuery implements the Cruddable interface
func (u *Team) DeleteQuery() string {
	return "UPDATE teams SET archived=true WHERE id=$1 RETURNING id, name, archived"
}
