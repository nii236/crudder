package main

import (
	"fmt"
	"regexp"
)

// UserList is a list of users
type UserList []*User

// ListQuery implements the Collection interface
func (u *UserList) ListQuery() string {
	return "SELECT id, name, archived FROM users WHERE archived = false"
}

// ReferenceQuery implements the Collection interface
func (u *UserList) ReferenceQuery(table string, column string) string {
	isAlpha := regexp.MustCompile(`^[A-Za-z]+$`).MatchString
	if !isAlpha(table) {
		panic("non alphanumeric table name provided")
	}

	if !isAlpha(column) {
		panic("non alphanumeric column name provided: " + column)
	}

	return fmt.Sprintf(`
SELECT 
	users.id, 
	users.name, 
	users.archived 
FROM users 
INNER JOIN "%s" on "%s".id = users."%s"
WHERE users.archived = false
AND "%s".id = $1
`, table, table, column, table)

}

// GetManyQuery implements the Collection interface
func (u *UserList) GetManyQuery() string {
	return `
SELECT id, name, archived 
FROM users 
WHERE archived = false
AND id = ANY($1)
`
}

// UpdateManyQuery implements the Collection interface
func (u *UserList) UpdateManyQuery() string {
	return `
UPDATE users 
SET name = $1
WHERE archived = false 
AND id = ANY($2)
RETURNING id, name, archived
`
}

// DeleteManyQuery implements the Collection interface
func (u *UserList) DeleteManyQuery() string {
	return `
UPDATE users 
SET archived = true
WHERE id = ANY($1)
RETURNING id, name, archived
`
}
