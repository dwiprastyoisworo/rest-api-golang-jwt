package repositories

import (
	"context"
	"database/sql"
	"rest-api-golang-jwt/src/entity"
)

type UserRepository struct {
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (r UserRepository) Get(ctx context.Context, db *sql.DB) ([]entity.Users, error) {
	conn, err := db.Conn(ctx)
	if err != nil {
		return []entity.Users{}, err
	}
	defer conn.Close()
	rows, err := conn.QueryContext(ctx, `select id, "name" , address from users u`)
	if err != nil {
		return []entity.Users{}, err
	}
	defer rows.Close()
	var users []entity.Users
	for rows.Next() {
		var user entity.Users
		errRowScan := rows.Scan(&user.ID, &user.Name, &user.Address)
		if errRowScan != nil {
			return []entity.Users{}, errRowScan
		}
		users = append(users, user)
	}
	return users, nil
}

func (r UserRepository) Insert(ctx context.Context, db *sql.DB, user entity.UserRequest) (int, error) {
	conn, err := db.Conn(ctx)
	if err != nil {
		return 0, nil
	}
	defer conn.Close()
	stmt, err := conn.PrepareContext(ctx, `insert into users (id, "name", password,address,email) values ($1, $2, $3,$4,$5)`)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	_, err = stmt.ExecContext(ctx, user.ID, user.Name, user.Password, user.Address, user.Email)
	if err != nil {
		return 0, err
	}
	return 1, nil
}

// create function to update data user
func (r UserRepository) Update(ctx context.Context, db *sql.DB, user entity.UserRequest) (int, error) {
	conn, err := db.Conn(ctx)
	if err != nil {
		return 0, nil
	}
	defer conn.Close()
	stmt, err := conn.PrepareContext(ctx, `update users set "name" = $1, address = $2 where id = $3`)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	_, err = stmt.ExecContext(ctx, user.Name, user.Address, user.ID)
	if err != nil {
		return 0, err
	}
	return 1, nil
}

// create function to delete data user
func (r UserRepository) Delete(ctx context.Context, db *sql.DB, id string) (int, error) {
	conn, err := db.Conn(ctx)
	if err != nil {
		return 0, nil
	}
	defer conn.Close()
	stmt, err := conn.PrepareContext(ctx, `delete from users where id = $1`)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	_, err = stmt.ExecContext(ctx, id)
	if err != nil {
		return 0, err
	}
	return 1, nil
}

// create function to get data user by email
func (r UserRepository) GetByEmail(ctx context.Context, db *sql.DB, email string) (entity.Users, error) {
	conn, err := db.Conn(ctx)
	if err != nil {
		return entity.Users{}, err
	}
	defer conn.Close()
	row := conn.QueryRowContext(ctx, `select id, "name", password, address from users where email = $1`, email)
	var user entity.Users
	err = row.Scan(&user.ID, &user.Name, &user.Password, &user.Address)
	if err != nil {
		return entity.Users{}, err
	}
	return user, nil
}

// create function to get data user by id
func (r UserRepository) GetByID(ctx context.Context, db *sql.DB, id string) (entity.Users, error) {
	conn, err := db.Conn(ctx)
	if err != nil {
		return entity.Users{}, err
	}
	defer conn.Close()
	row := conn.QueryRowContext(ctx, `select id, "name", password, address from users where id = $1`, id)
	var user entity.Users
	err = row.Scan(&user.ID, &user.Name, &user.Password, &user.Address)
	if err != nil {
		return entity.Users{}, err
	}
	return user, nil
}
