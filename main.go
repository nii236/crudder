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

	fmt.Println("Seed users")
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

	all := UserList{}
	err = svc.List(&all)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("Seed teams")
	for _, v := range all {
		for i := 0; i < 5; i++ {
			team := &Team{
				Name:   f.CompanyName(),
				UserID: v.ID,
			}

			err = svc.Create(team)
			if err != nil {
				fmt.Println(err)
				return
			}

		}
	}

	fmt.Println("Get teams for user")
	teamsByUser := TeamList{}
	svc.Reference(&teamsByUser, "users", "user_id", "2")
	for _, team := range teamsByUser {
		fmt.Printf("%+v\n", team)
	}
	test(svc)
}

func test(svc *DB) {
	fmt.Println("Get first user...")
	firstUser := &User{}
	err := svc.Read(firstUser, "1")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", firstUser)

	fmt.Println("Update first user...")
	firstUser.Name = "John Nguyen"
	err = svc.Update(firstUser, "1")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", firstUser)

	fmt.Println("Delete first user...")

	err = svc.Delete(firstUser, "1")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%+v\n", firstUser)

	fmt.Println("GetMany...")
	getManyUsers := UserList{}
	err = svc.GetMany(&getManyUsers, []string{"1", "2", "3"})
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, user := range getManyUsers {
		fmt.Printf("%+v\n", user)
	}

	fmt.Println("UpdateMany...")
	updateManyUsers := UserList{}
	updateTo := &User{Name: "Johnny"}

	err = svc.UpdateMany(&updateManyUsers, updateTo, []string{"1", "2", "5"})
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, user := range updateManyUsers {
		fmt.Printf("%+v\n", user)
	}

	// fmt.Println("DeleteMany...")
	// deleteManyUsers := UserList{}
	// err = svc.DeleteMany(&deleteManyUsers, 1, 2, 3, 4)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// for _, user := range deleteManyUsers {
	// 	fmt.Printf("%+v\n", user)
	// }

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

	fmt.Println("Done.")
}

func migrate(conn *sqlx.DB) {
	_, err := conn.Exec(`
	DROP TABLE IF EXISTS teams;
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

	CREATE TABLE IF NOT EXISTS teams (
		id SERIAL PRIMARY KEY NOT NULL,
		name text NOT NULL,
		user_id INTEGER NOT NULL REFERENCES users(id),
		archived boolean NOT NULL DEFAULT false
	);
	`)
	if err != nil {
		fmt.Println(err)
		return
	}
}
