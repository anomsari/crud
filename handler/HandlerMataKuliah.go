package handler

import (
	"CRUD/model"
	"encoding/json"
	"fmt"
	// "fmt"
	"net/http"
	// "strings"
	"io/ioutil"
)

//input
func HandlerMataKuliahPost(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var data model.Matakuliah
	if err = json.Unmarshal(body, &data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err = data.Insert(db); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(jsonData)
}

//delete
func HandlerMataKuliahDelete(w http.ResponseWriter, r *http.Request) {
	lastIndex := LastIndex(r)
	data := model.Matakuliah{Kd_mk: lastIndex}
	if err := data.Delete(db); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Write([]byte("true"))
}

//update
func HandlerMataKuliahPut(w http.ResponseWriter, r *http.Request) {
	fmt.Println("terpanggil")
	lastIndex := LastIndex(r)
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	jsonMap := make(map[string]interface{})
	err = json.Unmarshal(body, &jsonMap)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	data := model.Matakuliah{Kd_mk: lastIndex}
	err = data.Update(db, jsonMap)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if err = data.Get(db); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	jsonData, err := json.Marshal(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Write(jsonData)
}

//get
func HandlerMataKuliahGet(w http.ResponseWriter, r *http.Request) {
	lastIndex := LastIndex(r)
	if lastIndex == "matakuliah" {
		params, _ := r.URL.Query()["params"]
		data, err := model.GetAllMatakuliah(db, params...)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		jsonData, _ := json.Marshal(data)
		w.Write(jsonData)
		return
	} else {
		data := model.Matakuliah{Kd_mk: lastIndex}
		err := data.Get(db)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		jsonData, err := json.Marshal(data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(jsonData)
	}
}
