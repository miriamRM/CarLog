/*Script that handles the connection, table creation etc from proyect CarLog*/

package DbHandle

import(
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"fmt"
)

type Crud interface{
        Add()
        Read()
        Update()
        Delete()
}

type Car struct{
        Id      int
        Make    int
        Model   int
        Year    int
        Style   int
}

func (c Car) Add(db *sql.DB){
//Create a new car and store it to the database
	sql_additem := `
        INSERT OR REPLACE INTO Cars(
		MakeId,
                ModelId,
                Year,
		StyleId
        ) values(?,?,?,?)
        `

        stmt, err := db.Prepare(sql_additem)
        if err != nil{
		panic(err)
	}
        defer stmt.Close()

        _, err = stmt.Exec(c.Make,c.Model,c.Year,c.Style)
        if err != nil{
		panic(err)
	}
}

func (c Car) Read(){
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

type Mechanic struct{
        Id           int
        WorkshopName string
        MechanicName string
        Specialty    int
        Address      string
        Phone        int
}

type Log struct{
        Id       int
        Car      int //TODO: foreign key to the ID of car
        Mechanic int //Same as above from table mechanic
        Problem  string
        Solution string
        Date     string
        NextDate bool   //Is there going to be a next appointment?
        MailDate string //The date of the next appointment
}

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


func CreateFillCatalogs(db *sql.DB){
	CreateAllTables(db)
	FillCatalogs(db)
}


