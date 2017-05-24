/*In this first version I'll have a mess all around my code, next step is clean it up.*/

package main

import(
	//"database/sql"
	//_ "github.com/mattn/go-sqlite3"
	"ProyCarLogs/DbHandle"
	"time"
)

func main(){
	//OpenDB
	const DBPATH = "./data/CarLog.db"
	db := DbHandle.OpenDB(DBPATH)
	defer db.Close()

	//Fill catalogs that user won't manage
	DbHandle.CreateFillCatalogs(db)
	vibe := DbHandle.Car{ModelId: 3, Year: 2010, StyleId: 3}
	vibe.AddItems(db)

	intrepid := DbHandle.Car{ModelId: 2, Year: 2000, StyleId: 5}
	intrepid.AddItems(db)

	caravan := DbHandle.Car{ModelId: 1, Year: 1996, StyleId: 4}
	caravan.AddItems(db)

	malibu := DbHandle.Car{ModelId: 7, Year: 2016, StyleId: 5}
	malibu.AddItems(db)

	manuel := DbHandle.Mechanic{WorkshopName: "Taller Fulanitos",MechanicName: "Manuel", SpecialtyId: 1, Address: "Avenida fulana #123 Col centro", Phone: 1234567}
	manuel.AddItems(db)

	//Find the Id from intrepid

	date := time.Now()

	logIntrepid := DbHandle.Log{CarId: 2, MechanicId: 1, Problem: "Intrepid sigue goteando aceite y seguimos sin saber por donde cae la gota", Solution: "Se le cambiaron los empaques de no se que cosa", Date: date , NextDate: date.AddDate(0,0,1)}
	logIntrepid.AddItems(db)


}
