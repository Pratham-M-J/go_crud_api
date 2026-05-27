package sqlite

import (
	"database/sql"
	
	"github.com/Pratham-M-J/crud_api/internal/config"
	_ "modernc.org/sqlite" //underscore coz it's being used in the background not in the code itself
)

type Sqlite struct{
	Db *sql.DB
}
func (s *Sqlite) CreateStudent(name string, email string, age int,)(int64, error){
	//to avoid sql injection we first prepare the data and then execute
	stmt, err := s.Db.Prepare("INSERT INTO students (name, email, age) VALUES (?, ?, ?)")
	if err!=nil {
		return 0, err
	}

	defer stmt.Close()

	result, err := stmt.Exec(name, email, age)
	if err != nil{
		return 0, err
	}
	id, err := result.LastInsertId()

	if err != nil{
		return 0, err
	}
	return id, nil

}
func New(cfg *config.Config)(*Sqlite, error){
	db, err := sql.Open("sqlite", cfg.StoragePath)
	if err != nil{
		return nil, err
	}
	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS students(
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT,
		email TEXT UNIQUE,
		age INTEGER
	)
`)
	if err != nil{
		return nil, err
	}

	return &Sqlite{
		Db: db,
	}, nil

}