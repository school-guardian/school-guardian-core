package repository

import (
	"database/sql"
	"fmt"
	"gin/models"

	_ "github.com/lib/pq"
)

type DatabaseInterface interface {
	InsertNewUser(docs models.NewDocument) (bool, error)
	GetUser(id uint) ([]byte, error)
	GetAllUsers() ([]byte, error)
	UpdateUser(id uint) (bool, error)
	DeleteUser(id uint) (bool, error)
}

type Database struct {
	db *sql.DB
}

func NewDatabase(host, port, user, password, dbname string) (*Database, error) {
	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &Database{db: db}, nil
}

func (db *Database) Close() {
	if db.db != nil {
		db.db.Close()
	}
}

func (db *Database) InsertNewUser(docs models.NewDocument) (bool, error) {
	return true, nil
}

func (db *Database) GetUser(id int64) (bool, error) {
	return true, nil
}

func (db *Database) GetAllUsers() (bool, error) {
	return true, nil
}

func (db *Database) UpdateUser(id int64) (bool, error) {
	return true, nil
}

func (db *Database) DeleteUser(id int64) (bool, error) {
	return true, nil
}
