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

func TestMahasiswaHandler(t *testing.T) {
	db, err := initDatabase()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	var data = []model.Mahasiswa{
		model.Mahasiswa{NPM: "1234567890", Nama: "AIRTAS", Kelas: "5KA34"},
		model.Mahasiswa{NPM: "1233255422", Nama: "SARITAS", Kelas: "5KA14"},
		model.Mahasiswa{NPM: "1232341233", Nama: "UNIRAS", Kelas: "5KA24"},
		model.Mahasiswa{NPM: "1232455644", Nama: "DIA", Kelas: "5KA12"},
		model.Mahasiswa{NPM: "1223131454", Nama: "NOVI", Kelas: "5KA12"},
	}
	webHandler := http.HandlerFunc(handler.SS)
	handler.RegisDB(db)

	t.Run("Tes Post", func(t *testing.T) {
		for _, item := range data {
			res := httptest.NewRecorder()
			jsonMarshal, err := json.MarshalIndent(item, "", "")
			if err != nil {
				t.Fatal(err)
			}
			req, err := http.NewRequest(http.MethodPost, "/api/ss/mahasiswa", bytes.NewBuffer(jsonMarshal))
			if err != nil {
				t.Fatal(err)
			}
			webHandler.ServeHTTP(res, req)
			buff, err := ioutil.ReadAll(res.Body)
			if err != nil {
				t.Fatal(err)
			}
			got := model.Mahasiswa{}
			if err := json.Unmarshal(buff, &got); err != nil {
				t.Fatal(err)
			}
			// compareMahasiswa(t, got, item)
		}

	})

	t.Run("Tes Gets", func(t *testing.T) {
		res := httptest.NewRecorder()
		req, err := http.NewRequest(http.MethodGet, "/api/ss/mahasiswa", nil)
		if err != nil {
			t.Fatal(err)
		}
		webHandler.ServeHTTP(res, req)
		buff, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Fatal(err)
		}
		got := []model.Mahasiswa{}
		if err := json.Unmarshal(buff, &got); err != nil {
			t.Fatal(err)
		}
		for index, item := range got {
			compareMahasiswa(t, item, data[index])
		}

	})
	t.Run("Tes Gets With Params", func(t *testing.T) {
		res := httptest.NewRecorder()
		params := fmt.Sprintf("npm,=,%s;kelas,=,%s", data[0].NPM, data[0].Kelas)
		req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/api/ss/mahasiswa?params=%s", url.QueryEscape(params)), nil)
		if err != nil {
			t.Fatal(err)
		}
		webHandler.ServeHTTP(res, req)
		buff, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Fatal(err)
		}
		got := []model.Mahasiswa{}
		if err := json.Unmarshal(buff, &got); err != nil {
			t.Fatal(err)
		}
		// compareNilai(t, got[0], dataInsertNilai[0])
	})
	t.Run("Tes Get 1 dataInsertNilai", func(t *testing.T) {
		res := httptest.NewRecorder()
		// params := fmt.Sprintf("npm,=%s; kelas,=,%s", dataInsertNilai[0].NPM, dataInsertNilai[0].Kelas)
		req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/api/ss/mahasiswa/%v", data[0].NPM), nil)
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
		dataUpdate := map[string]interface{}{
			"npm": data[0].NPM, "kelas": "1KB20",
		}
		jsonUpdate, err := json.MarshalIndent(dataUpdate, "", "")
		if err != nil {
			t.Fatal(err)
		}
		req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("/api/ss/mahasiswa/%s", data[0].NPM), bytes.NewBuffer(jsonUpdate))
		if err != nil {
			t.Fatal(err)
		}
		webHandler.ServeHTTP(res, req)
		buff, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Fatal(err)
		}
		got := model.Mahasiswa{}
		if err := json.Unmarshal(buff, &got); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("Tes Delete", func(t *testing.T) {
		res := httptest.NewRecorder()
		req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("/api/ss/mahasiswa/%s", data[0].NPM), nil)
		if err != nil {
			t.Fatal(err)
		}
		webHandler.ServeHTTP(res, req)
		if fmt.Sprintf("%v", res.Body) != "true" {
			t.Fatal("npm tidak terhapus")
		}
	})

}
