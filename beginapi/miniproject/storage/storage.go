package storage

import (
	"database/sql"
	_ "github.com/lib/pq"
	"najottalim/january-22/miniproject/model"
)

func connect() (*sql.DB, error) {
	dsn := "user=postgres password=123 dbname=miniproject sslmode=disable"
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	return db, nil
}
func CreateUser(user *model.User) (*model.User, error) {
	db, err := connect()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	respUser := model.User{}
	err = db.QueryRow(`INSERT INTO users(id,first_name,last_name) VALUES ($1,$2,$3) RETURNING id,first_name,last_name`,
		user.ID,
		user.FirstName,
		user.LastName).Scan(
		&respUser.ID,
		&respUser.FirstName,
		&respUser.LastName)
	if err != nil {
		return nil, err
	}
	return &respUser, nil
}

func GetUser(userID string) (*model.User, error) {
	db, err := connect()
	if err != nil {
		return nil, err
	}

	respUser := model.User{}
	err = db.QueryRow(`SELECT id,first_name,last_name FROM users 	WHERE id = $1`,
		userID).Scan(
		&respUser.ID,
		&respUser.FirstName,
		&respUser.LastName)
	if err != nil {
		return nil, err
	}
	return &respUser, nil
}
func GetAll(page, limit int) (users []*model.User, err error) {
	db, err := connect()
	if err != nil {
		return nil, err
	}
	//offset = (limit-1) * page -> bu juda kerakli formuala
	offset := limit * (page - 1)
	rows, err := db.Query(`SELECT id,first_name,last_name FROM users LIMIT $1 OFFSET $2`, limit, offset)
	for rows.Next() {
		var user model.User
		err := rows.Scan(&user.ID, &user.FirstName, &user.LastName)
		if err != nil {
			return nil, err
		}

		users = append(users, &user)
	}
	return users, nil
}

func UpdatedUser(users []*model.User) (alusers []model.User, err error) {
	db, err := connect()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var respUsers []model.User
	for _, user := range users {
		var respUser model.User
		err := db.QueryRow(`UPDATE users SET first_name = $1, last_name = $2 WHERE id = $3 RETURNING id, first_name, last_name`, user.FirstName, user.LastName, user.ID).Scan(
			&respUser.ID,
			&respUser.FirstName,
			&respUser.LastName)
		if err != nil {
			return nil, err
		}
		respUsers = append(respUsers, respUser)
	}
	if err != nil {
		return nil, err
	}
	return respUsers, nil
}

func DeleteUser(users []*model.User) (alusers []model.User, err error) {
	db, err := connect()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var respUsers []model.User
	for _, user := range users {
		var respUser model.User
		err := db.QueryRow(`DELETE FROM users WHERE id = $1 RETURNING id, first_name, last_name`, user.ID).Scan(
			&respUser.ID,
			&respUser.FirstName,
			&respUser.LastName)
		if err != nil {
			return nil, err
		}
		respUsers = append(respUsers, respUser)
	}
	if err != nil {
		return nil, err
	}
	return respUsers, nil
}
