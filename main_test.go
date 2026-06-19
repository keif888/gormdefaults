package gormdefaultsv1

import (
	"fmt"
	"os"
	"testing"
)

// Initiate the tests after setting the database connections up
func TestMain(m *testing.M) {
	db, err := ConnectV1DB()
	if err != nil {
		fmt.Printf("connectv1db: failed with (%s)\n", err.Error())
		if db != nil {
			if err = CloseV1DB(); err != nil {
				fmt.Printf("closev1db: failed with (%s)\n", err.Error())
			}
		}
		os.Exit(1)
	}

	dbv2, errv2 := ConnectV2DB()
	if errv2 != nil {
		fmt.Printf("connectv2db: failed with (%s)\n", errv2.Error())
		if dbv2 != nil {
			if errv2 = CloseV2DB(); errv2 != nil {
				fmt.Printf("closev2db: failed with (%s)\n", errv2.Error())
			}
		}
		if err = CloseV1DB(); err != nil {
			fmt.Printf("closev1db: failed with (%s)\n", err.Error())
		}
		os.Exit(1)
	}

	code := m.Run()

	if err = CloseV1DB(); err != nil {
		fmt.Printf("closev1db: failed with (%s)\n", err.Error())
		os.Exit(1)
	}
	if err = CloseV2DB(); err != nil {
		fmt.Printf("closev2db: failed with (%s)\n", err.Error())
		os.Exit(1)
	}
	os.Exit(code)
}
