package mysqltests

import (
	"database/sql"
	"fmt"

	"github.com/lucastomic/syn-auth/internlas/database"
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
func connectTestingDB() (database.MYSQLDB, error) {
	db, err := sql.Open(
		"mysql",
		fmt.Sprintf("%s:%s@%s/%s?multiStatements=true", testUser, testPass, testHost, testDBName),
	)
	db.Ping()
	if err != nil {
		return database.MYSQLDB{}, err
	}
	return database.GetMYSQLDBWithDB(db)
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
