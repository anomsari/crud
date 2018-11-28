package model

import (
	"database/sql"
	"nyoba/Latihan/CRUD/lib"
)

type Nilai struct {
	NPM      string
	Kd_mk    string
	UAS      int
	UTS      int
	Akhir    string
	ID_nilai string
	Index    string
}

var TBNilai string = `
    CREATE TABLE Nilai(
   	Kd_mk VARCHAR(10) NOT NULL,
		NPM VARCHAR (8) NOT NULL,
    UTS INT(3) NOT NULL,
    UAS INT (3) NOT NULL,
    ID_nilai SERIAL PRIMARY KEY NOT NULL
    Akhir VARCHAR (1) NOT NULL,
    Index VARCHAR (1) NOT NULL
  );
`

func (k *Nilai) Name() string {
	return "Nilai"
}

func (m *Nilai) Field() ([]string, []interface{}) {
	fields := []string{"kd_mk", "NPM", "UTS", "UAS", "Akhir"}
	dst := []interface{}{&m.Kd_mk, &m.NPM, &m.UAS, &m.UTS, &m.Akhir}
	return fields, dst
}

func (m *Nilai) PrimaryKey() ([]string, []interface{}) {
	fields := []string{"ID_nilai"}
	dst := []interface{}{&m.ID_nilai}
	return fields, dst
}
func (m *Nilai) Structure() lib.Table {
	return &Nilai{}
}
func (n *Nilai) Insert(db *sql.DB, params ...string) error {
	result := 0
	if n.UAS != 0 && n.UTS != 0 {
		n.Index = "E"
	}
	result = n.UAS + n.UTS
	var z uint = uint(result)
	if z > 85 {
		n.Index = "A"
	} else if z > 75 {
		n.Index = "B"
	} else {
		n.Index = "E"
	}
	return lib.Insert(db, n)
}

func (m *Nilai) Delete(db *sql.DB) error {
	return lib.Delete(db, m)
}
func (m *Nilai) Get(db *sql.DB) error {
	return lib.Get(db, m)
}
func (m *Nilai) Update(db *sql.DB, data map[string]interface{}) error {
	return lib.Update(db, m, data)
}
