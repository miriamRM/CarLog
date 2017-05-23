/*In this first version I'll have a mess all around my code, next step is clean it up.*/

package main

import(
	//"database/sql"
	//_ "github.com/mattn/go-sqlite3"
	"ProyCarLogs/DbHandle"
)

/*
type Crud interface{
	Add()
	Read()
	Update()
	Delete()
}

type Car struct{
	Id	int
	Make	int
	Model	int
	Year	int
	Style	int
}

//func (c Car) Add(db *sql.DB){
//Create a new car and store it to the database
}//

type Mechanic struct{
	Id	     int
	WorkshopName string
	MechanicName string
	Specialty    int
	Address	     string
	Phone	     int
}

type Log struct{
	Id	 int
	Car	 int //TODO: foreign key to the ID of car
	Mechanic int //Same as above from table mechanic
	Problem	 string
	Solution string
	Date	 string
	NextDate bool   //Is there going to be a next appointment?
	MailDate string //The date of the next appointment
}*/

func main(){
	//OpenDB
	const DBPATH = "./data/CarLog.db"
	db := DbHandle.OpenDB(DBPATH)
	defer db.Close()

	//Fill catalogs that user won't manage
	DbHandle.CreateFillCatalogs(db)
	vibe := DbHandle.Car{Make: 7, Model: 3, Year: 2010, Style: 3}
	vibe.Add(db)

	intrepid := DbHandle.Car{Make: 1, Model: 2, Year: 2000, Style: 5}
	intrepid.Add(db)
}
