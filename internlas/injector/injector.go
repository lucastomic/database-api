// injector is the package in charge of inject all the dependencies
package injector

import "github.com/lucastomic/syn-auth/internlas/database"

// GetDatabase returns the database used in the app, without knowing the concret implementation
func GetDatabase() (database.Database, error) {
	return database.GetMYSQLDB()
}
