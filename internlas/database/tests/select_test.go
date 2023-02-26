package mysqltests

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/lucastomic/syn-auth/internlas/comparator"
)

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
		"caliber",
		map[string]any{
			"name":    "4-6kg",
			"amount":  12,
			"price":   14.5,
			"weight":  8.0,
			"product": "Merluza",
		},
	},
	{
		"sale",
		map[string]any{
			"date":          "2022-11-18 00:00:00",
			"payedAtMoment": true,
			"amount":        12,
			"userID":        1,
			"product":       "Merluza",
			"caliberName":   "1-3kg",
		},
	},
}

// Set of tests for InsertInto function which should return no error
var selectTests = []struct {
	table    string
	columns  []string
	expected []map[string]any
}{
	{
		table:   "product",
		columns: []string{"name", "arrivalDate"},
		expected: []map[string]any{
			{
				"name":        "Merluza",
				"arrivalDate": "2023-02-17 00:00:00",
			},
		},
	},
	{
		table:   "user",
		columns: []string{"*"},
		expected: []map[string]any{
			{
				"id":        1,
				"business":  nil,
				"name":      "Lucas Tomic",
				"nif":       "12345678Z",
				"phone":     "623029222",
				"ubication": "Madrid, Spain",
			},
		},
	},
	{
		table:   "caliber",
		columns: []string{"name", "amount"},
		expected: []map[string]any{
			{
				"name":   "1-3kg",
				"amount": 25,
			},
			{
				"name":   "4-6kg",
				"amount": 12,
			},
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
	for _, tt := range insertRowsForSelectTesting {
		db.InsertInto(tt.table, tt.body)
	}

	for _, tt := range selectTests {
		t.Run(tt.table, func(t *testing.T) {
			// t.Parallel()
			res, err := db.Select(tt.table, tt.columns)
			if err != nil {
				t.Errorf("Error selecting the values from the database: %v", err)
			}
			if !comparator.CompareMapSlices(res, tt.expected) {
				t.Errorf("Incorrect value retrived. Expected: %v, Got: %v", tt.expected, res)
			}
		})
	}
}
