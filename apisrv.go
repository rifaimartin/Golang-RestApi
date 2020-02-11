package main

import (
       "encoding/json"	
        "log"
       "net/http"
     //   "strconv"
        "github.com/gorilla/mux"
)

type Biodata struct {
	Nama string `json: "nama"`
	Umur int    `json: "umur"`
}

var data []Biodata = []Biodata {
	Biodata {
		Nama : "Rifai Martin",
		Umur : 18,
	},
	Biodata {
		Nama : "Mufidah ",
		Umur : 16,
	},
}

func main() {
	// Point 1
	 var r *mux.Router = mux.NewRouter()
	 log.Println("HTTP Services started @ 8080 ....")
	 
	// Point 2
	r.HandleFunc("/",
	func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		json.NewEncoder(w).Encode(&data)
		 }).Methods("GET")
 
	// Point 3
	 http.ListenAndServe(":8080", r)
 
 }
