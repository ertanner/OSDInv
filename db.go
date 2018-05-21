package main

import (
	"net/http"
	"log"
	"fmt"
	"github.com/gorilla/mux"
)

func osdInv(w http.ResponseWriter, r *http.Request) {
	j:= r.URL.Path
	params := mux.Vars(r)
	log.Println(params)
	//var msg string
	//var color string
	//var find string
	//var b = bpHit{}
	//rows, err := db.Query(`select * from DYLT_OSD_Inv_Mgt where found = 0 and description like '%?%'`, find)
	//if err != nil {
	//	log.Println(err)
	//}
	//i:= 0
	//for rows.Next() {
	//	err := rows.Scan(&statType, &statTime, &iref)
	//	if err != nil {
	//		log.Println(err)
	//	}
	//	d.Data[i].StatTime = statTime
	//	d.Data[i].Iref = iref
	//	i++
	//}
	//b.bp = msg
	//b.color = color
	//j, err := json.Marshal(b)
	//if err != nil {
	//	panic(err)
	//}
	log.Println(j)
	// w.Write(j)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	fmt.Fprint(w, j)
}

