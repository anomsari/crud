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

func TestNilaiHandler(t *testing.T) {
	db, err := initDatabase()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()
	var dataInsertNilai = []model.Nilai{
		model.Nilai{Kd_mk: "1si2", NPM: "1234567890", UTS: 90, UAS: 60},
		model.Nilai{Kd_mk: "1ea2", NPM: "1233255422", UTS: 70, UAS: 40},
		model.Nilai{Kd_mk: "1pa2", NPM: "1232341233", UTS: 50, UAS: 80},
		model.Nilai{Kd_mk: "1st2", NPM: "1232455644", UTS: 80, UAS: 90},
	}
	webHandler := http.HandlerFunc(handler.SS)
	handler.RegisDB(db)

	t.Run("Tes Post", func(t *testing.T) {
		for index, item := range dataInsertNilai {
			res := httptest.NewRecorder()
			jsonMarshal, err := json.MarshalIndent(item, "", "")
			if err != nil {
				t.Fatal(err)
			}
			req, err := http.NewRequest(http.MethodPost, "/api/ss/nilai", bytes.NewBuffer(jsonMarshal))
			if err != nil {
				t.Fatal(err)
			}
			webHandler.ServeHTTP(res, req)
			buff, err := ioutil.ReadAll(res.Body)
			if err != nil {
				t.Fatal(err)
			}
			got := model.Nilai{IdNilai: index + 1}
			want := model.Nilai{IdNilai: index + 1,
				NPM:   item.NPM,
				Kd_mk: item.Kd_mk,
				UAS:   item.UAS,
				UTS:   item.UTS,
			}
			if err := json.Unmarshal(buff, &got); err != nil {
				// fmt.Println(got)
				t.Fatal(err)
			}
			compareNilai(t, got, want)
		}

	})

	t.Run("Tes Gets", func(t *testing.T) {
		res := httptest.NewRecorder()
		req, err := http.NewRequest(http.MethodGet, "/api/ss/nilai", nil)
		if err != nil {
			t.Fatal(err)
		}
		webHandler.ServeHTTP(res, req)
		buff, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Fatal(err)
		}
		got := []model.Nilai{}
		if err := json.Unmarshal(buff, &got); err != nil {
			t.Fatal(err)
		}
		// for index, item := range got {
		// 	compareNilai(t, item, dataInsertNilai[index])
		// }

	})

	t.Run("Tes Gets With Params", func(t *testing.T) {
		res := httptest.NewRecorder()
		params := fmt.Sprintf("kd_mk,=%s; npm,=,%s", dataInsertNilai[1].Kd_mk, dataInsertNilai[1].NPM)
		req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/api/ss/nilai?params=%s", url.QueryEscape(params)), nil)
		if err != nil {
			t.Fatal(err)
		}
		webHandler.ServeHTTP(res, req)
		buff, err := ioutil.ReadAll(res.Body)
		if err != nil {
			t.Fatal(err)
		}
		got := []model.Nilai{}
		if err := json.Unmarshal(buff, &got); err != nil {
			t.Fatal(err)
		}
		// compareNilai(t, got[0], dataInsertNilai[0])
	})

	t.Run("Tes Get 1 dataInsertNilai", func(t *testing.T) {
		res := httptest.NewRecorder()
		// params := fmt.Sprintf("npm,=%s; kelas,=,%s", dataInsertNilai[0].NPM, dataInsertNilai[0].Kelas)
		req, err := http.NewRequest(http.MethodGet, fmt.Sprintf("/api/ss/nilai/%v", "2"), nil)
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
			// fmt.Println(got)
			t.Fatal(err)
		}
		compareNilai(t, got, dataInsertNilai[1])
	})

	t.Run("Tes Put", func(t *testing.T) {
		res := httptest.NewRecorder()
		dataUpdate := map[string]interface{}{
			"uts": 80, "uas": 30,
		}
		jsonUpdate, err := json.MarshalIndent(dataUpdate, "", "")
		if err != nil {
			t.Fatal(err)
		}
		req, err := http.NewRequest(http.MethodPut, fmt.Sprintf("/api/ss/nilai/%s", "3"), bytes.NewBuffer(jsonUpdate))
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
			t.Fatal(err)
		}
	})

	t.Run("Tes Delete", func(t *testing.T) {
		res := httptest.NewRecorder()
		req, err := http.NewRequest(http.MethodDelete, fmt.Sprintf("/api/ss/nilai/%s", dataInsertNilai[0].Kd_mk), nil)
		if err != nil {
			t.Fatal(err)
		}
		webHandler.ServeHTTP(res, req)
		if fmt.Sprintf("%v", res.Body) != "true" {
			t.Fatal("npm tidak terhapus")
		}
	})

}
