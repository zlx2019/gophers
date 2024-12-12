package user

import (
	"database/sql"
	"errors"

	"github.com/zlx2019/ecom/types"
)

var ErrNotFound = errors.New("No records found")

type Store struct {
	*sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db}
}

func (s *Store) GetUserByUsername(username string) (*types.User, error) {
	rows, err := s.Query("SELECT * FROM user WHERE username = ?", username)
	if err != nil {
		return nil, err
	}
	user := new(types.User)
	for rows.Next() {
		user, err = scanRowsIntoUser(rows)
		if err != nil {
			return nil, err
		}
	}
	if user.ID == 0 {
		return nil, ErrNotFound
	}
	return user, nil
}

func (s *Store) GetUserByID(id int) (*types.User, error) {
	return nil, nil
}
func (s *Store) CreateUser(*types.User) error {
	// res, err := s.Exec("insert into user(first_name,) values ()")
	return nil

}

// 将读取到的数据，映射到实体
func scanRowsIntoUser(rows *sql.Rows) (*types.User, error) {
	user := new(types.User)
	err := rows.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.Username,
		&user.Password,
		&user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return user, nil
}
