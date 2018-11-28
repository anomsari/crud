package test

import (
        "testing"
        "nyoba/Latihan/CRUD/model"
      )

func TestMatkul(t *testing.T) {
  var data = []model.Matkul{
    model.Matkul{Kd_mk: "12a1", Mata_Kuliah: "Algoritma"},
    model.Matkul{Kd_mk: "12a2", Mata_Kuliah: "Matematika"},
    model.Matkul{Kd_mk: "12a3", Mata_Kuliah: "Kimia"},
    model.Matkul{Kd_mk: "12a4", Mata_Kuliah: "Jaringan"},
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
      got := model.Matkul{Kd_mk: dataInsert.Kd_mk}
      if err := got.Get(db); err != nil {
        t.Fatal(err)
      }
      CompareMatkul(t, got, dataInsert)
    }


  })
  t.Run("Test Update", func(t *testing.T) {
    update := map[string]interface{}{
      "Mata_Kuliah": "Algoritma",
    }
    dataupdate:= data[0]
    if err := dataupdate.Update(db, update); err != nil {
      t.Fatal(err)
    }
    got := model.Matkul{Kd_mk: dataupdate.Kd_mk}
    if err := got.Get(db); err != nil {
      t.Fatal(err)
    }
    CompareMatkul(t, got, dataupdate)
  })

  t.Run("Test Gets", func(t *testing.T) {
    result, err := model.GetAllMatkul(db)
    if err != nil {
      t.Fatal(err)
    }
    for _, item := range result {
      got := model.Matkul{Kd_mk: item.Kd_mk}
      if err := got.Get(db); err != nil {
        t.Fatal(err)
      }
      CompareMatkul(t, got, *item)
    }
  })
  t.Run("Test get with parameter", func(t *testing.T) {
    params := "Mata_kuliah, =, Algoritma"
    result, err := model.GetAllMatkul(db, params)
    if err != nil {
      t.Fatal(err)
    }
    for _, item := range result {
      got := model.Matkul{Kd_mk: item.Kd_mk}
      if err := got.Get(db); err != nil {
        t.Fatal(err)
      }
      CompareMatkul(t, got, *item)
    }
  })

  t.Run("Test Delete", func(t *testing.T) {
    m := model.Matkul{Kd_mk: data[0].Kd_mk}
    if err := m.Delete(db); err != nil {
      t.Fatal(err)
    }
  })
}
func CompareMatkul(t *testing.T, got, want model.Matkul) {
	if got.Kd_mk!= want.Kd_mk{
		t.Fatalf("got : %s want : %s kode matkul tidak sama", got.Kd_mk, want.Kd_mk)
	}
	if got.Mata_Kuliah != want.Mata_Kuliah {
		t.Fatalf("got : %s want : %s nama tidak sama", got.Mata_Kuliah, want.Mata_Kuliah)
	}
}
