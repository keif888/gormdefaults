package gormdefaultsv1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//*****************************************************************************
// The following tests show that using a Pointer based type allows zero's and empty strings to be applied
// Create does not apply ddl level defaults.
// Create does allow zero/empty pointer fields, but not other fields
// Create using Omit for the ddl level columns allows the db to apply ddl level defaults.
// Save on top of an existing record sets all columns in the struct.
// Save without a PK, does not apply ddl level defaults.
// Updates from a struct applies zero/empty pointer fields, but not other fields
// UpdateColumns from a struct applies zero/empty pointer fields, but not other fields
// Updates from a map does apply zero/empty fields
// UpdateColumns from a map does apply zero/empty fields

func TestExampleV3(t *testing.T) {
	r0 := Itemv3{
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

		r1 := Itemv3{
			NameNoDefault:    "Test Record 1 CreateNonDefaults",
			InStockNoDefault: 1771,
		}
		// INSERT INTO `itemv3` (`name_no_default`,`name_blank_default`,`name_text_default`,`name_db_default`,`in_stock_no_default`,`in_stock_zero_default`,`in_stock_number_default`,`in_stock_db_default`,`created_at`,`updated_at`,`deleted_at`) VALUES ("Test Record 1 CreateNonDefaults","","NaN","",1771,0,-1,0,"2026-06-19 22:18:34.992","2026-06-19 22:18:34.992",NULL) RETURNING `id`
		if !assert.NoError(t, dbV2.Create(&r1).Error) {
			return
		} else if assert.GreaterOrEqual(t, r1.ID, uint(1)) {
			assert.Equal(t, "Test Record 1 CreateNonDefaults", r1.NameNoDefault)
			assert.Equal(t, "", *r1.NameBlankDefault)
			assert.Equal(t, "NaN", *r1.NameTextDefault)
			assert.Equal(t, "", r1.NameDBDefault) // The database ddl says this should be NaS
			assert.Equal(t, 1771, r1.InStockNoDefault)
			assert.Equal(t, -1, *r1.InStockNumberDefault)
			assert.Equal(t, 0, *r1.InStockZeroDefault)
			assert.Equal(t, 0, r1.InStockDBDefault) // The database ddl says this should be -1!
		}
	})

	t.Run("CreateNonDefaultsOmitDBDefaults", func(t *testing.T) {
		//*******************************************************************************************************
		// Create a new record with only non default columns populated.
		// Omit the database default columns from the create
		// And the requery to prove that database defaults worked.
		//*******************************************************************************************************

		r2 := Itemv3{
			NameNoDefault:    "Test Record 2 CreateNonDefaultsOmitDBDefaults",
			InStockNoDefault: 1772,
		}

		// INSERT INTO `itemv3` (`name_no_default`,`name_blank_default`,`name_text_default`,`in_stock_no_default`,`in_stock_zero_default`,`in_stock_number_default`,`created_at`,`updated_at`,`deleted_at`) VALUES ("Test Record 2 CreateNonDefaultsOmitDBDefaults","","NaN",1772,0,-1,"2026-06-19 22:18:34.995","2026-06-19 22:18:34.995",NULL) RETURNING `id`
		if !assert.NoError(t, dbV2.Omit("name_db_default", "in_stock_db_default").Create(&r2).Error) {
			return
		} else if assert.GreaterOrEqual(t, r2.ID, uint(1)) {
			// The insert doesn't select columns the omitted columns, so they are 0 or empty
			assert.Equal(t, "Test Record 2 CreateNonDefaultsOmitDBDefaults", r2.NameNoDefault)
			assert.Equal(t, "", *r2.NameBlankDefault)
			assert.Equal(t, "NaN", *r2.NameTextDefault)
			assert.Equal(t, "", r2.NameDBDefault) // The database ddl says this should be NaS
			assert.Equal(t, 1772, r2.InStockNoDefault)
			assert.Equal(t, -1, *r2.InStockNumberDefault)
			assert.Equal(t, 0, *r2.InStockZeroDefault)
			assert.Equal(t, 0, r2.InStockDBDefault) // The database ddl says this should be -1!

			r2 = Itemv3{ID: r2.ID}
			// SELECT * FROM `itemv3` WHERE `itemv3`.`id` = 2 ORDER BY `itemv3`.`id` LIMIT 1
			assert.NoError(t, dbV2.First(&r2).Error)
			assert.Equal(t, "Test Record 2 CreateNonDefaultsOmitDBDefaults", r2.NameNoDefault)
			assert.Equal(t, "", *r2.NameBlankDefault)
			assert.Equal(t, "NaN", *r2.NameTextDefault)
			assert.Equal(t, "NaS", r2.NameDBDefault) // The database ddl says this should be NaS
			assert.Equal(t, 1772, r2.InStockNoDefault)
			assert.Equal(t, -1, *r2.InStockNumberDefault)
			assert.Equal(t, 0, *r2.InStockZeroDefault)
			assert.Equal(t, -1, r2.InStockDBDefault) // The database ddl says this should be -1!
		}
	})

	t.Run("SaveAllToZero", func(t *testing.T) {
		//*******************************************************************************************************
		// Create a new record with all columns populated.
		// Save an update to the record with everything set to empty string or 0
		// This shows that Save applies empty string and 0.
		//*******************************************************************************************************

		r3 := Itemv3{
			NameNoDefault:        "Test Record 3 SaveAllToZero",
			NameBlankDefault:     new("Test Record 3 SaveAllToZero"),
			NameTextDefault:      new("Test Record 3 SaveAllToZero"),
			NameDBDefault:        "Test Record 3 SaveAllToZero",
			InStockNoDefault:     1773,
			InStockNumberDefault: new(1773),
			InStockZeroDefault:   new(1773),
			InStockDBDefault:     1773,
		}

		// INSERT INTO `itemv3` (`name_no_default`,`name_blank_default`,`name_text_default`,`name_db_default`,`in_stock_no_default`,`in_stock_zero_default`,`in_stock_number_default`,`in_stock_db_default`,`created_at`,`updated_at`,`deleted_at`) VALUES ("Test Record 3 SaveAllToZero","Test Record 3 SaveAllToZero","Test Record 3 SaveAllToZero","Test Record 3 SaveAllToZero",1773,1773,1773,1773,"2026-06-19 22:18:34.998","2026-06-19 22:18:34.998",NULL) RETURNING `id`
		if !assert.NoError(t, dbV2.Create(&r3).Error) {
			return
		} else if assert.GreaterOrEqual(t, r3.ID, uint(1)) {
			// The insert doesn't select columns the omitted columns, so they are 0 or empty
			assert.Equal(t, "Test Record 3 SaveAllToZero", r3.NameNoDefault)
			assert.Equal(t, "Test Record 3 SaveAllToZero", *r3.NameBlankDefault)
			assert.Equal(t, "Test Record 3 SaveAllToZero", *r3.NameTextDefault)
			assert.Equal(t, "Test Record 3 SaveAllToZero", r3.NameDBDefault)
			assert.Equal(t, 1773, r3.InStockNoDefault)
			assert.Equal(t, 1773, *r3.InStockNumberDefault)
			assert.Equal(t, 1773, *r3.InStockZeroDefault)
			assert.Equal(t, 1773, r3.InStockDBDefault)

			r3.NameNoDefault = ""
			r3.NameBlankDefault = new("")
			r3.NameTextDefault = new("")
			r3.NameDBDefault = ""
			r3.InStockNoDefault = 0
			r3.InStockNumberDefault = new(0)
			r3.InStockZeroDefault = new(0)
			r3.InStockDBDefault = 0

			// UPDATE `itemv3` SET `name_no_default`="",`name_blank_default`="",`name_text_default`="",`name_db_default`="",`in_stock_no_default`=0,`in_stock_zero_default`=0,`in_stock_number_default`=0,`in_stock_db_default`=0,`created_at`="2026-06-19 22:18:34.998",`updated_at`="2026-06-19 22:18:35",`deleted_at`=NULL WHERE `id` = 3
			assert.NoError(t, r3.Save())
			// Reload just in case
			r3 = Itemv3{ID: r3.ID}
			// SELECT * FROM `itemv3` WHERE `itemv3`.`id` = 3 ORDER BY `itemv3`.`id` LIMIT 1
			assert.NoError(t, dbV2.First(&r3).Error)
			assert.Equal(t, "", r3.NameNoDefault)
			assert.Equal(t, "", *r3.NameBlankDefault)
			assert.Equal(t, "", *r3.NameTextDefault)
			assert.Equal(t, "", r3.NameDBDefault)
			assert.Equal(t, 0, r3.InStockNoDefault)
			assert.Equal(t, 0, *r3.InStockNumberDefault)
			assert.Equal(t, 0, *r3.InStockZeroDefault)
			assert.Equal(t, 0, r3.InStockDBDefault)
		}
	})

	t.Run("UpdatesAllToZero", func(t *testing.T) {
		//*******************************************************************************************************
		// Create a new record with all columns populated.
		// Update the record with everything set to empty string or 0
		// This proves that Updates ignores trying to set to "" or 0
		//*******************************************************************************************************

		r4 := Itemv3{
			NameNoDefault:        "Test Record 4 UpdatesAllToZero",
			NameBlankDefault:     new("Test Record 4 UpdatesAllToZero"),
			NameTextDefault:      new("Test Record 4 UpdatesAllToZero"),
			NameDBDefault:        "Test Record 4 UpdatesAllToZero",
			InStockNoDefault:     1774,
			InStockNumberDefault: new(1774),
			InStockZeroDefault:   new(1774),
			InStockDBDefault:     1774,
		}

		// INSERT INTO `itemv3` (`name_no_default`,`name_blank_default`,`name_text_default`,`name_db_default`,`in_stock_no_default`,`in_stock_zero_default`,`in_stock_number_default`,`in_stock_db_default`,`created_at`,`updated_at`,`deleted_at`) VALUES ("Test Record 4 UpdatesAllToZero","Test Record 4 UpdatesAllToZero","Test Record 4 UpdatesAllToZero","Test Record 4 UpdatesAllToZero",1774,1774,1774,1774,"2026-06-19 22:18:35.003","2026-06-19 22:18:35.003",NULL) RETURNING `id`
		if !assert.NoError(t, dbV2.Create(&r4).Error) {
			return
		} else if assert.GreaterOrEqual(t, r4.ID, uint(1)) {
			r4.NameNoDefault = ""
			r4.NameBlankDefault = new("")
			r4.NameTextDefault = new("")
			r4.NameDBDefault = ""
			r4.InStockNoDefault = 0
			r4.InStockNumberDefault = new(0)
			r4.InStockZeroDefault = new(0)
			r4.InStockDBDefault = 0

			// UPDATE `itemv3` SET `id`=4,`name_blank_default`="",`name_text_default`="",`in_stock_zero_default`=0,`in_stock_number_default`=0,`created_at`="2026-06-19 22:18:35.003",`updated_at`="2026-06-19 22:18:35.007" WHERE `id` = 4
			assert.NoError(t, dbV2.Model(&Itemv3{ID: r4.ID}).Updates(&r4).Error)
			// Reload just in case
			r4 = Itemv3{ID: r4.ID}
			// SELECT * FROM `itemv3` WHERE `itemv3`.`id` = 4 ORDER BY `itemv3`.`id` LIMIT 1
			assert.NoError(t, dbV2.First(&r4).Error)
			assert.NotEqual(t, "", r4.NameNoDefault)
			assert.Equal(t, "", *r4.NameBlankDefault) // <-- Pointer Field
			assert.Equal(t, "", *r4.NameTextDefault)  // <-- Pointer Field
			assert.NotEqual(t, "", r4.NameDBDefault)
			assert.NotEqual(t, 0, r4.InStockNoDefault)
			assert.Equal(t, 0, *r4.InStockNumberDefault) // <-- Pointer Field
			assert.Equal(t, 0, *r4.InStockZeroDefault)   // <-- Pointer Field
			assert.NotEqual(t, 0, r4.InStockDBDefault)
		}
	})

	t.Run("UpdateColumnsAllToZero", func(t *testing.T) {
		//*******************************************************************************************************
		// Create a new record with all columns populated.
		// Update the record with everything set to empty string or 0
		// This proves that UpdateColumns ignores trying to set to "" or 0
		//*******************************************************************************************************

		r5 := Itemv3{
			NameNoDefault:        "Test Record 5 UpdateColumnsAllToZero",
			NameBlankDefault:     new("Test Record 5 UpdateColumnsAllToZero"),
			NameTextDefault:      new("Test Record 5 UpdateColumnsAllToZero"),
			NameDBDefault:        "Test Record 5 UpdateColumnsAllToZero",
			InStockNoDefault:     1775,
			InStockNumberDefault: new(1775),
			InStockZeroDefault:   new(1775),
			InStockDBDefault:     1775,
		}

		// INSERT INTO `itemv3` (`name_no_default`,`name_blank_default`,`name_text_default`,`name_db_default`,`in_stock_no_default`,`in_stock_zero_default`,`in_stock_number_default`,`in_stock_db_default`,`created_at`,`updated_at`,`deleted_at`) VALUES ("Test Record 5 UpdateColumnsAllToZero","Test Record 5 UpdateColumnsAllToZero","Test Record 5 UpdateColumnsAllToZero","Test Record 5 UpdateColumnsAllToZero",1775,1775,1775,1775,"2026-06-19 22:18:35.009","2026-06-19 22:18:35.009",NULL) RETURNING `id`
		if !assert.NoError(t, dbV2.Create(&r5).Error) {
			return
		} else if assert.GreaterOrEqual(t, r5.ID, uint(1)) {
			r5.NameNoDefault = ""
			r5.NameBlankDefault = new("")
			r5.NameTextDefault = new("")
			r5.NameDBDefault = ""
			r5.InStockNoDefault = 0
			r5.InStockNumberDefault = new(0)
			r5.InStockZeroDefault = new(0)
			r5.InStockDBDefault = 0

			// UPDATE `itemv3` SET `id`=5,`name_blank_default`="",`name_text_default`="",`in_stock_zero_default`=0,`in_stock_number_default`=0,`created_at`="2026-06-19 22:18:35.009",`updated_at`="2026-06-19 22:18:35.009" WHERE `id` = 5
			assert.NoError(t, dbV2.Model(&Itemv3{ID: r5.ID}).UpdateColumns(&r5).Error)
			// Reload just in case
			r5 = Itemv3{ID: r5.ID}
			// SELECT * FROM `itemv3` WHERE `itemv3`.`id` = 5 ORDER BY `itemv3`.`id` LIMIT 1
			assert.NoError(t, dbV2.First(&r5).Error)
			assert.NotEqual(t, "", r5.NameNoDefault)
			assert.Equal(t, "", *r5.NameBlankDefault) // <-- Pointer Field
			assert.Equal(t, "", *r5.NameTextDefault)  // <-- Pointer Field
			assert.NotEqual(t, "", r5.NameDBDefault)
			assert.NotEqual(t, 0, r5.InStockNoDefault)
			assert.Equal(t, 0, *r5.InStockNumberDefault) // <-- Pointer Field
			assert.Equal(t, 0, *r5.InStockZeroDefault)   // <-- Pointer Field
			assert.NotEqual(t, 0, r5.InStockDBDefault)
		}
	})

	t.Run("UpdatesAllToZeroMap", func(t *testing.T) {
		//*******************************************************************************************************
		// Create a new record with all columns populated.
		// Update the record with everything set to empty string or 0 via a Map
		//*******************************************************************************************************

		r6 := Itemv3{
			NameNoDefault:        "Test Record 6 UpdatesAllToZeroMap",
			NameBlankDefault:     new("Test Record 6 UpdatesAllToZeroMap"),
			NameTextDefault:      new("Test Record 6 UpdatesAllToZeroMap"),
			NameDBDefault:        "Test Record 6 UpdatesAllToZeroMap",
			InStockNoDefault:     1776,
			InStockNumberDefault: new(1776),
			InStockZeroDefault:   new(1776),
			InStockDBDefault:     1776,
		}

		// INSERT INTO `itemv3` (`name_no_default`,`name_blank_default`,`name_text_default`,`name_db_default`,`in_stock_no_default`,`in_stock_zero_default`,`in_stock_number_default`,`in_stock_db_default`,`created_at`,`updated_at`,`deleted_at`) VALUES ("Test Record 6 UpdatesAllToZeroMap","Test Record 6 UpdatesAllToZeroMap","Test Record 6 UpdatesAllToZeroMap","Test Record 6 UpdatesAllToZeroMap",1776,1776,1776,1776,"2026-06-19 22:18:35.016","2026-06-19 22:18:35.016",NULL) RETURNING `id`
		if !assert.NoError(t, dbV2.Create(&r6).Error) {
			return
		} else if assert.GreaterOrEqual(t, r6.ID, uint(1)) {
			r6.NameNoDefault = ""
			r6.NameBlankDefault = new("")
			r6.NameTextDefault = new("")
			r6.NameDBDefault = ""
			r6.InStockNoDefault = 0
			r6.InStockNumberDefault = new(0)
			r6.InStockZeroDefault = new(0)
			r6.InStockDBDefault = 0

			// UPDATE `itemv3` SET `in_stock_db_default`=0,`in_stock_no_default`=0,`in_stock_number_default`=0,`in_stock_zero_default`=0,`name_blank_default`="",`name_db_default`="",`name_no_default`="",`name_text_default`="",`updated_at`="2026-06-19 22:18:35.019" WHERE `id` = 6
			assert.NoError(t, dbV2.Model(&r6).Updates(map[string]any{"name_no_default": "", "name_blank_default": "", "name_text_default": "", "name_db_default": "", "in_stock_no_default": 0, "in_stock_zero_default": 0, "in_stock_number_default": 0, "in_stock_db_default": 0}).Error)
			// Reload just in case
			r6 = Itemv3{ID: r6.ID}
			// SELECT * FROM `itemv3` WHERE `itemv3`.`id` = 6 ORDER BY `itemv3`.`id` LIMIT 1
			assert.NoError(t, dbV2.First(&r6).Error)
			assert.Equal(t, "", r6.NameNoDefault)
			assert.Equal(t, "", *r6.NameBlankDefault)
			assert.Equal(t, "", *r6.NameTextDefault)
			assert.Equal(t, "", r6.NameDBDefault)
			assert.Equal(t, 0, r6.InStockNoDefault)
			assert.Equal(t, 0, *r6.InStockNumberDefault)
			assert.Equal(t, 0, *r6.InStockZeroDefault)
			assert.Equal(t, 0, r6.InStockDBDefault)
		}
	})

	t.Run("UpdateColumnsAllToZeroMap", func(t *testing.T) {
		//*******************************************************************************************************
		// Create a new record with all columns populated.
		// Update the record with everything set to empty string or 0 via map
		//*******************************************************************************************************

		r7 := Itemv3{
			NameNoDefault:        "Test Record 7 UpdateColumnsAllToZeroMap",
			NameBlankDefault:     new("Test Record 7 UpdateColumnsAllToZeroMap"),
			NameTextDefault:      new("Test Record 7 UpdateColumnsAllToZeroMap"),
			NameDBDefault:        "Test Record 7 UpdateColumnsAllToZeroMap",
			InStockNoDefault:     1777,
			InStockNumberDefault: new(1777),
			InStockZeroDefault:   new(1777),
			InStockDBDefault:     1777,
		}

		// INSERT INTO `itemv3` (`name_no_default`,`name_blank_default`,`name_text_default`,`name_db_default`,`in_stock_no_default`,`in_stock_zero_default`,`in_stock_number_default`,`in_stock_db_default`,`created_at`,`updated_at`,`deleted_at`) VALUES ("Test Record 7 UpdateColumnsAllToZeroMap","Test Record 7 UpdateColumnsAllToZeroMap","Test Record 7 UpdateColumnsAllToZeroMap","Test Record 7 UpdateColumnsAllToZeroMap",1777,1777,1777,1777,"2026-06-19 22:18:35.026","2026-06-19 22:18:35.026",NULL) RETURNING `id`
		if !assert.NoError(t, dbV2.Create(&r7).Error) {
			return
		} else if assert.GreaterOrEqual(t, r7.ID, uint(1)) {
			r7.NameNoDefault = ""
			r7.NameBlankDefault = new("")
			r7.NameTextDefault = new("")
			r7.NameDBDefault = ""
			r7.InStockNoDefault = 0
			r7.InStockNumberDefault = new(0)
			r7.InStockZeroDefault = new(0)
			r7.InStockDBDefault = 0

			// UPDATE `itemv3` SET `in_stock_db_default`=0,`in_stock_no_default`=0,`in_stock_number_default`=0,`in_stock_zero_default`=0,`name_blank_default`="",`name_db_default`="",`name_no_default`="",`name_text_default`="" WHERE `id` = 7
			assert.NoError(t, dbV2.Model(&r7).UpdateColumns(map[string]any{"name_no_default": r7.NameNoDefault, "name_blank_default": r7.NameBlankDefault, "name_text_default": r7.NameTextDefault, "name_db_default": r7.NameDBDefault, "in_stock_no_default": r7.InStockNoDefault, "in_stock_zero_default": r7.InStockZeroDefault, "in_stock_number_default": r7.InStockNumberDefault, "in_stock_db_default": r7.InStockDBDefault}).Error)
			// Reload just in case
			r7 = Itemv3{ID: r7.ID}
			// SELECT * FROM `itemv3` WHERE `itemv3`.`id` = 7 ORDER BY `itemv3`.`id` LIMIT 1
			assert.NoError(t, dbV2.First(&r7).Error)
			assert.Equal(t, "", r7.NameNoDefault)
			assert.Equal(t, "", *r7.NameBlankDefault)
			assert.Equal(t, "", *r7.NameTextDefault)
			assert.Equal(t, "", r7.NameDBDefault)
			assert.Equal(t, 0, r7.InStockNoDefault)
			assert.Equal(t, 0, *r7.InStockNumberDefault)
			assert.Equal(t, 0, *r7.InStockZeroDefault)
			assert.Equal(t, 0, r7.InStockDBDefault)
		}
	})

	t.Run("SaveAllWithZeroThenZeroAgain", func(t *testing.T) {
		//*******************************************************************************************************
		// Create a new record with all columns populated.
		// Save an update to the record with everything set to empty string or 0
		// This shows that Save applies empty string and 0.
		//*******************************************************************************************************

		r8 := Itemv3{
			NameNoDefault:        "",
			NameBlankDefault:     new(""),
			NameTextDefault:      new(""),
			NameDBDefault:        "",
			InStockNoDefault:     0,
			InStockNumberDefault: new(0),
			InStockZeroDefault:   new(0),
			InStockDBDefault:     0,
		}

		// INSERT INTO `itemv3` (`name_no_default`,`name_blank_default`,`name_text_default`,`name_db_default`,`in_stock_no_default`,`in_stock_zero_default`,`in_stock_number_default`,`in_stock_db_default`,`created_at`,`updated_at`,`deleted_at`) VALUES ("","","","",0,0,0,0,"2026-06-19 22:18:35.035","2026-06-19 22:18:35.035",NULL) RETURNING `id`
		if !assert.NoError(t, dbV2.Save(&r8).Error) {
			return
		} else if assert.GreaterOrEqual(t, r8.ID, uint(1)) {
			// The insert doesn't select columns the omitted columns, so they are 0 or empty
			assert.Equal(t, "", r8.NameNoDefault)
			assert.Equal(t, "", *r8.NameBlankDefault)
			assert.Equal(t, "", *r8.NameTextDefault)
			assert.Equal(t, "", r8.NameDBDefault)
			assert.Equal(t, 0, r8.InStockNoDefault)
			assert.Equal(t, 0, *r8.InStockNumberDefault)
			assert.Equal(t, 0, *r8.InStockZeroDefault)
			assert.Equal(t, 0, r8.InStockDBDefault)

			r8.NameNoDefault = ""
			r8.NameBlankDefault = new("")
			r8.NameTextDefault = new("")
			r8.NameDBDefault = ""
			r8.InStockNoDefault = 0
			r8.InStockNumberDefault = new(0)
			r8.InStockZeroDefault = new(0)
			r8.InStockDBDefault = 0

			// UPDATE `itemv3` SET `name_no_default`="",`name_blank_default`="",`name_text_default`="",`name_db_default`="",`in_stock_no_default`=0,`in_stock_zero_default`=0,`in_stock_number_default`=0,`in_stock_db_default`=0,`created_at`="2026-06-19 22:18:35.035",`updated_at`="2026-06-19 22:18:35.038",`deleted_at`=NULL WHERE `id` = 8
			assert.NoError(t, r8.Save())
			// Reload just in case
			r8 = Itemv3{ID: r8.ID}
			// SELECT * FROM `itemv3` WHERE `itemv3`.`id` = 8 ORDER BY `itemv3`.`id` LIMIT 1
			assert.NoError(t, dbV2.First(&r8).Error)
			assert.Equal(t, "", r8.NameNoDefault)
			assert.Equal(t, "", *r8.NameBlankDefault)
			assert.Equal(t, "", *r8.NameTextDefault)
			assert.Equal(t, "", r8.NameDBDefault)
			assert.Equal(t, 0, r8.InStockNoDefault)
			assert.Equal(t, 0, *r8.InStockNumberDefault)
			assert.Equal(t, 0, *r8.InStockZeroDefault)
			assert.Equal(t, 0, r8.InStockDBDefault)
		}
	})

	t.Run("CreateAllWithZero", func(t *testing.T) {
		//*******************************************************************************************************
		// Create a new record with all columns set to empty or 0.
		// This shows that Create applies empty string and 0 to Pointer fields.
		//*******************************************************************************************************

		r9 := Itemv3{
			NameNoDefault:        "",
			NameBlankDefault:     new(""),
			NameTextDefault:      new(""),
			NameDBDefault:        "",
			InStockNoDefault:     0,
			InStockNumberDefault: new(0),
			InStockZeroDefault:   new(0),
			InStockDBDefault:     0,
		}

		// INSERT INTO `itemv3` (`name_no_default`,`name_blank_default`,`name_text_default`,`name_db_default`,`in_stock_no_default`,`in_stock_zero_default`,`in_stock_number_default`,`in_stock_db_default`,`created_at`,`updated_at`,`deleted_at`) VALUES ("","","","",0,0,0,0,"2026-06-19 23:00:17.059","2026-06-19 23:00:17.059",NULL) RETURNING `id`
		if !assert.NoError(t, dbV2.Create(&r9).Error) {
			return
		} else if assert.GreaterOrEqual(t, r9.ID, uint(1)) {
			assert.Equal(t, "", r9.NameNoDefault)
			assert.Equal(t, "", *r9.NameBlankDefault)
			assert.Equal(t, "", *r9.NameTextDefault)
			assert.Equal(t, "", r9.NameDBDefault)
			assert.Equal(t, 0, r9.InStockNoDefault)
			assert.Equal(t, 0, *r9.InStockNumberDefault)
			assert.Equal(t, 0, *r9.InStockZeroDefault)
			assert.Equal(t, 0, r9.InStockDBDefault)
		}
	})
}
