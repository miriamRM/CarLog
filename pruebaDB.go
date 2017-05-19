package main

import(
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"fmt"
)

type Users struct{
	Id	int
	Name	string
	Date	string
}

func CreateTable(db *sql.DB) {
	// create table if not exists
	sql_table := `
	CREATE TABLE IF NOT EXISTS users(
		Id INTEGER NOT NULL PRIMARY KEY,
		Name TEXT,
		InsertedDatetime DATETIME);`

	_, err := db.Exec(sql_table)
	if err != nil {
		panic(err)
	}
}

func StoreItem(db *sql.DB, user Users) {
	sql_additem := `
	INSERT OR REPLACE INTO users(
		Name,
		InsertedDatetime
	) values(?, CURRENT_TIMESTAMP)
	`

	stmt, err := db.Prepare(sql_additem)
	if err != nil { panic(err) }
	defer stmt.Close()

	_, err = stmt.Exec(user.Name)
	if err != nil { panic(err) }
}

func ReadItem(db *sql.DB) []Users {
	sql_readall := `
	SELECT * FROM users
	ORDER BY datetime(InsertedDatetime) DESC
	`

	rows, err := db.Query(sql_readall)
	if err != nil { panic(err) }
	defer rows.Close()

	var result []Users
	for rows.Next() {
		var user Users
		err := rows.Scan(&user.Id, &user.Name, &user.Date)
		if err != nil { panic(err) }
		result = append(result, user)
	}
	return result
}

func main(){
	db, err := sql.Open("sqlite3","./data/test.db")
	if err != nil {
		log.Fatal(err)
	}
	if db == nil {
		panic("db nil")
	}
	defer db.Close()

	CreateTable(db)

	mimi := Users{Name: "Miriam"}
	StoreItem(db, mimi)

	oscar := Users{Name: "Oscar"}
	StoreItem(db, oscar)

	users := ReadItem(db)
	for _, user := range users{
		fmt.Println(user)
	}
}
