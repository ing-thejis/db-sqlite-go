package models

import(
	"errors"
	"time"
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func GetConn() *sql.DB {
	if db != nil{
		return db
	}

	var err error

	db, err = sql.Open("sqlite3", "notes.db")
	if err != nil{
		panic(err)
	}

	return db
}

func (n Note) Create() error{
	db := GetConn()

	stmt, err := db.Prepare("INSERT INTO notes (title, description, update_at) VALUES (?, ? ,?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(n.Title, n.Description, time.Now())
	if err != nil{
		return err
	}
	if i, err := result.RowsAffected(); err != nil || i != 1 {
		return errors.New("ERROR: se esperaba a row affected")
	}
	return nil
}

func (n *Note) GetAll() ([]Note, error){
	db := GetConn()

	rows, err := db.Query("SELECT id, title, description, create_at, update_at FROM notes")
	if err != nil{
		return []Note{}, err
	}
	defer rows.Close()

	notes := []Note{}

	for rows.Next() {
		rows.Scan(&n.ID, &n.Title, &n.Description, &n.CreatedAt, &n.UpdateAt)
		notes = append(notes, *n)
	}
	return notes, nil
}

func (n Note) Update() error {
	db := GetConn()

	stmt, err := db.Prepare("UPDATE notes set title=?, description=?, update_at=? WHERE id=?")
	if err != nil{
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(n.Title, n.Description, time.Now(), n.ID)

	if i, err := result.RowsAffected(); err != nil || i != 1 {
		return errors.New("ERROR: se esperaba a row affected")
	}
	return nil
}

func (n Note) Delete(id int) error {
	db := GetConn()

	stmt, err := db.Prepare("DELETE FROM notes WHERE id=?")
	if err != nil{
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(id)

	if i, err := result.RowsAffected(); err != nil || i != 1 {
		return errors.New("ERROR: se esperaba a row affected")
	}
	return nil
}

