/*Script that handles the connection, table creation etc from proyect CarLog*/

package DbHandle

import(
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	//"fmt"
	"time"
)

// Below are the structs
type Mechanic struct{
        Id           int
        WorkshopName string
        MechanicName string
        SpecialtyId  int
        SpecialtyStr string
	Address      string
        Phone        int
}

type Log struct{
        Id         int
        CarId      int
        MechanicId int
        Problem    string
        Solution   string
        Date       time.Time
        NextDate   time.Time //The date of the next appointment
}

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

type Model struct{
	Id	int
	MakeId	int
	Model	string
}

//I dont need these yet
type Make struct{
	Id	int
	Make	string
}

type Style struct{
	Id	int
	Style	string
}

type Specialty struct{
	Id	  int
	Specialty string
}

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
        AddItems()
   //     ReadAllItems()
        UpdateItems()
	SearchItems()
        DeleteItems()
}

//Cars Methods
func (c Car) AddItems(db *sql.DB){
	sqlAddItem := `
        INSERT INTO Cars(
		MakeId,
                ModelId,
                Year,
		StyleId
        ) values((SELECT MakeId FROM Model WHERE Id = ?),?,?,?)`

        stmt, err := db.Prepare(sqlAddItem)
	checkErr(err)
        defer stmt.Close()

        _, err = stmt.Exec(c.ModelId,c.ModelId,c.Year,c.StyleId)
	checkErr(err)
}

func (c Car) SearchItems(db *sql.DB) []Car{ //Search only for the ones of the same model
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

//Mechanic methods
func (m Mechanic) AddItems(db *sql.DB){
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

        _, err = stmt.Exec(m.WorkshopName, m.MechanicName, m.SpecialtyId, m.Address, m.Phone)
	checkErr(err)
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
	WHERE Id = ?
`

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


//Log methods
func (l Log) AddItems(db *sql.DB){
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

        _, err = stmt.Exec(l.CarId, l.MechanicId, l.Problem, l.Solution, l.Date, l.NextDate)
	checkErr(err)
}

/* AddDate(<año>,<mes>,<dia>)
today := time.Now()
fmt.Println(today)
oneMonth := today.AddDate(0,1,0) 
fmt.Println(oneMonth)
*/

func (l Log) UpdateItems(db *sql.DB){

}

func (l Log) SearchItems(db *sql.DB){

}

func (l Log) DeleteItems(db *sql.DB){

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
/*
	sql_readall := `
	SELECT Ma.Make, Mo.Model, C.Year, S.Style
   	FROM Makes as Ma, Model as Mo, Cars as C, Styles as S
   	WHERE C.MakeId = Ma.Id 
   	AND C.ModelId = Mo.Id
   	AND C.StyleId = S.Id`

        rows, err := db.Query(sql_readall)
	checkErr(err)
        defer rows.Close()

        for rows.Next() {
                //var car Car
                err := rows.Scan(&car., &user.Name, &user.Date)
		checkErr(err)
                result = append(result, user)
        }
        return result
*/
}

func ReadAllLogs(db *sql.DB){
/*
	sql_readall := `
	SELECT Ma.Make, Mo.Model, C.Year, S.Style
   	FROM Makes as Ma, Model as Mo, Cars as C, Styles as S
   	WHERE C.MakeId = Ma.Id 
   	AND C.ModelId = Mo.Id
   	AND C.StyleId = S.Id`

        rows, err := db.Query(sql_readall)
	checkErr(err)
        defer rows.Close()

        for rows.Next() {
                //var car Car
                err := rows.Scan(&car., &user.Name, &user.Date)
		checkErr(err)
                result = append(result, user)
        }
        return result
*/
}

