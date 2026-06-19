package gormdefaultsv1

import (
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Stores the database connection for all V2 tests
var dbV2 *gorm.DB

// Connects the Gorm V2 database to SQLite, creates the test table, and applies DDL modifications for database default columns.
func ConnectV2DB() (*gorm.DB, error) {
	const dbName = ".gormdefaultsv2.db"
	_ = os.Remove(dbName)
	db, err := gorm.Open(sqlite.Open(dbName), &gorm.Config{})
	if enableDebug {
		dbV2 = db.Debug()
	} else {
		dbV2 = db
	}
	if err != nil {
		return nil, err
	} else if err = dbV2.AutoMigrate(&Itemv2{}); err != nil {
		return dbV2, err
	} else if err = dbV2.Exec("ALTER TABLE itemv2 RENAME COLUMN name_db_default TO name_db_default_old;").Error; err != nil {
		return dbV2, err
	} else if err = dbV2.Exec(`ALTER TABLE itemv2 ADD COLUMN name_db_default BLOB DEFAULT "NaS";`).Error; err != nil {
		return dbV2, err
	} else if err = dbV2.Exec("ALTER TABLE itemv2 DROP COLUMN name_db_default_old;").Error; err != nil {
		return dbV2, err
	} else if err = dbV2.Exec("ALTER TABLE itemv2 RENAME COLUMN in_stock_db_default TO in_stock_db_default_old;").Error; err != nil {
		return dbV2, err
	} else if err = dbV2.Exec("ALTER TABLE itemv2 ADD COLUMN in_stock_db_default INTEGER DEFAULT -1;").Error; err != nil {
		return dbV2, err
	} else if err = dbV2.Exec("ALTER TABLE itemv2 DROP COLUMN in_stock_db_default_old;").Error; err != nil {
		return dbV2, err
	} else if err = dbV2.AutoMigrate(&Itemv3{}); err != nil {
		return dbV2, err
	} else if err = dbV2.Exec("ALTER TABLE itemv3 RENAME COLUMN name_db_default TO name_db_default_old;").Error; err != nil {
		return dbV2, err
	} else if err = dbV2.Exec(`ALTER TABLE itemv3 ADD COLUMN name_db_default BLOB DEFAULT "NaS";`).Error; err != nil {
		return dbV2, err
	} else if err = dbV2.Exec("ALTER TABLE itemv3 DROP COLUMN name_db_default_old;").Error; err != nil {
		return dbV2, err
	} else if err = dbV2.Exec("ALTER TABLE itemv3 RENAME COLUMN in_stock_db_default TO in_stock_db_default_old;").Error; err != nil {
		return dbV2, err
	} else if err = dbV2.Exec("ALTER TABLE itemv3 ADD COLUMN in_stock_db_default INTEGER DEFAULT -1;").Error; err != nil {
		return dbV2, err
	} else if err = dbV2.Exec("ALTER TABLE itemv3 DROP COLUMN in_stock_db_default_old;").Error; err != nil {
		return dbV2, err
	}
	return dbV2, err
}

// Closes the database for V2
func CloseV2DB() error {
	if dbV2 != nil {
		sqlDB, _ := dbV2.DB()
		return sqlDB.Close()
	}
	return nil
}
