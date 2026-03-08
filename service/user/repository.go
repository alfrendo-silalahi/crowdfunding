package user

import "database/sql"

type Repository interface {
	Save(user User) error
	FindByEmail(email string) (User, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &repository{db}
}

func (r *repository) Save(user User) error {
	err := r.db.QueryRow(`
	insert into users (username, occupation, email, password, role)
	values ($1, $2, $3, $4, $5)
	`, user.Username, user.Occupation, user.Email, user.Password, user.Role).Err()

	if err != nil {
		return err
	}

	return nil
}
func (r *repository) FindByEmail(email string) (User, error) {
	var user User

	err := r.db.QueryRow(`
	select email, password from users where email = $1
	`, email).Scan(&user.Email, &user.Password)

	if err != nil {
		return user, err
	}

	return user, nil
}
