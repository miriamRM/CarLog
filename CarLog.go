/*In this first version I'll have a mess all around my code, next step is clean it up.*/

package main

import(
	//"database/sql"
	//_ "github.com/mattn/go-sqlite3"
	"ProyCarLogs/DbHandle"
	"time"
	"fmt"
)

func printCars(cars []DbHandle.Car){
	fmt.Println("Id \t Make \t\t Model \t Year \t Style")
	for _, c := range cars{
		fmt.Printf("%v \t %v \t   %v \t %v \t %v \n", c.Id, c.MakeStr, c.ModelStr, c.Year, c.StyleStr)
	}
	fmt.Println("")
}

func main(){
	//OpenDB
	const DBPATH = "./data/CarLog.db"
	db := DbHandle.OpenDB(DBPATH)
	defer db.Close()

	//Fill catalogs that user won't manage
	DbHandle.CreateFillCatalogs(db)

	//CARS
	//Add Cars to the database
	vibe := DbHandle.Car{ModelId: 3, Year: 2010, StyleId: 3}
	vibe.AddItems(db)

	intrepid := DbHandle.Car{ModelId: 2, Year: 2000, StyleId: 5}
	intrepid.AddItems(db)

	caravan := DbHandle.Car{ModelId: 1, Year: 1996, StyleId: 4}
	caravan.AddItems(db)

	malibu := DbHandle.Car{ModelId: 7, Year: 2016, StyleId: 5}
	malibu.AddItems(db)

	//Find all the cars
	allCars := DbHandle.ReadAllCars(db)
	printCars(allCars)

	//Find  all the info from the cars whose model is the same as the Car given
	intrepids := intrepid.SearchItems(db)
	if len(intrepids) == 1 {
		intrepid = intrepids[0]
	}
	printCars(intrepids)

	//change the year of the car and update it on the database
	intrepid.Year = 1998
	intrepid.UpdateItems(db)

	//see if the change was done
	intrepids = intrepid.SearchItems(db)
	printCars(intrepids)

	//delete the car
	intrepid.DeleteItems(db)

	//list all the cars to see if the car was deleted.
	allCars = DbHandle.ReadAllCars(db)
	printCars(allCars)


	//MECHANICS
	//Add Mechanics to the database
	manuel := DbHandle.Mechanic{WorkshopName: "Taller Fulanitos",MechanicName: "Manuel", SpecialtyId: 1, Address: "Avenida fulana #123 Col centro", Phone: 1234567}
	manuel.AddItems(db)


	//LOGS
	//Add Logs to the database
	date := time.Now()
	logIntrepid := DbHandle.Log{CarId: 2, MechanicId: 1, Problem: "Intrepid sigue goteando aceite y seguimos sin saber por donde cae la gota", Solution: "Se le cambiaron los empaques de no se que cosa", Date: date , NextDate: date.AddDate(0,0,1)}
	logIntrepid.AddItems(db)


	//Delete all from tables.
	//DbHandle.DeleteAllData(db)

}
