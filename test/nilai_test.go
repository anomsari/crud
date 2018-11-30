package test

import (
	"CRUD/model"
	"fmt"
	"testing"
)

func TestNilai(t *testing.T) {
	var dataInsertNilai = []model.Nilai{
		model.Nilai{Kd_mk: "1si2", NPM: "1234567890", UTS: 90, UAS: 60},
		model.Nilai{Kd_mk: "1ea2", NPM: "1233255422", UTS: 70, UAS: 40},
		model.Nilai{Kd_mk: "1pa2", NPM: "1232341233", UTS: 50, UAS: 80},
		model.Nilai{Kd_mk: "1st2", NPM: "1232455644", UTS: 80, UAS: 90},
	}

	db, err := initDatabase()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	t.Run("testing Insert", func(t *testing.T) {
		for index, dataInsert := range dataInsertNilai {
			err := dataInsert.Insert(db)
			if err != nil {
				t.Fatal(err)

			}
			got := model.Nilai{IdNilai: index + 1}
			want := model.Nilai{
				IdNilai: index + 1,
				Kd_mk:   dataInsert.Kd_mk,
				NPM:     dataInsert.NPM,
				UTS:     dataInsert.UTS,
				UAS:     dataInsert.UAS,
			}
			if err := got.Get(db); err != nil {
				t.Fatal(err)
			}
			// fmt.Println(t, got, want)
			compareNilai(t, got, want)

		}
	})

	t.Run("testing Update Get", func(t *testing.T) {
		update := map[string]interface{}{
			"uts": 20,
			"uas": 3,
		}

		dataUpdate := model.Nilai{IdNilai: 1}
		if err := dataUpdate.Update(db, update); err != nil {
			t.Fatal(err)
		}
		got := model.Nilai{IdNilai: dataUpdate.IdNilai}
		// want := model.Nilai{Kd_mk: "1si2", NPM: "1234567890", UTS: 20, UAS: 3}
		if err := got.Get(db); err != nil {
			t.Fatal(err)

		}
		// compareNilai(t, got, want)
	})

	t.Run("Testing Delete", func(t *testing.T) {
		m := model.Nilai{IdNilai: 3}
		if err := m.Delete(db); err != nil {
			t.Fatal(err)
		}

	})

	t.Run("testing Gets", func(t *testing.T) {
		result, err := model.GetAllNilai(db)
		if err != nil {
			t.Fatal(err)
		}
		for _, item := range result {
			fmt.Println(*item)

			got := model.Nilai{IdNilai: item.IdNilai}
			if err := got.Get(db); err != nil {
				t.Fatal(err)
			}
			compareNilai(t, got, *item)
		}
	})

	t.Run("testing gets with Paramaters", func(t *testing.T) {
		params := "uts,=,90"
		result, err := model.GetAllNilai(db, params)
		if err != nil {
			t.Fatal(err)
		}
		for _, item := range result {
			fmt.Println(*item)

			// 			if err := got.Get(db); err != nil {
			// 				t.Fatal(err)
			// 			}
			// 			compareNilai(t, got, *item)
			//
		}
	})

}

func compareNilai(t *testing.T, got, want model.Nilai) {
	if got.Kd_mk != want.Kd_mk {
		t.Fatalf("got : %s want :%s KdMK tidak sama", got.Kd_mk, want.Kd_mk)
	}
	if got.NPM != want.NPM {
		t.Fatalf("got :%s want :%s MataKuliah tidak Sama", got.NPM, want.NPM)
	}
}
