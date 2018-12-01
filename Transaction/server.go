package main

import (
	//"fmt"
	"net/http"
	"encoding/json"
	//"log"
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var mongodb_server = "localhost"
var mongodb_database = "bookstore"
var mongodb_collection = "transaction"


//var mongodb_server string

//var mongodb_database string

//var mongodb_collection string

// NewServer configures and returns a Server.
func NewServer() *negroni.Negroni {
	formatter := render.New(render.Options{
		IndentJSON: true,
	})
	
	//mongodb_server = os.Getenv("MONGO_SERVER")
	//mongodb_database = os.Getenv("MONGO_DB")
	//mongodb_collection = os.Getenv("MONGO_COLLECTION")
	
	n := negroni.Classic()
	mx := mux.NewRouter()
	initRoutes(mx, formatter)
	n.UseHandler(mx)
	return n
}

// API Routes
func initRoutes(mx *mux.Router, formatter *render.Render) {
	mx.HandleFunc("/createTransaction", transactionCreateHandler(formatter)).Methods("POST")
	mx.HandleFunc("/getAllTransactionByUser/{id}", transactionHandler(formatter)).Methods("GET")
	
}

func transactionCreateHandler(formatter *render.Render) http.HandlerFunc{
	return func(w http.ResponseWriter, req *http.Request) {	
		setDefaultHeaders(w)	
		session, err := mgo.Dial(mongodb_server)
		if err != nil {
			panic(err)
			return
		}
		defer session.Close()
		session.SetMode(mgo.Monotonic, true)
		c := session.DB(mongodb_database).C(mongodb_collection)
		
		decoder := json.NewDecoder(req.Body)

		var transaction Transaction
		err2 := decoder.Decode(&transaction)
		if err2 != nil {
			panic(err2)
		}
		
		tempbooks:=transaction.MyBooks

		tempid := bson.NewObjectId()
		err1 := c.Insert(&Transaction{ID: tempid,MyBooks:tempbooks})

		if err1 != nil {
			panic(err1)
		}
		
		formatter.JSON(w, http.StatusOK, tempid)
		

	 }

}

func transactionHandler(formatter *render.Render) http.HandlerFunc{
	return func(w http.ResponseWriter, req *http.Request) {	
		setDefaultHeaders(w)	
		session, err := mgo.Dial(mongodb_server)
		if err != nil {
			panic(err)
			return
		}
		defer session.Close()
		session.SetMode(mgo.Monotonic, true)
		//c := session.DB(mongodb_database).C(mongodb_collection)
		//h := session.DB(mongodb_database).C(mongodb_collection1)
		// 	var books []Books
		// 	err = c.Find(bson.M{}).All(&books)
		// 	if err != nil {
		// 		log.Fatal(err)
		// 	}
		// fmt.Println("All Books are :", books)
		// formatter.JSON(w, http.StatusOK, books)
		
		//decoder := json.NewDecoder(req.Body)

		

		
		
		

		
	 }

}



func setDefaultHeaders(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	w.Header().Set("Cache-Control", "no-store, no-cache, must-revalidate, post-check=0, pre-check=0")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Vary", "Accept-Encoding")
}