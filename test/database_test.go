package test

import (
	"CRUD/lib"
	"CRUD/model"
	"database/sql"
	"testing"
)

var username, password, namaDatabase, databaseDefault string

func init() {
	username = "postgres"
	password = "P@ssw0rd"
	namaDatabase = "latcrud"
	databaseDefault = "postgres"
}
func TestDatabase(t *testing.T) {
	t.Run("Testing untuk koneksi postgres Database ", func(t *testing.T) {
		db, err := lib.Connect(username, password, databaseDefault)

		defer db.Close()

		if err != nil {
			t.Fatal(err)
		}

	})
	t.Run("Testing Untuk Drop Databse ", func(t *testing.T) {
		db, err := lib.Connect(username, password, databaseDefault)

		defer db.Close()

		if err != nil {
			t.Fatal(err)
		}

		err = lib.DropDB(db, namaDatabase)

		if err != nil {
			t.Fatal(err)
		}
	})
	t.Run("Testing Untuk Create Databse ", func(t *testing.T) {
		db, err := lib.Connect(username, password, databaseDefault)

		defer db.Close()

		if err != nil {
			t.Fatal(err)
		}

		err = lib.CreateDB(db, namaDatabase)

		if err != nil {
			t.Fatal(err)
		}
	})

}

func TestTableMahasiswa(t *testing.T) {
	// t.Run("Testing Untuk Drop Databse ", func(t *testing.T) {
	// 	db, err := lib.Connect(username, password, namaDatabase)

	// 	defer db.Close()

	// 	if err != nil {
	// 		t.Fatal(err)
	// 	}

	// 	if err = lib.DropTable(db, model.TbMahasiswa); err != nil {
	// 		t.Fatal(err)
	// 	}
	// })
	t.Run("Testing Untuk Create Databse ", func(t *testing.T) {
		db, err := lib.Connect(username, password, namaDatabase)
		defer db.Close()

		if err != nil {
			t.Fatal(err)
		}
		err = lib.CreateTable(db, model.TbMahasiswa)
		if err != nil {
			t.Fatal(err)
		}
	})

}

func initDatabase() (*sql.DB, error) {
	db, err := lib.Connect(username, password, databaseDefault)
	if err != nil {
		return nil, err
	}
	if err = lib.DropDB(db, namaDatabase); err != nil {
		return nil, err
	}
	if err = lib.CreateDB(db, namaDatabase); err != nil {
		return nil, err
	}
	db.Close()

	db, err = lib.Connect(username, password, namaDatabase)
	if err != nil {
		return nil, err
	}
	if err = lib.CreateTable(db, model.TbMahasiswa); err != nil {
		return nil, err
	}

	if err = lib.CreateTable(db, model.TbMatkul); err != nil {
		return nil, err
	}
	if err = lib.CreateTable(db, model.TbNilai); err != nil {
		return nil, err
	}
	return db, nil
}
