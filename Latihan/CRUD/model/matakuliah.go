package model

import (
	"database/sql"
	"nyoba/Latihan/CRUD/lib"
)

type Matkul struct {
	Kd_mk       string
	Mata_Kuliah string
}

var TBMatkul string = `
    CREATE TABLE Matkul(
   	Kd_mk VARCHAR(10)PRIMARY KEY NOT NULL,
		mata_kuliah VARCHAR (20) NOT NULL
  );
`

func (k *Matkul) Name() string {
	return "Matkul"

}
func (m *Matkul) Field() ([]string, []interface{}) {
	fields := []string{"kd_mk", "Mata_Kuliah"}
	dst := []interface{}{&m.Kd_mk, &m.Mata_Kuliah}
	return fields, dst
}
func (m *Matkul) PrimaryKey() ([]string, []interface{}) {
	fields := []string{"kd_mk"}
	dst := []interface{}{&m.Kd_mk}
	return fields, dst
}
func (m *Matkul) Structure() lib.Table {
	return &Matkul{}
}
func (m *Matkul) Insert(db *sql.DB) error {
	return lib.Insert(db, m)
}
func (m *Matkul) Delete(db *sql.DB) error {
	return lib.Delete(db, m)
}
func (m *Matkul) Get(db *sql.DB) error {
	return lib.Get(db, m)
}
func (m *Matkul) Update(db *sql.DB, data map[string]interface{}) error {
	return lib.Update(db, m, data)
}

func GetAllMatkul(db *sql.DB, params ...string) ([]*Matkul, error) {
	m := &Matkul{}
	data, err := lib.Gets(db, m, params...)
	if err != nil {
		return nil, err
	}
	matkul := make([]*Matkul, len(data))
	for index, item := range data {
		matkul[index] = item.(*Matkul)
	}
	return matkul, nil
}
