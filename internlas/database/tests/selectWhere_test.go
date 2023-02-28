package mysqltests

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/lucastomic/syn-auth/internlas/comparator"
)

// Set of rows to insert for test the select function
var insertRowsForSelectWhereTesting = []struct {
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
var selectWhereTests = []struct {
	table    string
	columns  []string
	where    string
	expected []map[string]any
}{
	{
		table:   "caliber",
		columns: []string{"name", "amount"},
		where:   "amount = 25",
		expected: []map[string]any{
			{
				"name":   "1-3kg",
				"amount": 25,
			},
		},
	},
	{
		table:   "caliber",
		columns: []string{"name", "amount"},
		where:   "amount = 12",
		expected: []map[string]any{
			{
				"name":   "4-6kg",
				"amount": 12,
			},
		},
	},
	{
		table:   "product",
		columns: []string{"name"},
		where:   "name = besugo",
		expected: []map[string]any{
			{},
		},
	},
	{
		table:   "caliber",
		columns: []string{"name", "amount"},
		where:   "name = Merluza AND amount = 12",
		expected: []map[string]any{
			{
				"name":   "4-6kg",
				"amount": 12,
			},
		},
	},
}

func TestSelectWhere(t *testing.T) {
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
