package test

import (
	"CRUD/handler"
	"CRUD/model"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	// "net/url"
	"testing"
)

func TestMatakuliahHandler(t *testing.T) {
	db, err := initDatabase()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	var dataInsertMatkul = []model.Matakuliah{
		model.Matakuliah{Kd_mk: "1si2", Mata_kuliah: "Sistem Informasi"},
		model.Matakuliah{Kd_mk: "1ea2", Mata_kuliah: "Sistem dataInsertMatkul"},
		model.Matakuliah{Kd_mk: "1pa2", Mata_kuliah: "Sistem Terdistribusi"},
		model.Matakuliah{Kd_mk: "1st2", Mata_kuliah: "Algortima Pemrograman"},
	}
	webHandler := http.HandlerFunc(handler.SS)
	handler.RegisDB(db)

	t.Run("Tes Post", func(t *testing.T) {
		for _, item := range dataInsertMatkul {
			res := httptest.NewRecorder()
			jsonMarshal, err := json.MarshalIndent(item, "", "")
			if err != nil {
				t.Fatal(err)
			}
			req, err := http.NewRequest(http.MethodPost, "/api/ss/matakuliah", bytes.NewBuffer(jsonMarshal))
			if err != nil {
				t.Fatal(err)
			}
			webHandler.ServeHTTP(res, req)
			buff, err := ioutil.ReadAll(res.Body)
			if err != nil {
				t.Fatal(err)
			}
			got := model.Matakuliah{}
			if err := json.Unmarshal(buff, &got); err != nil {
				// fmt.Println(got)
				t.Fatal(err)
			}
			// compareMatakuliah(t, got, item)
		}

	})

	t.Run("Tes Gets", func(t *testing.T) {
		res := httptest.NewRecorder()
		req, err := http.NewRequest(http.MethodGet, "/api/ss/matakuliah", nil)
		if err != nil {
			t.Fatal(err)
		}
		webHandler.ServeHTTP(res, req)
		buff, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Fatal(err)
		}
		got := []model.Matakuliah{}
		if err := json.Unmarshal(buff, &got); err != nil {
			t.Fatal(err)
		}
		// for index, item := range got {
		// 	compareMatakuliah(t, item, dataInsertMatkul[index])
		// }

	})
	t.Run("Tes Gets With Params", func(t *testing.T) {
		res := httptest.NewRecorder()
		params := fmt.Sprintf("kd_mk,=,%s;mata_kuliah,=,%s", dataInsertMatkul[2].Kd_mk, dataInsertMatkul[2].Mata_kuliah)
		req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/api/ss/matakuliah?params=%s", url.QueryEscape(params)), nil)
		if err != nil {
			t.Fatal(err)
		}
		webHandler.ServeHTTP(res, req)
		buff, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Fatal(err)
		}
		got := []model.Matakuliah{}
		if err := json.Unmarshal(buff, &got); err != nil {
			t.Fatal(err)
		}
		// compareNilai(t, got[0], dataInsertMatkulInsertNilai[0])
	})

	t.Run("Tes Get 1 dataInsertNilai", func(t *testing.T) {
		res := httptest.NewRecorder()
		// params := fmt.Sprintf("npm,=%s; kelas,=,%s", dataInsertNilai[0].NPM, dataInsertNilai[0].Kelas)
		req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/api/ss/matakuliah/%v", dataInsertMatkul[0].Kd_mk), nil)
		if err != nil {
			t.Fatal(err)
		}
		webHandler.ServeHTTP(res, req)

		buff, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Fatal(err)
		}
		got := model.Nilai{}
		if err := json.Unmarshal(buff, &got); err != nil {
			fmt.Println(got)
			t.Fatal(err)
		}
		// compareNilai(t, got, data[1])
	})

	t.Run("Tes Put", func(t *testing.T) {
		res := httptest.NewRecorder()
		dataInsertMatkulUpdate := map[string]interface{}{
			"mata_kuliah": "sistem",
		}
		jsonUpdate, err := json.MarshalIndent(dataInsertMatkulUpdate, "", "")
		if err != nil {
			t.Fatal(err)
		}
		req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("/api/ss/matakuliah/%s", dataInsertMatkul[1].Kd_mk), bytes.NewBuffer(jsonUpdate))
		if err != nil {
			t.Fatal(err)
		}
		webHandler.ServeHTTP(res, req)
		buff, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Fatal(err)
		}
		got := model.Matakuliah{}
		if err := json.Unmarshal(buff, &got); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Tes Delete", func(t *testing.T) {
		res := httptest.NewRecorder()
		req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("/api/ss/matakuliah/%s", dataInsertMatkul[0].Kd_mk), nil)
		// fmt.Println(req)
		if err != nil {
			t.Fatal(err)
		}
		webHandler.ServeHTTP(res, req)
		if fmt.Sprintf("%v", res.Body) != "true" {
			t.Fatal("npm tidak terhapus")
		}
	})

}
