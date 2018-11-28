package test

import (
	"database/sql"
	"nyoba/Latihan/CRUD/lib"
	"nyoba/Latihan/CRUD/model"
	"testing"
)

var username, pasword, namadatabase, databaseDefault string

func init() {
	username = "postgres"
	pasword = "postgres"
	namadatabase = "latihancrud"
	databaseDefault = "postgres"
}

func TestDatabase(t *testing.T) {
	t.Run("Testing koneksi postgres", func(t *testing.T) {
		db, err := lib.Connect(username, pasword, databaseDefault)

		defer db.Close()

		if err != nil {
			t.Fatal(err)
		}

	})
	t.Run("Testing Drop postgres", func(t *testing.T) {
		db, err := lib.Connect(username, pasword, databaseDefault)

		defer db.Close()

		if err != err {
			t.Fatal(err)
		}
		err = lib.DropDB(db, namadatabase)
		if err != nil {
			t.Fatal(err)
		}
	})
	t.Run("Testing Create postgres", func(t *testing.T) {
		db, err := lib.Connect(username, pasword, databaseDefault)

		defer db.Close()

		if err != err {
			t.Fatal(err)
		}
		err = lib.CreateDB(db, namadatabase)
		if err != nil {
			t.Fatal(err)
		}
	})

	// 	t.Run("Testing Create table Mahasiswa", func(t *testing.T) {
	// 		db, err := lib.Connect(username, pasword, namadatabase)
	// 		defer db.Close()
	//
	// 		if err != err {
	// 			t.Fatal(err)
	// 		}
	// 		if err = lib.CreateTable(db, model.TBMahasiswa); err != nil {
	// 			t.Fatal(err)
	// 		}
	// 	})
	//
}

func initDatabase() (*sql.DB, error) {
	db, err := lib.Connect(username, pasword, databaseDefault)
	if err != err {
		return nil, err
	}
	if err = lib.DropDB(db, namadatabase); err != nil {
		return nil, err
	}
	if err = lib.CreateDB(db, namadatabase); err != nil {
		return nil, err
	}
	db.Close()

	db, err = lib.Connect(username, pasword, namadatabase)
	if err != err {
		return nil, err
	}
	if err = lib.CreateTable(db, model.TBMahasiswa); err != nil {
		return nil, err
	}
	db.Close()

	db, err = lib.Connect(username, pasword, namadatabase)
	if err != err {
		return nil, err
	}
	if err = lib.CreateTable(db, model.TBMatkul); err != nil {
		return nil, err
	}
	return db, nil
}
