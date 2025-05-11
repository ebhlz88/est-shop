package store

import (
	"fmt"
	"time"

	"github.com/ebhlz88/est-shop/models"
)

func (s PostgresStore) CreateUserTable() error {
	query := `
	create table if not exists users(
		id serial primary key,
		name varchar(50),
		username varchar(50),
		password varchar(150),
		number numeric,
		created_at timestamp
		)
		`
	_, err := s.db.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

func (s *PostgresStore) CreateUser(name, username, password string, number int, createAt time.Time) error {
	query := `insert into users
	(name, username, password, number, created_at) values ($1, $2, $3, $4, $5)`
	_, err := s.db.Exec(query, name, username, password, number, createAt)
	return err
}

func (s *PostgresStore) GetUser() ([]models.User, error) {
	query := `select id, name, username,password, number, created_at from users`
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var users []models.User
	for rows.Next() {
		var user models.User
		err := rows.Scan(&user.ID, &user.Name, &user.UserName, &user.Password, &user.Number, &user.CreateAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (s *PostgresStore) GetUserById(id int) (models.User, error) {
	query := `select id, name, number, created_at from users where id=$1`
	var user models.User
	err := s.db.QueryRow(query, id).Scan(&user.ID, &user.Name, &user.Number, &user.CreateAt)
	return user, err
}

func (s *PostgresStore) GetUserByUsername(username string) (models.User, error) {
	fmt.Println(username)
	query := `select id, username, password, created_at from users where username=$1`
	var user models.User
	err := s.db.QueryRow(query, username).Scan(&user.ID, &user.UserName, &user.Password, &user.CreateAt)
	return user, err
}
