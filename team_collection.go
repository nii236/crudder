package main

import (
	"fmt"
	"regexp"
)

// TeamList is a list of teams
type TeamList []*Team

// ListQuery implements the Collection interface
func (u *TeamList) ListQuery() string {
	return "SELECT id, name, archived FROM teams WHERE archived = false"
}

// ReferenceQuery implements the Collection interface
func (u *TeamList) ReferenceQuery(table string, column string) string {
	isAlpha := regexp.MustCompile(`^[A-Za-z_]+$`).MatchString
	if !isAlpha(table) {
		panic("non alphanumeric table name provided")
	}

	if !isAlpha(column) {
		panic("non alphanumeric column name provided: " + column)
	}

	return fmt.Sprintf(`
SELECT 
	teams.id, 
	teams.name, 
	teams.archived 
FROM teams 
INNER JOIN "%s" on "%s".id = teams."%s"
WHERE teams.archived = false
AND "%s".id = $1
`, table, table, column, table)

}

// GetManyQuery implements the Collection interface
func (u *TeamList) GetManyQuery() string {
	return `
SELECT id, name, archived 
FROM teams 
WHERE archived = false
AND id = ANY($1)
`
}

// UpdateManyQuery implements the Collection interface
func (u *TeamList) UpdateManyQuery() string {
	return `
UPDATE teams 
SET name = $2 
WHERE archived = false 
AND id = ANY($1)
RETURNING id, name, archived
`
}

// DeleteManyQuery implements the Collection interface
func (u *TeamList) DeleteManyQuery() string {
	return `
UPDATE teams 
SET archived = true
WHERE id = ANY($1)
RETURNING id, name, archived
`
}
