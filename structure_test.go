package gormdefaultsv1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestValidateStructureV1 makes sure that the table definition has the expected data types and defaults.
func TestValidateStructureV1(t *testing.T) {
	type TableInfo struct {
		CID       int `gorm:"Column:cid"`
		Name      string
		Type      string
		Notnull   bool
		DfltValue *string
		Pk        bool `gorm:"Column:pk"`
	}

	expected := []TableInfo{
		{
			CID:     0,
			Name:    "id",
			Type:    "INTEGER",
			Notnull: false,
			Pk:      true,
		},
		{
			CID:     1,
			Name:    "name_no_default",
			Type:    "bytes",
			Notnull: false,
			Pk:      false,
		},
		{
			CID:       2,
			Name:      "name_blank_default",
			Type:      "bytes",
			Notnull:   false,
			DfltValue: new(`''`),
			Pk:        false,
		},
		{
			CID:       3,
			Name:      "name_text_default",
			Type:      "bytes",
			Notnull:   false,
			DfltValue: new(`'NaN'`),
			Pk:        false,
		},
		{
			CID:     4,
			Name:    "in_stock_no_default",
			Type:    "INTEGER",
			Notnull: false,
			Pk:      false,
		},
		{
			CID:       5,
			Name:      "in_stock_zero_default",
			Type:      "INTEGER",
			Notnull:   false,
			DfltValue: new(`0`),
			Pk:        false,
		},
		{
			CID:       6,
			Name:      "in_stock_number_default",
			Type:      "INTEGER",
			Notnull:   false,
			DfltValue: new(`-1`),
			Pk:        false,
		},
		{
			CID:     7,
			Name:    "created_at",
			Type:    "datetime",
			Notnull: false,
			Pk:      false,
		},
		{
			CID:     8,
			Name:    "updated_at",
			Type:    "datetime",
			Notnull: false,
			Pk:      false,
		},
		{
			CID:     9,
			Name:    "deleted_at",
			Type:    "datetime",
			Notnull: false,
			Pk:      false,
		},
		{
			CID:       10,
			Name:      "name_db_default",
			Type:      "bytes",
			Notnull:   false,
			DfltValue: new(`'NaS'`),
			Pk:        false,
		},
		{
			CID:       11,
			Name:      "in_stock_db_default",
			Type:      "INTEGER",
			Notnull:   false,
			DfltValue: new(`-1`),
			Pk:        false,
		},
	}
	actual := []TableInfo{}

	result := dbV1.Raw("PRAGMA table_info(itemv1);").Scan(&actual)
	if assert.NoError(t, result.Error) {
		assert.Equal(t, int64(12), result.RowsAffected)
		assert.ElementsMatch(t, expected, actual)
	}
}

// TestValidateStructureV2 makes sure that the table definition has the expected data types and defaults.
func TestValidateStructureV2(t *testing.T) {
	type TableInfo struct {
		CID       int `gorm:"Column:cid"`
		Name      string
		Type      string
		Notnull   bool
		DfltValue *string
		Pk        bool `gorm:"Column:pk"`
	}

	expected := []TableInfo{
		{
			CID:     0,
			Name:    "id",
			Type:    "INTEGER",
			Notnull: false,
			Pk:      true,
		},
		{
			CID:     1,
			Name:    "name_no_default",
			Type:    "BLOB",
			Notnull: false,
			Pk:      false,
		},
		{
			CID:       2,
			Name:      "name_blank_default",
			Type:      "BLOB",
			Notnull:   false,
			DfltValue: new(`""`),
			Pk:        false,
		},
		{
			CID:       3,
			Name:      "name_text_default",
			Type:      "BLOB",
			Notnull:   false,
			DfltValue: new(`"NaN"`),
			Pk:        false,
		},
		{
			CID:     4,
			Name:    "in_stock_no_default",
			Type:    "INTEGER",
			Notnull: false,
			Pk:      false,
		},
		{
			CID:       5,
			Name:      "in_stock_zero_default",
			Type:      "INTEGER",
			Notnull:   false,
			DfltValue: new(`0`),
			Pk:        false,
		},
		{
			CID:       6,
			Name:      "in_stock_number_default",
			Type:      "INTEGER",
			Notnull:   false,
			DfltValue: new(`-1`),
			Pk:        false,
		},
		{
			CID:     7,
			Name:    "created_at",
			Type:    "datetime",
			Notnull: false,
			Pk:      false,
		},
		{
			CID:     8,
			Name:    "updated_at",
			Type:    "datetime",
			Notnull: false,
			Pk:      false,
		},
		{
			CID:     9,
			Name:    "deleted_at",
			Type:    "datetime",
			Notnull: false,
			Pk:      false,
		},
		{
			CID:       10,
			Name:      "name_db_default",
			Type:      "BLOB",
			Notnull:   false,
			DfltValue: new(`"NaS"`),
			Pk:        false,
		},
		{
			CID:       11,
			Name:      "in_stock_db_default",
			Type:      "INTEGER",
			Notnull:   false,
			DfltValue: new(`-1`),
			Pk:        false,
		},
	}
	actual := []TableInfo{}

	result := dbV2.Raw("PRAGMA table_info(itemv2);").Scan(&actual)
	if assert.NoError(t, result.Error) {
		assert.Equal(t, int64(12), result.RowsAffected)
		assert.ElementsMatch(t, expected, actual)
	}
}

// TestValidateStructureV3 makes sure that the table definition has the expected data types and defaults.
func TestValidateStructureV3(t *testing.T) {
	type TableInfo struct {
		CID       int `gorm:"Column:cid"`
		Name      string
		Type      string
		Notnull   bool
		DfltValue *string
		Pk        bool `gorm:"Column:pk"`
	}

	expected := []TableInfo{
		{
			CID:     0,
			Name:    "id",
			Type:    "INTEGER",
			Notnull: false,
			Pk:      true,
		},
		{
			CID:     1,
			Name:    "name_no_default",
			Type:    "BLOB",
			Notnull: false,
			Pk:      false,
		},
		{
			CID:       2,
			Name:      "name_blank_default",
			Type:      "BLOB",
			Notnull:   false,
			DfltValue: new(`""`),
			Pk:        false,
		},
		{
			CID:       3,
			Name:      "name_text_default",
			Type:      "BLOB",
			Notnull:   false,
			DfltValue: new(`"NaN"`),
			Pk:        false,
		},
		{
			CID:     4,
			Name:    "in_stock_no_default",
			Type:    "INTEGER",
			Notnull: false,
			Pk:      false,
		},
		{
			CID:       5,
			Name:      "in_stock_zero_default",
			Type:      "INTEGER",
			Notnull:   false,
			DfltValue: new(`0`),
			Pk:        false,
		},
		{
			CID:       6,
			Name:      "in_stock_number_default",
			Type:      "INTEGER",
			Notnull:   false,
			DfltValue: new(`-1`),
			Pk:        false,
		},
		{
			CID:     7,
			Name:    "created_at",
			Type:    "datetime",
			Notnull: false,
			Pk:      false,
		},
		{
			CID:     8,
			Name:    "updated_at",
			Type:    "datetime",
			Notnull: false,
			Pk:      false,
		},
		{
			CID:     9,
			Name:    "deleted_at",
			Type:    "datetime",
			Notnull: false,
			Pk:      false,
		},
		{
			CID:       10,
			Name:      "name_db_default",
			Type:      "BLOB",
			Notnull:   false,
			DfltValue: new(`"NaS"`),
			Pk:        false,
		},
		{
			CID:       11,
			Name:      "in_stock_db_default",
			Type:      "INTEGER",
			Notnull:   false,
			DfltValue: new(`-1`),
			Pk:        false,
		},
	}
	actual := []TableInfo{}

	result := dbV2.Raw("PRAGMA table_info(itemv3);").Scan(&actual)
	if assert.NoError(t, result.Error) {
		assert.Equal(t, int64(12), result.RowsAffected)
		assert.ElementsMatch(t, expected, actual)
	}
}
