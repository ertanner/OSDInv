package main

import (
	"database/sql"
	_ "github.com/alexbrainman/odbc"
	"log"
	"os"
	"encoding/json"
	"fmt"
	"github.com/rs/cors"
	"net/http"
	"text/template"
	"github.com/gorilla/mux"
)

var db *sql.DB
var r = mux.NewRouter().StrictSlash(true)
var err error
var tpl *template.Template

func init(){
	tpl = template.Must(template.ParseGlob("html/*.html"))
}

type Configuration struct {
	HttpPort string
	ConnectionString string
	Appname string
	Runmode string
	sqluser string
	sqlpass string
	sqldb string
	SessionName string
}

func main() {
	configuration := Configuration{}
	filename := "app.json"
	log.Println(filename)
	pwd, _ := os.Getwd()
	fto := pwd+"\\"+filename
	log.Println(fto)
	file, err := os.Open(fto)
	if err != nil {
		log.Println("File Open error")
		os.Exit(500) //return err
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&configuration)
	if err != nil {
		log.Println("json erorr")
		os.Exit(500)//return err
	}

	// ********************************************************
	// Create the database handle, confirm driver is present
	// *********************************************************
	connectString := configuration.sqluser + ":" + configuration.sqlpass + configuration.sqldb
	log.Println(connectString)
	db, err = sql.Open("odbc", "DSN=DYLT_IMP"  )
	if err != nil {
		log.Fatalf("Error on initializing database connection: %s", err.Error())
	}
	fmt.Println("db opened at root:****@/test")
	db.SetMaxIdleConns(100)
	defer db.Close()
	// make sure connection is available
	err = db.Ping()
	if err != nil {
		log.Fatalf("Error on opening database connection: %s", err.Error())
	}else {fmt.Println("verified db is open")}


	r.HandleFunc("/", homePage)
	r.HandleFunc("addItem", addItem)
	r.HandleFunc("removeItems", removeItems)
	r.HandleFunc("foundItem", foundItem)
	r.HandleFunc("/osdInv/{id}", osdInv)

	// open for business
	fmt.Println("Router is open for business on port " + configuration.HttpPort)
	port := ":"+ configuration.HttpPort
	log.Println(port)
	handler := cors.Default().Handler(r)
	http.ListenAndServe(port, handler)
}

func homePage(w http.ResponseWriter, r *http.Request)  {
	vars := mux.Vars(r)
	key := vars["id"]
	log.Println(key)
	tpl.Execute(w)
}