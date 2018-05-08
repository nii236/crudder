package main

type User struct {
	ID       int
	Name     string
	Archived bool
}

type UserList []*User

func (u *UserList) ListParams() []interface{} {
	return []interface{}{}
}

func (u *UserList) ListTargets() []interface{} {
	return []interface{}{}
}

func (u *UserList) ListQuery() string {
	return "SELECT id, name, archived FROM users WHERE archived = false"
}

func (u *User) CreateParams() []interface{} {
	return []interface{}{u.Name}
}

func (u *User) CreateTargets() []interface{} {
	return []interface{}{
		&u.ID,
		&u.Name,
		&u.Archived,
	}
}

func (u *User) CreateQuery() string {
	return "INSERT INTO users (name) VALUES ($1) RETURNING id, name, archived"
}

func (u *User) ReadParams() []interface{} {
	return []interface{}{u.ID}
}

func (u *User) ReadTargets() []interface{} {
	return []interface{}{
		&u.ID,
		&u.Name,
		&u.Archived,
	}
}

func (u *User) ReadQuery() string {
	return "SELECT id, name, archived FROM users WHERE id=$1 AND archived = false"
}

func (u *User) UpdateParams() []interface{} {
	return []interface{}{u.ID, u.Name}
}

func (u *User) UpdateTargets() []interface{} {
	return []interface{}{
		&u.ID,
		&u.Name,
		&u.Archived,
	}
}

func (u *User) UpdateQuery() string {
	return "UPDATE users SET name = $2 WHERE id=$1 AND archived = false RETURNING id, name, archived"
}

func (u *User) DeleteParams() []interface{} {
	return []interface{}{u.ID}
}

func (u *User) DeleteTargets() []interface{} {
	return []interface{}{
		&u.ID,
		&u.Name,
		&u.Archived,
	}
}

func (u *User) DeleteQuery() string {
	return "UPDATE users SET archived=true WHERE id=$1 RETURNING id, name, archived"
}
