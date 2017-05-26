package main

import(
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"fmt"
	"time"
)

type Users struct{
	Id	int
	Name	string
	Date	time.Time
	NextDate time.Time
}

func CreateTable(db *sql.DB) {
	// create table if not exists
	sql_table := `
	CREATE TABLE IF NOT EXISTS users(
		Id INTEGER NOT NULL PRIMARY KEY,
		Name TEXT,
		InsertedDatetime DATETIME,
		NextDate DATETIME);`

	_, err := db.Exec(sql_table)
	if err != nil {
		panic(err)
	}
}

func (user Users) StoreItem(db *sql.DB) {
	sql_additem := `
	INSERT OR REPLACE INTO users(
		Name,
		InsertedDatetime,
		NextDate
	) values(?,?,?)` //CURRENT_TIMESTAMP)


	stmt, err := db.Prepare(sql_additem)
	if err != nil { panic(err) }
	defer stmt.Close()

	_, err = stmt.Exec(user.Name, user.Date, user.NextDate)
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
		err := rows.Scan(&user.Id, &user.Name, &user.Date, &user.NextDate)
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

	now := time.Now()
	y,m,d := time.Now().Date()

	fmt.Println("now", now)
	fmt.Println("date", d,m,y)

	mimi := Users{Name: "Miriam", Date: now, NextDate: now.AddDate(0,1,0)}
	mimi.StoreItem(db)

	oscar := Users{Name: "Oscar", Date: now, NextDate: now.AddDate(1,0,3)}
	oscar.StoreItem(db)

	users := ReadItem(db)
	fmt.Println("Id \t Name \t Date \t NextDate")
	for _, user := range users{
		fmt.Print(user.Id, "\t", user.Name, "\t")
		y, m, d = user.Date.Date()
		fmt.Print(d, m, y, "\t")
		y, m, d = user.NextDate.Date()
		fmt.Print(d, m, y, "\n")
	}
}
