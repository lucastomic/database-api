package mysqltests

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

// TODO:Terminar

// Set of tests for InsertInto function which should return no error
var selectTestsRight = []struct {
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

// Set of tests for InsertInto function which should return error
var selectTestsWrong = []struct {
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

// Set of rows to insert for test the select function
var insertRowsForSelectTesting = []struct {
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

func TestSelect(t *testing.T) {
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
