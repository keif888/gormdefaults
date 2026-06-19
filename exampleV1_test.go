package gormdefaultsv1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//*****************************************************************************
// The following tests show that
// Create does not apply ddl level defaults.
// Save on top of an existing record sets all columns in the struct.
// Save without a PK, does not apply ddl level defaults.
// Create using Omit for the ddl level columns allows the db to apply ddl level defaults.
// Updates from a struct does not apply zero/empty fields
// UpdateColumns from a struct does not apply zero/empty fields
// Updates from a map does apply zero/empty fields
// UpdateColumns from a map does apply zero/empty fields

func TestExampleV1(t *testing.T) {
	r0 := Itemv1{
		ID:               1000,
		NameNoDefault:    "Test Record 0 Reset ID",
		InStockNoDefault: 1770,
	}
	assert.NoError(t, r0.Create())

	t.Run("CreateNonDefaults", func(t *testing.T) {
		//*******************************************************************************************************
		// Create a new record with only non default columns populated.
		// Show that gorm annotation defaults are applied,
		// and show that the ddl level defaults are not applied.
		//*******************************************************************************************************

		r1 := Itemv1{
			NameNoDefault:    "Test Record 1 CreateNonDefaults",
			InStockNoDefault: 1771,
		}
		// INSERT INTO "itemv1" ("name_no_default","name_db_default","in_stock_no_default","in_stock_db_default","created_at","updated_at","deleted_at") VALUES ('Test Record 1 CreateNonDefaults','',1771,0,'2026-06-19 20:19:13','2026-06-19 20:19:13',NULL)
		// SELECT "name_blank_default", "name_text_default", "in_stock_zero_default", "in_stock_number_default" FROM "itemv1"  WHERE (id = 1)
		if !assert.NoError(t, dbV1.Create(&r1).Error) {
			return
		} else if assert.GreaterOrEqual(t, r1.ID, uint(1)) {
			assert.Equal(t, "Test Record 1 CreateNonDefaults", r1.NameNoDefault)
			assert.Equal(t, "", r1.NameBlankDefault)
			assert.Equal(t, "NaN", r1.NameTextDefault)
			assert.Equal(t, "", r1.NameDBDefault) // The database ddl says this should be NaS
			assert.Equal(t, 1771, r1.InStockNoDefault)
			assert.Equal(t, -1, r1.InStockNumberDefault)
			assert.Equal(t, 0, r1.InStockZeroDefault)
			assert.Equal(t, 0, r1.InStockDBDefault) // The database ddl says this should be -1!
		}
	})

	t.Run("CreateNonDefaultsOmitDBDefaults", func(t *testing.T) {
		//*******************************************************************************************************
		// Create a new record with only non default columns populated.
		// Omit the database default columns from the create
		// And the requery to prove that database defaults worked.
		//*******************************************************************************************************

		r2 := Itemv1{
			NameNoDefault:    "Test Record 2 CreateNonDefaultsOmitDBDefaults",
			InStockNoDefault: 1772,
		}

		// INSERT INTO "itemv1" ("name_no_default","in_stock_no_default","created_at","updated_at","deleted_at") VALUES ('Test Record 2 CreateNonDefaultsOmitDBDefaults',1772,'2026-06-19 20:27:54','2026-06-19 20:27:54',NULL)
		// SELECT "name_blank_default", "name_text_default", "in_stock_zero_default", "in_stock_number_default" FROM "itemv1"  WHERE (id = 2)
		if !assert.NoError(t, dbV1.Omit("name_db_default", "in_stock_db_default").Create(&r2).Error) {
			return
		} else if assert.GreaterOrEqual(t, r2.ID, uint(1)) {
			// The insert doesn't select columns the omitted columns, so they are 0 or empty
			assert.Equal(t, "Test Record 2 CreateNonDefaultsOmitDBDefaults", r2.NameNoDefault)
			assert.Equal(t, "", r2.NameBlankDefault)
			assert.Equal(t, "NaN", r2.NameTextDefault)
			assert.Equal(t, "", r2.NameDBDefault) // The database ddl says this should be NaS
			assert.Equal(t, 1772, r2.InStockNoDefault)
			assert.Equal(t, -1, r2.InStockNumberDefault)
			assert.Equal(t, 0, r2.InStockZeroDefault)
			assert.Equal(t, 0, r2.InStockDBDefault) // The database ddl says this should be -1!

			r2 = Itemv1{ID: r2.ID}
			// SELECT * FROM "itemv1"  WHERE "itemv1"."deleted_at" IS NULL AND "itemv1"."id" = 2 ORDER BY "itemv1"."id" ASC LIMIT 1
			assert.NoError(t, dbV1.First(&r2).Error)
			assert.Equal(t, "Test Record 2 CreateNonDefaultsOmitDBDefaults", r2.NameNoDefault)
			assert.Equal(t, "", r2.NameBlankDefault)
			assert.Equal(t, "NaN", r2.NameTextDefault)
			assert.Equal(t, "NaS", r2.NameDBDefault) // The database ddl says this should be NaS
			assert.Equal(t, 1772, r2.InStockNoDefault)
			assert.Equal(t, -1, r2.InStockNumberDefault)
			assert.Equal(t, 0, r2.InStockZeroDefault)
			assert.Equal(t, -1, r2.InStockDBDefault) // The database ddl says this should be -1!
		}
	})

	t.Run("SaveAllToZero", func(t *testing.T) {
		//*******************************************************************************************************
		// Create a new record with all columns populated.
		// Save an update to the record with everything set to empty string or 0
		// This shows that Save applies empty string and 0.
		//*******************************************************************************************************

		r3 := Itemv1{
			NameNoDefault:        "Test Record 3 SaveAllToZero",
			NameBlankDefault:     "Test Record 3 SaveAllToZero",
			NameTextDefault:      "Test Record 3 SaveAllToZero",
			NameDBDefault:        "Test Record 3 SaveAllToZero",
			InStockNoDefault:     1773,
			InStockNumberDefault: 1773,
			InStockZeroDefault:   1773,
			InStockDBDefault:     1773,
		}

		// INSERT INTO "itemv1" ("name_no_default","name_blank_default","name_text_default","name_db_default","in_stock_no_default","in_stock_zero_default","in_stock_number_default","in_stock_db_default","created_at","updated_at","deleted_at") VALUES ('Test Record 3 SaveAllToZero','Test Record 3 SaveAllToZero','Test Record 3 SaveAllToZero','Test Record 3 SaveAllToZero',1773,1773,1773,1773,'2026-06-19 21:02:14','2026-06-19 21:02:14',NULL)
		if !assert.NoError(t, dbV1.Create(&r3).Error) {
			return
		} else if assert.GreaterOrEqual(t, r3.ID, uint(1)) {
			// The insert doesn't select columns the omitted columns, so they are 0 or empty
			assert.Equal(t, "Test Record 3 SaveAllToZero", r3.NameNoDefault)
			assert.Equal(t, "Test Record 3 SaveAllToZero", r3.NameBlankDefault)
			assert.Equal(t, "Test Record 3 SaveAllToZero", r3.NameTextDefault)
			assert.Equal(t, "Test Record 3 SaveAllToZero", r3.NameDBDefault)
			assert.Equal(t, 1773, r3.InStockNoDefault)
			assert.Equal(t, 1773, r3.InStockNumberDefault)
			assert.Equal(t, 1773, r3.InStockZeroDefault)
			assert.Equal(t, 1773, r3.InStockDBDefault)

			r3.NameNoDefault = ""
			r3.NameBlankDefault = ""
			r3.NameTextDefault = ""
			r3.NameDBDefault = ""
			r3.InStockNoDefault = 0
			r3.InStockNumberDefault = 0
			r3.InStockZeroDefault = 0
			r3.InStockDBDefault = 0

			// UPDATE "itemv1" SET "name_no_default" = '', "name_blank_default" = '', "name_text_default" = '', "name_db_default" = '', "in_stock_no_default" = 0, "in_stock_zero_default" = 0, "in_stock_number_default" = 0, "in_stock_db_default" = 0, "created_at" = '2026-06-19 21:02:14', "updated_at" = '2026-06-19 21:02:14', "deleted_at" = NULL  WHERE "itemv1"."deleted_at" IS NULL AND "itemv1"."id" = 3
			assert.NoError(t, r3.Save())
			// Reload just in case
			r3 = Itemv1{ID: r3.ID}
			// SELECT * FROM "itemv1"  WHERE "itemv1"."deleted_at" IS NULL AND "itemv1"."id" = 3 ORDER BY "itemv1"."id" ASC LIMIT 1
			assert.NoError(t, dbV1.First(&r3).Error)
			assert.Equal(t, "", r3.NameNoDefault)
			assert.Equal(t, "", r3.NameBlankDefault)
			assert.Equal(t, "", r3.NameTextDefault)
			assert.Equal(t, "", r3.NameDBDefault)
			assert.Equal(t, 0, r3.InStockNoDefault)
			assert.Equal(t, 0, r3.InStockNumberDefault)
			assert.Equal(t, 0, r3.InStockZeroDefault)
			assert.Equal(t, 0, r3.InStockDBDefault)
		}
	})

	t.Run("UpdatesAllToZero", func(t *testing.T) {
		//*******************************************************************************************************
		// Create a new record with all columns populated.
		// Update the record with everything set to empty string or 0
		// This proves that Updates ignores trying to set to "" or 0
		//*******************************************************************************************************

		r4 := Itemv1{
			NameNoDefault:        "Test Record 4 UpdatesAllToZero",
			NameBlankDefault:     "Test Record 4 UpdatesAllToZero",
			NameTextDefault:      "Test Record 4 UpdatesAllToZero",
			NameDBDefault:        "Test Record 4 UpdatesAllToZero",
			InStockNoDefault:     1774,
			InStockNumberDefault: 1774,
			InStockZeroDefault:   1774,
			InStockDBDefault:     1774,
		}

		// INSERT INTO "itemv1" ("name_no_default","name_blank_default","name_text_default","name_db_default","in_stock_no_default","in_stock_zero_default","in_stock_number_default","in_stock_db_default","created_at","updated_at","deleted_at") VALUES ('Test Record 4 UpdatesAllToZero','Test Record 4 UpdatesAllToZero','Test Record 4 UpdatesAllToZero','Test Record 4 UpdatesAllToZero',1774,1774,1774,1774,'2026-06-19 21:02:14','2026-06-19 21:02:14',NULL)
		if !assert.NoError(t, dbV1.Create(&r4).Error) {
			return
		} else if assert.GreaterOrEqual(t, r4.ID, uint(1)) {
			r4.NameNoDefault = ""
			r4.NameBlankDefault = ""
			r4.NameTextDefault = ""
			r4.NameDBDefault = ""
			r4.InStockNoDefault = 0
			r4.InStockNumberDefault = 0
			r4.InStockZeroDefault = 0
			r4.InStockDBDefault = 0

			// UPDATE "itemv1" SET "created_at" = '2026-06-19 21:02:14', "id" = 4, "updated_at" = '2026-06-19 21:02:14'  WHERE "itemv1"."deleted_at" IS NULL AND "itemv1"."id" = 4
			assert.NoError(t, dbV1.Model(&Itemv1{}).Updates(&r4).Error)
			// Reload just in case
			r4 = Itemv1{ID: r4.ID}
			// SELECT * FROM "itemv1"  WHERE "itemv1"."deleted_at" IS NULL AND "itemv1"."id" = 4 ORDER BY "itemv1"."id" ASC LIMIT 1
			assert.NoError(t, dbV1.First(&r4).Error)
			assert.NotEqual(t, "", r4.NameNoDefault)
			assert.NotEqual(t, "", r4.NameBlankDefault)
			assert.NotEqual(t, "", r4.NameTextDefault)
			assert.NotEqual(t, "", r4.NameDBDefault)
			assert.NotEqual(t, 0, r4.InStockNoDefault)
			assert.NotEqual(t, 0, r4.InStockNumberDefault)
			assert.NotEqual(t, 0, r4.InStockZeroDefault)
			assert.NotEqual(t, 0, r4.InStockDBDefault)
		}
	})

	t.Run("UpdateColumnsAllToZero", func(t *testing.T) {
		//*******************************************************************************************************
		// Create a new record with all columns populated.
		// Update the record with everything set to empty string or 0
		// This proves that UpdateColumns ignores trying to set to "" or 0
		//*******************************************************************************************************

		r5 := Itemv1{
			NameNoDefault:        "Test Record 5 UpdateColumnsAllToZero",
			NameBlankDefault:     "Test Record 5 UpdateColumnsAllToZero",
			NameTextDefault:      "Test Record 5 UpdateColumnsAllToZero",
			NameDBDefault:        "Test Record 5 UpdateColumnsAllToZero",
			InStockNoDefault:     1775,
			InStockNumberDefault: 1775,
			InStockZeroDefault:   1775,
			InStockDBDefault:     1775,
		}

		// INSERT INTO "itemv1" ("name_no_default","name_blank_default","name_text_default","name_db_default","in_stock_no_default","in_stock_zero_default","in_stock_number_default","in_stock_db_default","created_at","updated_at","deleted_at") VALUES ('Test Record 5 UpdateColumnsAllToZero','Test Record 5 UpdateColumnsAllToZero','Test Record 5 UpdateColumnsAllToZero','Test Record 5 UpdateColumnsAllToZero',1775,1775,1775,1775,'2026-06-19 21:12:34','2026-06-19 21:12:34',NULL)
		if !assert.NoError(t, dbV1.Create(&r5).Error) {
			return
		} else if assert.GreaterOrEqual(t, r5.ID, uint(1)) {
			r5.NameNoDefault = ""
			r5.NameBlankDefault = ""
			r5.NameTextDefault = ""
			r5.NameDBDefault = ""
			r5.InStockNoDefault = 0
			r5.InStockNumberDefault = 0
			r5.InStockZeroDefault = 0
			r5.InStockDBDefault = 0

			// UPDATE "itemv1" SET "created_at" = '2026-06-19 21:12:34', "id" = 5, "updated_at" = '2026-06-19 21:12:34'  WHERE "itemv1"."deleted_at" IS NULL AND "itemv1"."id" = 5
			assert.NoError(t, dbV1.Model(&Itemv1{}).UpdateColumns(&r5).Error)
			// Reload just in case
			r5 = Itemv1{ID: r5.ID}
			// SELECT * FROM "itemv1"  WHERE "itemv1"."deleted_at" IS NULL AND "itemv1"."id" = 5 ORDER BY "itemv1"."id" ASC LIMIT 1
			assert.NoError(t, dbV1.First(&r5).Error)
			assert.NotEqual(t, "", r5.NameNoDefault)
			assert.NotEqual(t, "", r5.NameBlankDefault)
			assert.NotEqual(t, "", r5.NameTextDefault)
			assert.NotEqual(t, "", r5.NameDBDefault)
			assert.NotEqual(t, 0, r5.InStockNoDefault)
			assert.NotEqual(t, 0, r5.InStockNumberDefault)
			assert.NotEqual(t, 0, r5.InStockZeroDefault)
			assert.NotEqual(t, 0, r5.InStockDBDefault)
		}
	})

	t.Run("UpdatesAllToZeroMap", func(t *testing.T) {
		//*******************************************************************************************************
		// Create a new record with all columns populated.
		// Update the record with everything set to empty string or 0 via a Map
		//*******************************************************************************************************

		r6 := Itemv1{
			NameNoDefault:        "Test Record 6 UpdatesAllToZeroMap",
			NameBlankDefault:     "Test Record 6 UpdatesAllToZeroMap",
			NameTextDefault:      "Test Record 6 UpdatesAllToZeroMap",
			NameDBDefault:        "Test Record 6 UpdatesAllToZeroMap",
			InStockNoDefault:     1776,
			InStockNumberDefault: 1776,
			InStockZeroDefault:   1776,
			InStockDBDefault:     1776,
		}

		// INSERT INTO "itemv1" ("name_no_default","name_blank_default","name_text_default","name_db_default","in_stock_no_default","in_stock_zero_default","in_stock_number_default","in_stock_db_default","created_at","updated_at","deleted_at") VALUES ('Test Record 6 UpdatesAllToZeroMap','Test Record 6 UpdatesAllToZeroMap','Test Record 6 UpdatesAllToZeroMap','Test Record 6 UpdatesAllToZeroMap',1776,1776,1776,1776,'2026-06-19 21:02:16','2026-06-19 21:02:16',NULL)
		if !assert.NoError(t, dbV1.Create(&r6).Error) {
			return
		} else if assert.GreaterOrEqual(t, r6.ID, uint(1)) {
			r6.NameNoDefault = ""
			r6.NameBlankDefault = ""
			r6.NameTextDefault = ""
			r6.NameDBDefault = ""
			r6.InStockNoDefault = 0
			r6.InStockNumberDefault = 0
			r6.InStockZeroDefault = 0
			r6.InStockDBDefault = 0

			// UPDATE "itemv1" SET "in_stock_db_default" = 0, "in_stock_no_default" = 0, "in_stock_number_default" = 0, "in_stock_zero_default" = 0, "name_blank_default" = '', "name_db_default" = '', "name_no_default" = '', "name_text_default" = '', "updated_at" = '2026-06-19 21:27:42'  WHERE "itemv1"."deleted_at" IS NULL AND "itemv1"."id" = 6
			assert.NoError(t, dbV1.Model(&r6).Updates(map[string]any{"name_no_default": "", "name_blank_default": "", "name_text_default": "", "name_db_default": "", "in_stock_no_default": 0, "in_stock_zero_default": 0, "in_stock_number_default": 0, "in_stock_db_default": 0}).Error)
			// Reload just in case
			r6 = Itemv1{ID: r6.ID}
			// SELECT * FROM "itemv1"  WHERE "itemv1"."deleted_at" IS NULL AND "itemv1"."id" = 6 ORDER BY "itemv1"."id" ASC LIMIT 1
			assert.NoError(t, dbV1.First(&r6).Error)
			assert.Equal(t, "", r6.NameNoDefault)
			assert.Equal(t, "", r6.NameBlankDefault)
			assert.Equal(t, "", r6.NameTextDefault)
			assert.Equal(t, "", r6.NameDBDefault)
			assert.Equal(t, 0, r6.InStockNoDefault)
			assert.Equal(t, 0, r6.InStockNumberDefault)
			assert.Equal(t, 0, r6.InStockZeroDefault)
			assert.Equal(t, 0, r6.InStockDBDefault)
		}
	})

	t.Run("UpdateColumnsAllToZeroMap", func(t *testing.T) {
		//*******************************************************************************************************
		// Create a new record with all columns populated.
		// Update the record with everything set to empty string or 0 via map
		//*******************************************************************************************************

		r7 := Itemv1{
			NameNoDefault:        "Test Record 7 UpdateColumnsAllToZeroMap",
			NameBlankDefault:     "Test Record 7 UpdateColumnsAllToZeroMap",
			NameTextDefault:      "Test Record 7 UpdateColumnsAllToZeroMap",
			NameDBDefault:        "Test Record 7 UpdateColumnsAllToZeroMap",
			InStockNoDefault:     1777,
			InStockNumberDefault: 1777,
			InStockZeroDefault:   1777,
			InStockDBDefault:     1777,
		}

		// INSERT INTO "itemv1" ("name_no_default","name_blank_default","name_text_default","name_db_default","in_stock_no_default","in_stock_zero_default","in_stock_number_default","in_stock_db_default","created_at","updated_at","deleted_at") VALUES ('Test Record 7 UpdateColumnsAllToZeroMap','Test Record 7 UpdateColumnsAllToZeroMap','Test Record 7 UpdateColumnsAllToZeroMap','Test Record 7 UpdateColumnsAllToZeroMap',1777,1777,1777,1777,'2026-06-19 21:12:34','2026-06-19 21:12:34',NULL)
		if !assert.NoError(t, dbV1.Create(&r7).Error) {
			return
		} else if assert.GreaterOrEqual(t, r7.ID, uint(1)) {
			r7.NameNoDefault = ""
			r7.NameBlankDefault = ""
			r7.NameTextDefault = ""
			r7.NameDBDefault = ""
			r7.InStockNoDefault = 0
			r7.InStockNumberDefault = 0
			r7.InStockZeroDefault = 0
			r7.InStockDBDefault = 0

			// UPDATE "itemv1" SET "in_stock_db_default" = 0, "in_stock_no_default" = 0, "in_stock_number_default" = 0, "in_stock_zero_default" = 0, "name_blank_default" = '', "name_db_default" = '', "name_no_default" = '', "name_text_default" = ''  WHERE "itemv1"."deleted_at" IS NULL AND "itemv1"."id" = 7
			assert.NoError(t, dbV1.Model(&r7).UpdateColumns(map[string]any{"name_no_default": r7.NameNoDefault, "name_blank_default": r7.NameBlankDefault, "name_text_default": r7.NameTextDefault, "name_db_default": r7.NameDBDefault, "in_stock_no_default": r7.InStockNoDefault, "in_stock_zero_default": r7.InStockZeroDefault, "in_stock_number_default": r7.InStockNumberDefault, "in_stock_db_default": r7.InStockDBDefault}).Error)
			// Reload just in case
			r7 = Itemv1{ID: r7.ID}
			// SELECT * FROM "itemv1"  WHERE "itemv1"."deleted_at" IS NULL AND "itemv1"."id" = 7 ORDER BY "itemv1"."id" ASC LIMIT 1
			assert.NoError(t, dbV1.First(&r7).Error)
			assert.Equal(t, "", r7.NameNoDefault)
			assert.Equal(t, "", r7.NameBlankDefault)
			assert.Equal(t, "", r7.NameTextDefault)
			assert.Equal(t, "", r7.NameDBDefault)
			assert.Equal(t, 0, r7.InStockNoDefault)
			assert.Equal(t, 0, r7.InStockNumberDefault)
			assert.Equal(t, 0, r7.InStockZeroDefault)
			assert.Equal(t, 0, r7.InStockDBDefault)
		}
	})

	t.Run("SaveAllWithZeroThenZeroAgain", func(t *testing.T) {
		//*******************************************************************************************************
		// Create a new record with all columns populated.
		// Save an update to the record with everything set to empty string or 0
		// This shows that Save applies empty string and 0.
		//*******************************************************************************************************

		r8 := Itemv1{
			NameNoDefault:        "",
			NameBlankDefault:     "",
			NameTextDefault:      "",
			NameDBDefault:        "",
			InStockNoDefault:     0,
			InStockNumberDefault: 0,
			InStockZeroDefault:   0,
			InStockDBDefault:     0,
		}

		// INSERT INTO "itemv1" ("name_no_default","name_db_default","in_stock_no_default","in_stock_db_default","created_at","updated_at","deleted_at") VALUES ('','',0,0,'2026-06-19 21:47:29','2026-06-19 21:47:29',NULL)
		// SELECT "name_blank_default", "name_text_default", "in_stock_zero_default", "in_stock_number_default" FROM "itemv1"  WHERE (id = 8)
		if !assert.NoError(t, dbV1.Save(&r8).Error) {
			return
		} else if assert.GreaterOrEqual(t, r8.ID, uint(1)) {
			// The insert doesn't select columns the omitted columns, so they are 0 or empty
			assert.Equal(t, "", r8.NameNoDefault)
			assert.Equal(t, "", r8.NameBlankDefault)   // <-- Gets Gorm Default
			assert.NotEqual(t, "", r8.NameTextDefault) // <-- Gets Gorm Default
			assert.Equal(t, "", r8.NameDBDefault)
			assert.Equal(t, 0, r8.InStockNoDefault)
			assert.NotEqual(t, 0, r8.InStockNumberDefault) // <-- Gets Gorm Default
			assert.Equal(t, 0, r8.InStockZeroDefault)      // <-- Gets Gorm Default
			assert.Equal(t, 0, r8.InStockDBDefault)

			r8.NameNoDefault = ""
			r8.NameBlankDefault = ""
			r8.NameTextDefault = ""
			r8.NameDBDefault = ""
			r8.InStockNoDefault = 0
			r8.InStockNumberDefault = 0
			r8.InStockZeroDefault = 0
			r8.InStockDBDefault = 0

			// UPDATE "itemv1" SET "name_no_default" = '', "name_blank_default" = '', "name_text_default" = '', "name_db_default" = '', "in_stock_no_default" = 0, "in_stock_zero_default" = 0, "in_stock_number_default" = 0, "in_stock_db_default" = 0, "created_at" = '2026-06-19 21:48:47', "updated_at" = '2026-06-19 21:48:47', "deleted_at" = NULL  WHERE "itemv1"."deleted_at" IS NULL AND "itemv1"."id" = 8
			assert.NoError(t, r8.Save())
			// Reload just in case
			r8 = Itemv1{ID: r8.ID}
			// SELECT * FROM "itemv1"  WHERE "itemv1"."deleted_at" IS NULL AND "itemv1"."id" = 8 ORDER BY "itemv1"."id" ASC LIMIT 1
			assert.NoError(t, dbV1.First(&r8).Error)
			assert.Equal(t, "", r8.NameNoDefault)
			assert.Equal(t, "", r8.NameBlankDefault)
			assert.Equal(t, "", r8.NameTextDefault)
			assert.Equal(t, "", r8.NameDBDefault)
			assert.Equal(t, 0, r8.InStockNoDefault)
			assert.Equal(t, 0, r8.InStockNumberDefault)
			assert.Equal(t, 0, r8.InStockZeroDefault)
			assert.Equal(t, 0, r8.InStockDBDefault)
		}
	})
}
