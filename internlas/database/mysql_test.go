package database

import (
	"database/sql"
	"fmt"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

// For testing databse we use the next container:
// docker run --name mysql -e MYSQL_ROOT_PASSWORD=secret -v mysql/var/lib/mysql -p 33306:3306 mysql:8.0.32
// before testing we have to create all the database  (i don't know why if I have a docker volume)
// TODO: make docker-compose.yml for testing DB

// Relevant information to connect to the testing DB
var testUser = "root"
var testPass = "secret"
var testHost = "tcp(127.0.0.1:33306)"
var testDBName = "naturalYSalvaje"

// connectTestingDB connects to a testing databse to make secure tests
func connectTestingDB() (MYSQLDB, error) {
	db, err := sql.Open(
		"mysql",
		fmt.Sprintf("%s:%s@%s/%s?multiStatements=true", testUser, testPass, testHost, testDBName),
	)
	if err != nil {
		return MYSQLDB{}, err
	}
	return MYSQLDB{db}, nil
}

// truncateAllTables truncates all de tables from the testing DB.
func truncateAllTables() {
	db, err := sql.Open(
		"mysql",
		fmt.Sprintf("%s:%s@%s/%s?multiStatements=true", testUser, testPass, testHost, testDBName),
	)
	if err != nil {
		fmt.Println(err)
	}
	stm, _ := db.Prepare("SET FOREIGN_KEY_CHECKS=0")
	stm.Exec()
	stm, _ = db.Prepare("TRUNCATE naturalYSalvaje.sale;")
	stm.Exec()
	stm, _ = db.Prepare("TRUNCATE naturalYSalvaje.product")
	stm.Exec()
	stm, _ = db.Prepare("TRUNCATE naturalYSalvaje.user")
	stm.Exec()
	stm, _ = db.Prepare("TRUNCATE naturalYSalvaje.caliber")
	stm.Exec()
	stm, _ = db.Prepare("SET FOREIGN_KEY_CHECKS=1")
	stm.Exec()
	defer stm.Close()
}

// Set of tests for InsertInto function which whould return no error
var insertIntoTestsRight = []struct {
	table string
	body  map[string]any
}{
	{
		"product",
		map[string]any{
			"name":         "Merluza",
			"arrivalDate":  "2023-02-17",
			"arrivalPlace": "Barcelona",
		},
	},
	{
		"user",
		map[string]any{
			"name":      "Lucas Tomic",
			"nif":       "12345678Z",
			"phone":     "623029222",
			"ubication": "Madrid, Spain",
		},
	},
	{
		"caliber",
		map[string]any{
			"name":    "1-3kg",
			"amount":  25,
			"price":   10.5,
			"weight":  5.6,
			"product": "Merluza",
		},
	},
	{
		"sale",
		map[string]any{
			"date":          "2022-11-18",
			"payedAtMoment": true,
			"amount":        12,
			"userID":        1,
			"product":       "Merluza",
			"caliberName":   "1-3kg",
		},
	},
}

// Set of tests for InsertInto function which whould return error
var insertIntoTestsWrong = []struct {
	table string
	body  map[string]any
}{
	{
		"product",
		map[string]any{
			"arrivalDate":  "2023-02-17",
			"arrivalPlace": "Barcelona",
		},
	},
	{
		"user",
		map[string]any{
			"nif":       "12345678Z",
			"phone":     "623029222",
			"ubication": "Madrid, Spain",
		},
	},
	{
		"caliber",
		map[string]any{
			"name":    "1-3kg",
			"amount":  "A lot",
			"price":   10.5,
			"weight":  5.6,
			"product": "Merluza",
		},
	},
	{
		"sale",
		map[string]any{
			"payedAtMoment": "yes",
			"amount":        12,
			"caliberName":   "1-3kg",
		},
	},
}

func TestInsertInto(t *testing.T) {
	truncateAllTables()

	db, err := connectTestingDB()
	if err != nil {
		t.Errorf("Error connecting to the database: %v", err)
		return
	}
	for _, tt := range insertIntoTestsRight {
		t.Run(tt.table, func(t *testing.T) {
			// t.Parallel()
			err := db.InsertInto(tt.table, tt.body)
			if err != nil {
				t.Errorf("Error inserting the values in the database: %v", err)
			}
		})
	}
	for _, tt := range insertIntoTestsWrong {
		t.Run(tt.table, func(t *testing.T) {
			// t.Parallel()
			err := db.InsertInto(tt.table, tt.body)
			if err == nil {
				t.Error("Should return error")
			}
		})
	}
}
