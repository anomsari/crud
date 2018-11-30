package handler

import (
	"CRUD/model"
	"encoding/json"
	"strconv"
	// "fmt"
	"net/http"
	// "strings"
	"io/ioutil"
)

//input
func HandlerNilaiPost(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var data model.Nilai
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
func HandlerNilaiDelete(w http.ResponseWriter, r *http.Request) {
	lastIndex := LastIndex(r)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// }
	data := model.Nilai{NPM: lastIndex}
	if err := data.Delete(db); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Write([]byte("true"))
}

//update
func HandlerNilaiPut(w http.ResponseWriter, r *http.Request) {
	lastIndex := LastIndex(r)
	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusBadRequest)
	// }
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
	i, _ := strconv.Atoi(lastIndex)
	data := model.Nilai{IdNilai: i}
	err = data.Update(db, jsonMap)
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

//get
func HandlerNilaiGet(w http.ResponseWriter, r *http.Request) {
	lastIndex := LastIndex(r)
	if lastIndex == "nilai" {
		// 	http.Error(w, err.Error(), http.StatusInternalServerError)
		// params, _ := r.URL.Query()["params"]
		data, err := model.GetAllNilai(db)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		jsonData, _ := json.Marshal(data)
		w.Write(jsonData)
		return
	} else {
		i, _ := strconv.Atoi(lastIndex)
		data := model.Nilai{IdNilai: i}
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
