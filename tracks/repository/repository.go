package repository

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// the database
type Repository struct {
	DB *sql.DB
}

// creates an instance of the database
var repo Repository

// initializes the db
func Init() {
	if db, err := sql.Open("sqlite3", "/tmp/microservices.db"); err == nil {
		repo = Repository{DB: db}
	} else {
		log.Fatal("Database initialisation")
	}
}

// creates a table called Tracks
func Create() int {
	const sql = "CREATE TABLE IF NOT EXISTS Tracks" +
		"(Id TEXT PRIMARY KEY, Audio TEXT)"
	if _, err := repo.DB.Exec(sql); err == nil {
		return 0
	} else {
		return -1
	}
}

// clear the Tracks table (removes all rows)
func Clear() int {
	const sql = "DELETE FROM Tracks"
	if _, err := repo.DB.Exec(sql); err == nil {
		return 0
	} else {
		return -1
	}
}

// update a row in the Tracks table
func Update(t Track) int64 {
	const sql = "UPDATE Tracks SET Audio = ? WHERE Id = ?"
	// use a prepared statement
	if stmt, err := repo.DB.Prepare(sql); err == nil {
		defer stmt.Close()
		if res, err := stmt.Exec(t.Audio, t.Id); err == nil {
			if n, err := res.RowsAffected(); err == nil {
				return n
			}
		}
	}
	return -1
}

// insert a new row to the Tracks table
func Insert(t Track) (Track, int64) {
	const sql = "INSERT INTO Tracks(Id, Audio) VALUES (?, ?)"
	// use a prepared statement
	if stmt, err := repo.DB.Prepare(sql); err == nil {
		defer stmt.Close()
		if res, err := stmt.Exec(t.Id, t.Audio); err == nil {
			if n, err := res.RowsAffected(); err == nil {
				return t, n
			}
		}
	}
	return t, -1 // duplicate row entered or other error

}

// list all the rows in the Track table
func List() ([]string, int64) {
	const sql = "SELECT * FROM Tracks"
	var tracks []string
	if rows, err := repo.DB.Query(sql); err == nil {
		defer rows.Close()
		var id string
		var audio string
		var count int64 = 0
		for rows.Next() {
			if err := rows.Scan(&id, &audio); err == nil { // read row values into variables
				tracks = append(tracks, id)
				count += 1
			} else {
				return tracks, -1 // error
			}
		}
		return tracks, count // 0 or more rows affected

	}
	return tracks, -1

}

// read one row from the Tracks table
func Read(id string) (Track, int64) {
	const sql = "SELECT * FROM Tracks WHERE Id = ?"
	// use a prepared statement
	if stmt, err := repo.DB.Prepare(sql); err == nil {
		defer stmt.Close()
		var t Track
		row := stmt.QueryRow(id) // returns one row
		if err := row.Scan(&t.Id, &t.Audio); err == nil {
			return t, 1
		} else {
			return Track{}, 0
		}
	}
	return Track{}, -1
}

// delete a row from the Tracks table
func Delete(id string) int64 {
	const sql = "DELETE FROM Tracks WHERE Id = ?"
	// use a prepared statement
	if stmt, err := repo.DB.Prepare(sql); err == nil {
		defer stmt.Close()
		if res, err := stmt.Exec(id); err == nil {
			if n, err := res.RowsAffected(); err == nil {
				return n
			}
		}
	}
	return -1

}
