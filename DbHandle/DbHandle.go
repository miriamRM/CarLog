/*Script that handles the connection, table creation etc from proyect CarLog*/

package DbHandle

import(
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"ProyCarLogs/GoogleAPIs/CalendarAPI"
	"ProyCarLogs/GoogleAPIs/GmailAPI"
	"time"
	"fmt"
)

func checkErr(err error){
	if err != nil{
		panic(err)
	}
}

// Below are the functions to open and create the tables on the DB
func OpenDB(path string) *sql.DB{
	db, err := sql.Open("sqlite3", path)
	checkErr(err)
	if db == nil {
		panic("db nil")
	}
	return db
}

func CreateTableStyles(db *sql.DB){
	sql_table := `
		CREATE TABLE IF NOT EXISTS Styles(
		Id INTEGER NOT NULL PRIMARY KEY,
		Style TEXT NOT NULL);`

	_, err := db.Exec(sql_table)
	checkErr(err)
}

func CreateTableMakes(db *sql.DB){
	sql_table := `
		CREATE TABLE IF NOT EXISTS Makes(
		Id INTEGER NOT NULL PRIMARY KEY,
		Make TEXT NOT NULL);`

	_, err := db.Exec(sql_table)
	checkErr(err)
}


func CreateTableModel(db *sql.DB){
	sql_table := `
		CREATE TABLE IF NOT EXISTS Model(
		Id INTEGER NOT NULL PRIMARY KEY,
		MakeId INTEGER NOT NULL,
		Model TEXT NOT NULL,
		FOREIGN KEY(MakeId) REFERENCES Makes(Id));`

	_, err := db.Exec(sql_table)
	checkErr(err)
}

func CreateTableCars(db *sql.DB){
	sql_table := `
		CREATE TABLE IF NOT EXISTS Cars(
		Id INTEGER NOT NULL PRIMARY KEY,
		MakeId INTEGER NOT NULL,
		ModelId INTEGER NOT NULL,
		Year INTEGER NOT NULL,
		StyleId INTEGER NOT NULL,
		FOREIGN KEY(StyleId) REFERENCES Styles(Id),
		FOREIGN KEY(MakeId) REFERENCES Makes(Id),
		FOREIGN KEY(ModelId) REFERENCES Model(Id));`

	_, err := db.Exec(sql_table)
	checkErr(err)
}

func CreateTableSpecialty(db *sql.DB){
	sql_table := `
		CREATE TABLE IF NOT EXISTS Specialty(
		Id INTEGER NOT NULL PRIMARY KEY,
		Specialty TEXT NOT NULL);`

	_, err := db.Exec(sql_table)
	checkErr(err)
}

func CreateTableMechanic(db *sql.DB){
	sql_table := `
		CREATE TABLE IF NOT EXISTS Mechanic(
		Id INTEGER NOT NULL PRIMARY KEY,
		WorkshopName TEXT NOT NULL,
		MechanicName TEXT NOT NULL,
		SpecialtyId INTEGER NOT NULL,
		Address TEXT,
		Phone INTEGER NOT NULL,
		FOREIGN KEY(SpecialtyId) REFERENCES Specialty(Id));`

	_, err := db.Exec(sql_table)
	checkErr(err)
}

func CreateTableLog(db *sql.DB){
	sql_table := `
		CREATE TABLE IF NOT EXISTS Log(
		Id INTEGER NOT NULL PRIMARY KEY,
		CarId INTEGER NOT NULL,
		MechanicId INTEGER NOT NULL,
		Problem TEXT NOT NULL,
		Solution TEXT NOT NULL,
		Date DATETIME NOT NULL, 
		NextDate DATETIME,
		FOREIGN KEY(CarId) REFERENCES Cars(Id),
		FOREIGN KEY(MechanicId) REFERENCES Mechanic(Id));` //al insertar una fecha hacerlo como CURRENT_TIMESTAMP

	_, err := db.Exec(sql_table)
	checkErr(err)
}

func CreateAllTables(db *sql.DB){
	CreateTableStyles(db)
	CreateTableMakes(db)
	CreateTableModel(db)
	CreateTableCars(db)
	CreateTableSpecialty(db)
	CreateTableMechanic(db)
	CreateTableLog(db)
}

// Below are the functions to Fill the catalogs that user do not manage (things that barely change)
func FillStyleTable(db *sql.DB){
	styles := []string{"SUV","PickUp","Hatchback","Van","Sedan"}
	sql_additem := `
	INSERT OR REPLACE INTO Styles(
		Style) 
	values(?)`

	stmt, err := db.Prepare(sql_additem)
	checkErr(err)
	defer stmt.Close()

	for _, style := range styles{
		_, err = stmt.Exec(style)
		checkErr(err)
	}
}

func FillMakesTable(db *sql.DB){
	makes := []string{"Dodge","Chrysler","Ford","Toyota","Nissan","Chevrolet","Pontiac"}
	sql_additem := `
	INSERT OR REPLACE INTO Makes(
		Make) 
	values(?)`

	stmt, err := db.Prepare(sql_additem)
	checkErr(err)
	defer stmt.Close()

	for _, m := range makes{ //make is a special word from GO
		_, err = stmt.Exec(m)
		checkErr(err)
	}
}

type Model struct{
	Id	int
	MakeId	int
	Model	string
}

func FillModelTable(db *sql.DB){
	models := []Model{
			{MakeId: 1, Model: "Caravan"},
			{MakeId: 1, Model: "Intrepid"},
			{MakeId: 7, Model: "Vibe"},
			{MakeId: 2, Model: "Town n Country"},
			{MakeId: 3, Model: "Escape"},
			{MakeId: 6, Model: "Spark"},
			{MakeId: 6, Model: "Malibu"},}

	sql_additem := `
	INSERT OR REPLACE INTO Model(
		MakeId,
		Model) 
	values(?,?)`

	stmt, err := db.Prepare(sql_additem)
	checkErr(err)
	defer stmt.Close()

	for _, model := range models{
		_, err = stmt.Exec(model.MakeId, model.Model)
		checkErr(err)
	}
}

func FillSpecialtyTable(db *sql.DB){
	specialty := []string{"General","Electrical","Transmission","Motor"}
        sql_additem := `
        INSERT OR REPLACE INTO Specialty(
                Specialty) 

        values(?)`

        stmt, err := db.Prepare(sql_additem)
	checkErr(err)
        defer stmt.Close()

        for _, special := range specialty{
                _, err = stmt.Exec(special)
		checkErr(err)
        }
}

func FillCatalogs(db *sql.DB){
	FillStyleTable(db)
	FillMakesTable(db)
	FillModelTable(db)
	FillSpecialtyTable(db)
}

// Below is the function used on the main to create and fill the catalogs
func CreateFillCatalogs(db *sql.DB){
	CreateAllTables(db)
	FillCatalogs(db)
}

//Functions to Delete all data on each table
func DeleteAllStyles(db *sql.DB){
	sqlDelAll := `
	DELETE FROM Styles`

	stmt, err := db.Prepare(sqlDelAll)
	checkErr(err)
	defer stmt.Close()

	resp, err := stmt.Exec()
	checkErr(err)

	row, err := resp.RowsAffected()
	checkErr(err)
	if row < 1 {
		panic("No rows affected")
	}
}

func DeleteAllMakes(db *sql.DB){
	sqlDelAll := `
	DELETE FROM Makes`

	stmt, err := db.Prepare(sqlDelAll)
	checkErr(err)
	defer stmt.Close()

	resp, err := stmt.Exec()
	checkErr(err)

	row, err := resp.RowsAffected()
	checkErr(err)
	if row < 1 {
		panic("No rows affected")
	}
}

func DeleteAllModels(db *sql.DB){
	sqlDelAll := `
	DELETE FROM Model`

	stmt, err := db.Prepare(sqlDelAll)
	checkErr(err)
	defer stmt.Close()

	resp, err := stmt.Exec()
	checkErr(err)

	row, err := resp.RowsAffected()
	checkErr(err)
	if row < 1 {
		panic("No rows affected")
	}
}

func DeleteAllSpecialties(db *sql.DB){
	sqlDelAll := `
	DELETE FROM Specialty`

	stmt, err := db.Prepare(sqlDelAll)
	checkErr(err)
	defer stmt.Close()

	resp, err := stmt.Exec()
	checkErr(err)

	row, err := resp.RowsAffected()
	checkErr(err)
	if row < 1 {
		panic("No rows affected")
	}
}

func DeleteAllCars(db *sql.DB){
	sqlDelAll := `
	DELETE FROM Cars`

	stmt, err := db.Prepare(sqlDelAll)
	checkErr(err)
	defer stmt.Close()

	resp, err := stmt.Exec()
	checkErr(err)

	row, err := resp.RowsAffected()
	checkErr(err)
	if row < 1 {
		panic("No rows affected")
	}
}

func DeleteAllMechanics(db *sql.DB){
	sqlDelAll := `
	DELETE FROM Mechanic`

	stmt, err := db.Prepare(sqlDelAll)
	checkErr(err)
	defer stmt.Close()

	resp, err := stmt.Exec()
	checkErr(err)

	row, err := resp.RowsAffected()
	checkErr(err)
	if row < 1 {
		panic("No rows affected")
	}
}

func DeleteAllLogs(db *sql.DB){
	sqlDelAll := `
	DELETE FROM Log`

	stmt, err := db.Prepare(sqlDelAll)
	checkErr(err)
	defer stmt.Close()

	resp, err := stmt.Exec()
	checkErr(err)

	row, err := resp.RowsAffected()
	checkErr(err)
	if row < 1 {
		panic("No rows affected")
	}
}

func DeleteAllData(db *sql.DB){
	DeleteAllStyles(db)
	DeleteAllMakes(db)
	DeleteAllModels(db)
	DeleteAllSpecialties(db)
	DeleteAllCars(db)
	DeleteAllMechanics(db)
	DeleteAllLogs(db)
}

// Below is the interface to manage the items.
type Crud interface{
        //AddItems(*sql.DB)      //No funciona porque regresa tipos diferentes
        //ReadAllItems(*sql.DB) //No funciona porque no requiere de un receiver
	//SearchItems(*sql.DB)  //No funciona porque regresa tipos diferentes
        UpdateItems(*sql.DB)
        DeleteItems(*sql.DB)
}

//Car Struct
type Car struct{
        Id        int
        MakeId    int
	MakeStr   string
        ModelId   int
	ModelStr  string
        Year      int
        StyleId   int
	StyleStr  string
}

//Cars Methods
func (c Car) AddItems(db *sql.DB) Car{
	sqlGetMake := `
	SELECT MakeId FROM Model WHERE Id = ?`

        rows, err := db.Query(sqlGetMake, c.ModelId)
	checkErr(err)
	defer rows.Close()

	var makeId int
        for rows.Next() {
                err := rows.Scan(&makeId)
                checkErr(err)
        }
	c.MakeId = makeId

	sqlAddItem := `
        INSERT INTO Cars(
		MakeId,
                ModelId,
                Year,
		StyleId
        ) values(?,?,?,?)`

        stmt, err := db.Prepare(sqlAddItem)
	checkErr(err)
        defer stmt.Close()

        res, err := stmt.Exec(makeId,c.ModelId,c.Year,c.StyleId)
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)
	c.Id = int(id)

	return c
}

func (c Car) SearchItems(db *sql.DB) []Car{
	sqlReadAll := `
	SELECT C.Id, C.MakeId, Ma.Make, C.ModelId, Mo.Model, C.Year, C.StyleId, S.Style
   	FROM Makes as Ma, Model as Mo, Cars as C, Styles as S
   	WHERE C.MakeId = Ma.Id 
   	AND C.ModelId = Mo.Id
   	AND C.StyleId = S.Id 
	AND C.ModelId = ?
	ORDER BY C.Id ASC`

	rows, err := db.Query(sqlReadAll, c.ModelId)
	checkErr(err)
	defer rows.Close()

	var result []Car
	for rows.Next() {
		var car Car
		err := rows.Scan(&car.Id, &car.MakeId, &car.MakeStr, &car.ModelId, &car.ModelStr, &car.Year, &car.StyleId, &car.StyleStr)
		checkErr(err)
		result = append(result, car)
	}
	return result
}

func (c Car) UpdateItems(db *sql.DB){
	sqlUpdateItem := `
	UPDATE Cars
	SET MakeId = ?, ModelId = ?, Year = ?, StyleId = ?
	WHERE Id = ?`

	stmt, err := db.Prepare(sqlUpdateItem)
	checkErr(err)
	defer stmt.Close()

	resp, err := stmt.Exec(c.MakeId, c.ModelId, c.Year, c.StyleId, c.Id)
	checkErr(err)

	row, err := resp.RowsAffected()
	checkErr(err)
	if row < 1{
		panic("No rows affected")
	}
}

func (c Car) DeleteItems(db *sql.DB){
	sqlDelItem := `
	DELETE FROM Cars
	WHERE Id = ?`

	stmt, err := db.Prepare(sqlDelItem)
	checkErr(err)
	defer stmt.Close()

	resp, err := stmt.Exec(c.Id)
	checkErr(err)

	row, err := resp.RowsAffected()
	checkErr(err)
	if row < 1 {
		panic("No rows affected")
	}
}

func ReadAllCars(db *sql.DB) []Car{ //read all items
	sqlReadAll := `
	SELECT C.Id, C.MakeId, Ma.Make, C.ModelId, Mo.Model, C.Year, C.StyleId, S.Style
   	FROM Makes as Ma, Model as Mo, Cars as C, Styles as S
   	WHERE C.MakeId = Ma.Id 
   	AND C.ModelId = Mo.Id
   	AND C.StyleId = S.Id 
	ORDER BY C.Id ASC`

	rows, err := db.Query(sqlReadAll)
	checkErr(err)
	defer rows.Close()

	var result []Car
	for rows.Next() {
		var car Car
		err := rows.Scan(&car.Id, &car.MakeId, &car.MakeStr, &car.ModelId, &car.ModelStr, &car.Year, &car.StyleId, &car.StyleStr)
		checkErr(err)
		result = append(result, car)
	}
	return result
}

//Mechanic Struct
type Mechanic struct{
        Id           int
        WorkshopName string
        MechanicName string
        SpecialtyId  int
        SpecialtyStr string
	Address      string
        Phone        int
}

//Mechanic Methods
func (m Mechanic) AddItems(db *sql.DB) Mechanic{
	sqlAddItem := `
        INSERT INTO Mechanic(
		WorkshopName,
                MechanicName,
                SpecialtyId,
		Address,
		Phone
        ) values(?,?,?,?,?)`

        stmt, err := db.Prepare(sqlAddItem)
	checkErr(err)
        defer stmt.Close()

        res, err := stmt.Exec(m.WorkshopName, m.MechanicName, m.SpecialtyId, m.Address, m.Phone)
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)
	m.Id = int(id)

	return m
}

func (m Mechanic) SearchItems(db *sql.DB) []Mechanic{
	sqlReadAll := `
	SELECT Me.Id, Me.WorkshopName, Me.MechanicName, Me.SpecialtyId, Sp.Specialty, Me.Address, Me.Phone 
   	FROM Mechanic as Me, Specialty as Sp
   	WHERE Me.SpecialtyId = Sp.Id 
	AND Me.WorkshopName = ?
	ORDER BY Me.Id ASC`

	rows, err := db.Query(sqlReadAll, m.WorkshopName)
	checkErr(err)
	defer rows.Close()

	var result []Mechanic
	for rows.Next() {
		var Me Mechanic
		err := rows.Scan(&Me.Id, &Me.WorkshopName, &Me.MechanicName, &Me.SpecialtyId, &Me.SpecialtyStr, &Me.Address, &Me.Phone)
		checkErr(err)
		result = append(result, Me)
	}
	return result
}

func (m Mechanic) UpdateItems(db *sql.DB){
	sqlUpdateItem := `
	UPDATE Mechanic
	SET WorkshopName = ?, MechanicName = ?, SpecialtyId = ?, Address = ?, Phone = ?
	WHERE Id = ?`

	stmt, err := db.Prepare(sqlUpdateItem)
	checkErr(err)
	defer stmt.Close()

	resp, err := stmt.Exec(m.WorkshopName, m.MechanicName, m.SpecialtyId, m.Address, m.Phone, m.Id)
	checkErr(err)

	row, err := resp.RowsAffected()
	checkErr(err)
	if row < 1{
		panic("No rows affected")
	}
}

func (m Mechanic) DeleteItems(db *sql.DB){
	sqlDelItem := `
	DELETE FROM Mechanic
	WHERE Id = ?`

	stmt, err := db.Prepare(sqlDelItem)
	checkErr(err)
	defer stmt.Close()

	resp, err := stmt.Exec(m.Id)
	checkErr(err)

	row, err := resp.RowsAffected()
	checkErr(err)
	if row < 1 {
		panic("No rows affected")
	}
}

func ReadAllMechanic(db *sql.DB) []Mechanic{
	sqlReadAll := `
	SELECT Me.Id, Me.WorkshopName, Me.MechanicName, Me.SpecialtyId, Sp.Specialty, Me.Address, Me.Phone 
   	FROM Mechanic as Me, Specialty as Sp
   	WHERE Me.SpecialtyId = Sp.Id 
	ORDER BY Me.Id ASC`

	rows, err := db.Query(sqlReadAll)
	checkErr(err)
	defer rows.Close()

	var result []Mechanic
	for rows.Next() {
		var Me Mechanic
		err := rows.Scan(&Me.Id, &Me.WorkshopName, &Me.MechanicName, &Me.SpecialtyId, &Me.SpecialtyStr, &Me.Address, &Me.Phone)
		checkErr(err)
		result = append(result, Me)
	}
	return result
}

//Log Struct
type Log struct{
        Id          int
        CarId       int
        CarStr      string
	MechanicId  int
	MechanicStr string
        Problem     string
        Solution    string
        Date        time.Time
        NextDate    time.Time
}

//Log Methods
func (l Log) AddItems(db *sql.DB, mail string) Log{
	sqlAddItem := `
        INSERT OR REPLACE INTO Log(
		CarId,
                MechanicId,
                Problem,
		Solution,
		Date,
		NextDate
        ) values(?,?,?,?,?,?)`

        stmt, err := db.Prepare(sqlAddItem)
	checkErr(err)
        defer stmt.Close()

        res, err := stmt.Exec(l.CarId, l.MechanicId, l.Problem, l.Solution, l.Date, l.NextDate)
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)
	l.Id = int(id)

	if l.NextDate.IsZero() != true{
		CreateEventPlusMail(l, db, mail)
	}

	return l
}


func (l Log) SearchItems(db *sql.DB) []Log{
	fmt.Println("log antes de select",l)
	sqlReadAll := `
	SELECT l.Id, l.CarId, mo.Model, l.MechanicId, me.WorkshopName, l.Problem, l.Solution, l.Date, l.NextDate
	FROM Log as l, Model as mo, Mechanic as me, Cars as c
	WHERE l.Id = ?
	AND l.CarId = c.Id
	AND c.ModelId = mo.Id
	AND l.MechanicId = me.Id`

	rows, err := db.Query(sqlReadAll, l.Id)
	checkErr(err)
	defer rows.Close()

	var result []Log
	for rows.Next() {
		var log Log
		err := rows.Scan(&log.Id, &log.CarId, &log.CarStr, &log.MechanicId, &log.MechanicStr, &log.Problem, &log.Solution, &log.Date, &log.NextDate)
		checkErr(err)
		result = append(result, log)
	}
	return result
}

func (l Log) UpdateItems(db *sql.DB, mail string){
	sqlUpdateItem := `
	UPDATE Log
	SET Problem = ?, Solution = ?, Date = ?, NextDate = ?
	WHERE Id = ?`

	stmt, err := db.Prepare(sqlUpdateItem)
	checkErr(err)
	defer stmt.Close()

	resp, err := stmt.Exec(l.Problem, l.Solution, l.Date, l.NextDate, l.Id)
	checkErr(err)

	row, err := resp.RowsAffected()
	checkErr(err)
	if row < 1{
		panic("No rows affected")
	}

	if l.NextDate.IsZero() != true{
		CreateEventPlusMail(l, db, mail)
	}
}

func (l Log) DeleteItems(db *sql.DB){
	sqlDelItem :=`
	Delete from Log 
	WHERE Id = ?`

	stmt, err := db.Prepare(sqlDelItem)
	checkErr(err)
	defer stmt.Close()

	resp, err := stmt.Exec(l.Id)
	checkErr(err)

	row, err := resp.RowsAffected()
	checkErr(err)
	if row < 1 {
		panic("No rows affected")
	}
}

func CreateEventPlusMail(l Log,db *sql.DB, mail string){
	sqlGetWorkshop := `
	SELECT me.WorkshopName
	FROM Mechanic as me, Log as l
	WHERE l.MechanicId = me.Id
	AND l.Id = ?`

	rows, err := db.Query(sqlGetWorkshop, l.Id)
        checkErr(err)
        defer rows.Close()

        var mechanic string
        for rows.Next() {
                err := rows.Scan(&mechanic)
                checkErr(err)
        }

	CalendarAPI.CreateEventCalendar(mechanic, l.Solution, l.NextDate)
	GmailAPI.CreateSendMail(mail)

}

func ReadAllLogs(db *sql.DB) []Log{
	sqlReadAll := `
	SELECT l.Id, l.CarId, mo.Model, l.MechanicId, me.WorkshopName, l.Problem, l.Solution, l.Date, l.NextDate
	FROM Log as l, Model as mo, Mechanic as me, Cars as c
	WHERE l.CarId = c.Id
	AND c.ModelId = mo.Id
	AND l.MechanicId = me.Id`

	rows, err := db.Query(sqlReadAll)
	checkErr(err)
	defer rows.Close()

	var result []Log
	for rows.Next() {
		var log Log
		err := rows.Scan(&log.Id, &log.CarId, &log.CarStr, &log.MechanicId, &log.MechanicStr, &log.Problem, &log.Solution, &log.Date, &log.NextDate)
		checkErr(err)
		result = append(result, log)
	}
	return result
}

