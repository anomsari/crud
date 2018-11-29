package main

import (
	"CRUD/handler"
	"CRUD/lib"
	"CRUD/model"
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os/exec"
)

var (
	dbUser    = "postgres"
	dbPass    = "P@ssw0rd"
	dbDefault = "postgres"
	dbUse     = "latcrud"
	port      = 8080
)

func init() {
	flag.IntVar(&port, "server", 8080, "Your port")
}

func main() {
	// flag.Parse()
	cmd := exec.Command("polymer", "build")
	cmd.Dir = fmt.Sprintf("FrontEnd")
	err := cmd.Start()
	if err != nil {
		log.Printf("Command Finished with error:%v", err.Error())
	}
	log.Printf("waiting for command to finish..")
	err = cmd.Wait()
	if err != nil {
		log.Printf("Command Finished with error: %v", err.Error())
	} else {
		log.Print("Command Finished without error")
	}
	//
	db, err := lib.Connect(dbUser, dbPass, dbUse)
	if err != nil {
		db, err = initDatabase()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
	}

	defer db.Close()
	handler.RegisDB(db)

	http.HandleFunc("/api/ss/", handler.SS)
	polymer := http.FileServer(http.Dir("FrontEnd/build/es6-bundled"))
	http.Handle("/", http.StripPrefix("/", polymer))

	// fmt.Printf("Port started at localhost:%v\n", fmt.Sprintf("%v", port))
	// log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))

	log.Println("localhost:8889")
	http.ListenAndServe(":8889", nil)

}
func initDatabase() (*sql.DB, error) {
	dbInit, err := lib.Connect(dbUser, dbUse, dbDefault)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	if err = lib.CreateDB(dbInit, dbUse); err != nil {
		return nil, err
	}
	fmt.Println("Re-creating Database ...")
	dbInit.Close()

	db, err := lib.Connect(dbUser, dbPass, dbDefault)
	if err != nil {
		return nil, err
	}
	if err = lib.CreateTable(db, model.TbMahasiswa); err != nil {
		fmt.Println("table mahasiswa")
		return nil, err
	}

	if err = lib.CreateTable(db, model.TbMatkul); err != nil {
		fmt.Println("table matkul")
		return nil, err
	}
	if err = lib.CreateTable(db, model.TbNilai); err != nil {
		fmt.Println("table nilai")
		return nil, err
	}
	return db, nil
}
