/*Script that handles the connection, table creation etc from proyect CarLog*/

package DbHandle

import(
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"fmt"
	"time"
)

// Below are the structs
type Mechanic struct{
        Id           int
        WorkshopName string
        MechanicName string
        SpecialtyId  int
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
        ModelId   int
        Year      int
        StyleId   int
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

// Below are the functions to open and create the tables on the DB
func OpenDB(path string) *sql.DB{
	db, err := sql.Open("sqlite3", path)
	if err != nil {
		panic(err)
	}
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
	if err != nil {
		panic(err)
	}
}

func CreateTableMakes(db *sql.DB){
	sql_table := `
		CREATE TABLE IF NOT EXISTS Makes(
		Id INTEGER NOT NULL PRIMARY KEY,
		Make TEXT NOT NULL);`

	_, err := db.Exec(sql_table)
	if err != nil {
		panic(err)
	}
}


func CreateTableModel(db *sql.DB){
	sql_table := `
		CREATE TABLE IF NOT EXISTS Model(
		Id INTEGER NOT NULL PRIMARY KEY,
		MakeId INTEGER NOT NULL,
		Model TEXT NOT NULL,
		FOREIGN KEY(MakeId) REFERENCES Makes(Id));`

	_, err := db.Exec(sql_table)
	if err != nil {
		panic(err)
	}
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
	if err != nil {
		panic(err)
	}
}

func CreateTableSpecialty(db *sql.DB){
	sql_table := `
		CREATE TABLE IF NOT EXISTS Specialty(
		Id INTEGER NOT NULL PRIMARY KEY,
		Specialty TEXT NOT NULL);`

	_, err := db.Exec(sql_table)
	if err != nil {
		panic(err)
	}
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
	if err != nil {
		panic(err)
	}
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
	if err != nil {
		panic(err)
	}
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
	if err != nil{
		panic(err)
	}
	defer stmt.Close()

	for _, style := range styles{
		_, err = stmt.Exec(style)
		if err != nil{
			panic(err)
		}
	}
}

func FillMakesTable(db *sql.DB){
	makes := []string{"Dodge","Chrysler","Ford","Toyota","Nissan","Chevrolet","Pontiac"}
	sql_additem := `
	INSERT OR REPLACE INTO Makes(
		Make) 
	values(?)`

	stmt, err := db.Prepare(sql_additem)
	if err != nil{
		panic(err)
	}
	defer stmt.Close()

	for _, m := range makes{ //make is a special word from GO
		_, err = stmt.Exec(m)
		if err != nil { panic(err) }
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
	if err != nil{
		panic(err)
	}
	defer stmt.Close()

	for _, model := range models{
		_, err = stmt.Exec(model.MakeId, model.Model)
		if err != nil{
			panic(err)
		}
	}
}

func FillSpecialtyTable(db *sql.DB){
	specialty := []string{"General","Electrical","Transmission","Motor"}
        sql_additem := `
        INSERT OR REPLACE INTO Specialty(
                Specialty) 
        values(?)`

        stmt, err := db.Prepare(sql_additem)
        if err != nil{
                panic(err)
		fmt.Println(err)
        }
        defer stmt.Close()

        for _, special := range specialty{
                _, err = stmt.Exec(special)
                if err != nil{
			panic(err)
		}
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

// Below are the methods to manage the items.
type Crud interface{
        AddItems()
        ReadItems()
        UpdateItems()
	SearchItmes()
        DeleteItems()
}

func (c Car) AddItems(db *sql.DB){
	sql_additem := `
        INSERT OR REPLACE INTO Cars(
		MakeId,
                ModelId,
                Year,
		StyleId
        ) values((SELECT MakeId FROM Model WHERE Id = ?),?,?,?)
        `

        stmt, err := db.Prepare(sql_additem)
        if err != nil{
		panic(err)
	}
        defer stmt.Close()

        _, err = stmt.Exec(c.ModelId,c.ModelId,c.Year,c.StyleId)
        if err != nil{
		panic(err)
	}
}

func (c Car) ReadItems(db *sql.DB){
/*
	sql_readall := `
	SELECT Ma.Make, Mo.Model, C.Year, S.Style
   	FROM Makes as Ma, Model as Mo, Cars as C, Styles as S
   	WHERE C.MakeId = Ma.Id 
   	AND C.ModelId = Mo.Id
   	AND C.StyleId = S.Id`

        rows, err := db.Query(sql_readall)
        if err != nil{
		panic(err)
	}
        defer rows.Close()

        for rows.Next() {
                //var car Car
                err := rows.Scan(&car., &user.Name, &user.Date)
                if err != nil{
			panic(err)
		}
                result = append(result, user)
        }
        return result
*/
}

func (c Car) UpdateItems(db *sql.DB){

}

func (c Car) SearchItmes(db *sql.DB){

}

func (c Car) DeleteItems(db *sql.DB){

}

//Mechanic methods
func (m Mechanic) AddItems(db *sql.DB){
	sql_additem := `
        INSERT OR REPLACE INTO Mechanic(
		WorkshopName,
                MechanicName,
                SpecialtyId,
		Address,
		Phone
        ) values(?,?,?,?,?)
        `

        stmt, err := db.Prepare(sql_additem)
        if err != nil{
		panic(err)
	}
        defer stmt.Close()

        _, err = stmt.Exec(m.WorkshopName, m.MechanicName, m.SpecialtyId, m.Address, m.Phone)
        if err != nil{
		panic(err)
	}
}

func (m Mechanic) ReadItems(db *sql.DB){
/*
	sql_readall := `
	SELECT Ma.Make, Mo.Model, C.Year, S.Style
   	FROM Makes as Ma, Model as Mo, Cars as C, Styles as S
   	WHERE C.MakeId = Ma.Id 
   	AND C.ModelId = Mo.Id
   	AND C.StyleId = S.Id`

        rows, err := db.Query(sql_readall)
        if err != nil{
		panic(err)
	}
        defer rows.Close()

        for rows.Next() {
                //var car Car
                err := rows.Scan(&car., &user.Name, &user.Date)
                if err != nil{
			panic(err)
		}
                result = append(result, user)
        }
        return result
*/
}

/*func () UpdateItems(db *sql.DB){

}

func () SearchItmes(db *sql.DB){

}

func () DeleteItems(db *sql.DB){

}*/

//Log methods
func (l Log) AddItems(db *sql.DB){
	sql_additem := `
        INSERT OR REPLACE INTO Log(
		CarId,
                MechanicId,
                Problem,
		Solution,
		Date,
		NextDate
        ) values(?,?,?,?,?,?)
        `

        stmt, err := db.Prepare(sql_additem)
        if err != nil{
		panic(err)
	}
        defer stmt.Close()

        _, err = stmt.Exec(l.CarId, l.MechanicId, l.Problem, l.Solution, l.Date, l.NextDate)
        if err != nil{
		panic(err)
	}
}

/* AddDate(<año>,<mes>,<dia>)
today := time.Now()
fmt.Println(today)
oneMonth := today.AddDate(0,1,0) 
fmt.Println(oneMonth)
*/

func (l Log) ReadItems(db *sql.DB){
/*
	sql_readall := `
	SELECT Ma.Make, Mo.Model, C.Year, S.Style
   	FROM Makes as Ma, Model as Mo, Cars as C, Styles as S
   	WHERE C.MakeId = Ma.Id 
   	AND C.ModelId = Mo.Id
   	AND C.StyleId = S.Id`

        rows, err := db.Query(sql_readall)
        if err != nil{
		panic(err)
	}
        defer rows.Close()

        for rows.Next() {
                //var car Car
                err := rows.Scan(&car., &user.Name, &user.Date)
                if err != nil{
			panic(err)
		}
                result = append(result, user)
        }
        return result
*/
}

/*func () UpdateItems(db *sql.DB){

}

func () SearchItmes(db *sql.DB){

}

func () DeleteItems(db *sql.DB){

}*/
