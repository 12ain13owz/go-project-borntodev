package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/12ain13owz/project-borntodev/model"
)

var Database *sql.DB

func SetupDatabase() error {
	var err error
	dbType := "mysql"
	dbName := "/coursedb"
	user := "root"
	pass := ":123456"
	port := "@(127.0.0.1:3306)"
	Database, err = sql.Open(dbType, user+pass+port+dbName)

	if err != nil {
		fmt.Println("Cannot connect Database")
		return err
	}

	fmt.Println(Database)
	fmt.Println("Connected Successfully")
	Database.SetConnMaxLifetime(time.Minute * 3)
	Database.SetMaxOpenConns(10)
	Database.SetMaxIdleConns(10)
	return nil
}

func GetUsers() ([]model.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := "SELECT id, name, email FROM users"
	result, err := Database.QueryContext(ctx, query)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer result.Close()

	users := make([]model.User, 0)
	for result.Next() {
		var user model.User
		result.Scan(&user.ID, &user.Name, &user.Email)
		users = append(users, user)
	}

	return users, nil
}

func GetUserByID(id int) (*model.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	query := "SELECT id, name, email FROM users WHERE id = ?"
	row := Database.QueryRowContext(ctx, query, id)
	user := &model.User{}
	err := row.Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return user, nil
}
