/*In this first version I'll have a mess all around my code, next step is clean it up.*/

package main

import(
	"ProyCarLogs/DbHandle"
	"time"
	"fmt"
)

//Functions to print the information
func printCars(cars []DbHandle.Car){
	fmt.Println("Id \t Make \t\t Model \t Year \t Style")
	for _, c := range cars{
		fmt.Printf("%v \t %v \t   %v \t %v \t %v \n", c.Id, c.MakeStr, c.ModelStr, c.Year, c.StyleStr)
	}
	fmt.Println("")
}

func printMech(mech []DbHandle.Mechanic){
	fmt.Println("Id  Workshop  Mechanic  Specialty  Address  Phone")
	for _, m := range mech{
		fmt.Printf("%v  %v  %v  %v  %v  %v\n", m.Id, m.WorkshopName, m.MechanicName, m.SpecialtyStr, m.Address, m.Phone)
	}
	fmt.Println("")
}

func printLog(logs []DbHandle.Log){
	fmt.Println("Id  Car  Workshop  Problem  Solution  Date  Next date")
	for _, l := range logs{
		fmt.Printf("%v %v %v %v %v %v %v\n",l.Id, l.CarStr, l.MechanicStr, l.Problem, l.Solution, l.Date, l.NextDate)
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
	vibe := DbHandle.Car{
		ModelId: 3,
		Year: 2010,
		StyleId: 3}
	vibe = vibe.AddItems(db)

	intrepid := DbHandle.Car{
		ModelId: 2,
		Year: 2000,
		StyleId: 5}
	intrepid = intrepid.AddItems(db)

	malibu := DbHandle.Car{
		ModelId: 7,
		Year: 2016,
		StyleId: 5}
	malibu = malibu.AddItems(db)

	//Find all the cars
	allCars := DbHandle.ReadAllCars(db)
	printCars(allCars)

	//Get info from the cars whose model is the same as the Car given
	cars := vibe.SearchItems(db)
	printCars(cars)

	cars = intrepid.SearchItems(db)
	printCars(cars)

	//change the year of the car and update it on the database
	intrepid.Year = 1998
	intrepid.UpdateItems(db)

	//see if the change was done
	cars = intrepid.SearchItems(db)
	printCars(cars)

	//delete the car
	intrepid.DeleteItems(db)

	//list all the cars to see if the car was deleted.
	allCars = DbHandle.ReadAllCars(db)
	printCars(allCars)


	//MECHANICS
	//Add Mechanics to the database
	manuel := DbHandle.Mechanic{
		WorkshopName: "Taller Fulanitos",
		MechanicName: "Manuel",
		SpecialtyId: 1,
		Address: "Avenida fulana #123 Col centro",
		Phone: 1234567}
	manuel = manuel.AddItems(db)

	wero := DbHandle.Mechanic{
		WorkshopName: "El Wero",
		MechanicName: "Roger",
		SpecialtyId: 4,
		Address: "Calle Alisos Col Piedras Negras",
		Phone: 6462053540}
	wero = wero.AddItems(db)

	//Get all the mechanics from the DB
	allMech := DbHandle.ReadAllMechanic(db)
	printMech(allMech)

	//Search for a specific mechanic
	mech := wero.SearchItems(db)
	printMech(mech)

	//Update the changes
	wero.Address = "Calle Alisos #160 Col. Piedras Negras"
	wero.UpdateItems(db)

	//Check if the changes were made
	mech = wero.SearchItems(db)
	printMech(mech)

	//Delete a mechanic
	wero.DeleteItems(db)

	//Get all the mechanics from the BD
	allMech = DbHandle.ReadAllMechanic(db)
	printMech(allMech)


	//LOGS
	date := time.Now()
	mail := "tsuki4u@gmail.com"

	//Add Logs to the database
	logVibe := DbHandle.Log{
		CarId: 1,
		MechanicId: 1,
		Problem: "Problema",
		Solution: "Solucion",
		Date: date,
		NextDate: date.AddDate(0,0,1),
	}
	logVibe = logVibe.AddItems(db, mail)

	logMalibu := DbHandle.Log{
		CarId: malibu.Id,
		MechanicId: 1,
		Problem: "Problem 1",
		Solution: "Solution 1",
		Date: date,
	}
	logMalibu = logMalibu.AddItems(db, mail)

	//Get all the logs from the DB
	logs := DbHandle.ReadAllLogs(db)
	printLog(logs)

	//Get an specific log
	logs = logMalibu.SearchItems(db)
	printLog(logs)

	//Update an specific log
	logMalibu.Problem = "Problem 101"
	logMalibu.UpdateItems(db, mail)

	//Delete an specific log
	logVibe.DeleteItems(db)

	//Get all the logs from the DB
	logs = DbHandle.ReadAllLogs(db)
	printLog(logs)


	//Delete all the information from tables.
	DbHandle.DeleteAllData(db)
}
