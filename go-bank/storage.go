package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// DataBase
type Storage interface {
	GetAccountByID(int) (*Account, error)
	GetAccounts() ([]*Account, error)
	CreateAccount(*Account) error
	UpdateAccount(*Account) error
	DeleteAccount(int) error
}

// Postgresql 客户端实例
type PostgresStore struct {
	db *sql.DB
}

// 创建 Postgresql 客户端实例
func NewPostgresStore(user, password, dbName string) (*PostgresStore, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		"192.168.1.223", 5432, user, password, dbName)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	return &PostgresStore{db}, nil
}

// 创建新的账户，返回自增生成的ID
func (s *PostgresStore) CreateAccount(account *Account) (int, error) {
	// 插入数据，通过 RETURNING 语句取得生成的ID
	sql := `insert into account 
					(first_name, last_name, number, balance, create_at) values
					($1, $2, $3, $4, $5) RETURNING id`
	var id int
	err := s.db.QueryRow(sql, account.FirstName, account.LastName, account.Number, account.Balance, account.CreatedAt).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}

// 修改账户
func (s *PostgresStore) UpdateAccount(account *Account) error {
	return nil
}

// 删除账户
func (s *PostgresStore) DeleteAccount(id int) error {
	_, err := s.db.Query("delete from account where id = $1", id)
	return err
}

// 根据ID获取账户信息
func (s *PostgresStore) GetAccountByID(id int) (*Account, error) {
	rows, err := s.db.Query("select * from account where id = $1", id)
	if err != nil {
		return nil, err
	}
	if rows.Next() {
		if account, err := scanIntoAccount(rows); err == nil {
			return account, nil
		}
	}
	return nil, fmt.Errorf("account %d not found", id)
}

// 获取所有账户信息
func (s *PostgresStore) GetAccounts() ([]*Account, error) {
	rows, err := s.db.Query("select id, first_name, last_name, number, balance, create_at from account")
	if err != nil {
		return nil, err
	}
	accounts := []*Account{}
	for rows.Next() {
		account := new(Account)
		if err = rows.Scan(&account.ID, &account.FirstName, &account.LastName, &account.Number, &account.Balance, &account.CreatedAt); err != nil {
			return nil, err
		}
		accounts = append(accounts, account)
	}
	return accounts, nil
}

// 将从数据库读取到的一行记录，映射为 Account 对象
func scanIntoAccount(rows *sql.Rows) (*Account, error) {
	account := new(Account)
	err := rows.Scan(&account.ID, &account.FirstName, &account.LastName, &account.Number, &account.Balance, &account.CreatedAt)
	if err != nil {
		return nil, err
	}
	return account, nil
}

// 创建数据表
func (s *PostgresStore) createTables() error {
	sql := `create table if not exists account (
				id 			serial primary key,
				first_name 	varchar(50),
				last_name 	varchar(50),
				number 		serial,
				balance 	numeric(10,2),
				create_at 	timestamp
			)`
	if _, err := s.db.Exec(sql); err != nil {
		return err
	}
	return nil
}
