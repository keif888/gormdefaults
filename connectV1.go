package gormdefaultsv1

import (
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// Stores the database connection for all V1 tests
var dbV1 *gorm.DB

// Connects the Gorm V1 database to SQLite, creates the test table, and applies DDL modifications for database default columns.
func ConnectV1DB() (*gorm.DB, error) {
	const dbName = ".gormdefaultsv1.db"
	_ = os.Remove(dbName)
	db, err := gorm.Open("sqlite3", dbName)
	if enableDebug {
		dbV1 = db.Debug()
	} else {
		dbV1 = db
	}
	if err != nil {
		return nil, err
	} else if err = dbV1.AutoMigrate(&Itemv1{}).Error; err != nil {
		return dbV1, err
	} else if err = dbV1.Exec("ALTER TABLE itemv1 RENAME COLUMN name_db_default TO name_db_default_old;").Error; err != nil {
		return dbV1, err
	} else if err = dbV1.Exec("ALTER TABLE itemv1 ADD COLUMN name_db_default bytes DEFAULT 'NaS';").Error; err != nil {
		return dbV1, err
	} else if err = dbV1.Exec("ALTER TABLE itemv1 DROP COLUMN name_db_default_old;").Error; err != nil {
		return dbV1, err
	} else if err = dbV1.Exec("ALTER TABLE itemv1 RENAME COLUMN in_stock_db_default TO in_stock_db_default_old;").Error; err != nil {
		return dbV1, err
	} else if err = dbV1.Exec("ALTER TABLE itemv1 ADD COLUMN in_stock_db_default integer DEFAULT -1;").Error; err != nil {
		return dbV1, err
	} else if err = dbV1.Exec("ALTER TABLE itemv1 DROP COLUMN in_stock_db_default_old;").Error; err != nil {
		return dbV1, err
	}
	return dbV1, err
}

// Closes the database for V2
func CloseV1DB() error {
	if dbV1 != nil {
		return dbV1.Close()
	}
	return nil
}
