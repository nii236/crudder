package main

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/manveru/faker"
)

func main() {
	conn, err := sqlx.Connect("postgres", "user=dev dbname=dev password=dev sslmode=disable")
	if err != nil {
		fmt.Println(err)
		return
	}
	migrate(conn)
	svc := &DB{
		conn: conn,
	}

	f, err := faker.New("en")
	if err != nil {
		fmt.Println(err)
		return
	}
	for i := 0; i < 10; i++ {
		user := &User{
			Name: f.Name(),
		}

		err = svc.Create(user)
		if err != nil {
			fmt.Println(err)
			return
		}

	}

	fmt.Println("Get first user...")
	firstUser := &User{ID: 1}
	err = svc.Read(firstUser)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", firstUser)

	fmt.Println("Update first user...")
	firstUser.Name = "John Nguyen"
	err = svc.Update(firstUser)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", firstUser)

	fmt.Println("Delete first user...")
	err = svc.Delete(firstUser)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", firstUser)

	fmt.Println("List all users...")
	allUsers := UserList{}
	err = svc.List(&allUsers)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, user := range allUsers {
		fmt.Printf("%+v\n", user)
	}
}

func migrate(conn *sqlx.DB) {
	_, err := conn.Exec(`
DROP TABLE IF EXISTS users;
	`)
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = conn.Exec(`
CREATE TABLE IF NOT EXISTS users (
	id SERIAL PRIMARY KEY NOT NULL,
	name text NOT NULL,
	archived boolean NOT NULL DEFAULT false
);	
`)
	if err != nil {
		fmt.Println(err)
		return
	}
}
