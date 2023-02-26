package database

// DatabseInterface is the parent struct for databse systems
type DatabaseInterface interface {
	// InsertInto into inserts a row into the database.
	// In case of error, it returns it. Otherwise, returns nil.
	// Expects the table name as first parameter and a map with the values as second parameter.
	// For example:
	//
	// var values = map[string]any{
	// 	"name":"Lucas",
	// 	"age":22,
	// 	"country":"Argentina",
	// }
	// db.InsertInto("userTable", values)
	InsertInto(table string, body map[string]any) error
	// SelectWhere returns the values with the next query inforamtion:
	// table: string with the table's name where the data is stored
	// columns: string slice with the columns to get retrived.
	// To get all columns Should pass []string{"*"}
	// where: string with the where clauses
	//
	// it returns the information in a maps slice format ([]map[string]any).
	// If there is an error it returns it as second parameter
	// For example:
	// SelectWhere("product", []string{"name"},"arrivalDay='2023-04-23'")
	// returns the names of the products that arrives the 23 of april in 2023
	SelectWhere(table string, columns []string, where string) ([]map[string]any, error)
	// Select works exactly the same as SelectWhere, but without any where clauses.
	// For example:
	// Select("product", []string{"name"})
	// returns the names of all the products
	Select(table string, columns []string) ([]map[string]any, error)
	// Ping verifies if the connection to the databse is still
	// alive. It returns an error in case the databse is not alive
	// and nill otherwise
	Ping() error
}
