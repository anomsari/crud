package test

import (
	"nyoba/Latihan/CRUD/model"
	"testing"
)

func TestMahasiswa(t *testing.T) {
	var data = []model.Mahasiswa{
		model.Mahasiswa{NPM: "12345678", Nama: "Moon", Kelas: "1KB04"},
		model.Mahasiswa{NPM: "12222111", Nama: "Xen", Kelas: "1KB04"},
		model.Mahasiswa{NPM: "12345645", Nama: "Lo", Kelas: "1KB04"},
		model.Mahasiswa{NPM: "12345679", Nama: "jo", Kelas: "1KB04"},
	}
	db, err := initDatabase()
	if err != err {
		t.Fatal(err)
	}
	defer db.Close()

	t.Run("Test Insert", func(t *testing.T) {
		for _, dataInsert := range data {
			err := dataInsert.Insert(db)
			if err != nil {
				t.Fatal(err)
			}
			got := model.Mahasiswa{NPM: dataInsert.NPM}
			if err := got.Get(db); err != nil {
				t.Fatal(err)
			}
			CompareMahasiswa(t, got, dataInsert)
		}
	})

	t.Run("Test Update", func(t *testing.T) {
		update := map[string]interface{}{
			"nama": "Moon",
		}
		dataupdate := data[0]
		if err := dataupdate.Update(db, update); err != nil {
			t.Fatal(err)
		}
		got := model.Mahasiswa{NPM: dataupdate.NPM}
		if err := got.Get(db); err != nil {
			t.Fatal(err)
		}
		CompareMahasiswa(t, got, dataupdate)
	})

	t.Run("Test Gets", func(t *testing.T) {
		result, err := model.GetAllMahasiswa(db)
		if err != nil {
			t.Fatal()
		}
		for _, item := range result {
			got := model.Mahasiswa{NPM: item.NPM}
			if err := got.Get(db); err != nil {
				t.Fatal(err)
			}
			CompareMahasiswa(t, got, *item)
		}
	})

	t.Run("Test get with parameter", func(t *testing.T) {
		params := "nama, =, Moon"
		result, err := model.GetAllMahasiswa(db, params)
		if err != nil {
			t.Fatal(err)
		}
		for _, item := range result {
			got := model.Mahasiswa{NPM: item.NPM}
			if err := got.Get(db); err != nil {
				t.Fatal(err)
			}
			CompareMahasiswa(t, got, *item)
		}
	})

	t.Run("Test Delete", func(t *testing.T) {
		m := model.Mahasiswa{NPM: data[0].NPM}
		if err := m.Delete(db); err != nil {
			t.Fatal(err)
		}
	})
}

func CompareMahasiswa(t *testing.T, got, want model.Mahasiswa) {
	if got.NPM != want.NPM {
		t.Fatalf("got : %s want : %s npm tidak sama", got.NPM, want.NPM)
	}
	if got.Nama != want.Nama {
		t.Fatalf("got : %s want : %s npm tidak sama", got.Nama, want.Nama)
	}
	if got.Kelas != want.Kelas {
		t.Fatalf("got : %s want : %s npm tidak sama", got.Kelas, want.Kelas)
	}
}
