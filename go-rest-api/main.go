package main

import (
	"fmt"
	"log"
	"encoding/json"
	"net/http"
	"io/ioutil"

	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Article struct {
	gorm.Model
    Title string
    Desc string
    Content string
}

var db *gorm.DB
var err error

func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Welcome to the HomePage!")
    fmt.Println("Endpoint Hit: homePage")
}

func getArticles(w http.ResponseWriter, r *http.Request){
	fmt.Println("Endpoint Hit: getArticles")
	var articles []Article
	
	db.Find(&articles)
    
	json.NewEncoder(w).Encode(&articles)
}

func getArticle(w http.ResponseWriter, r *http.Request){
    params := mux.Vars(r)
	var article Article

	db.First(&article, params["id"])

    json.NewEncoder(w).Encode(&article)
}

func createArticle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: createArticle")
    reqBody, _ := ioutil.ReadAll(r.Body)
    var article Article 
	json.Unmarshal(reqBody, &article)
	
	db.Create(&article)

    json.NewEncoder(w).Encode(article)
}

func deleteArticle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: deleteArticle")
    params := mux.Vars(r)
	var article Article
	
	db.First(&article, params["id"])
	db.Delete(&article)
	  
	json.NewEncoder(w).Encode(&article)
}

func updateArticle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: updateArticle")
	params := mux.Vars(r)
	reqBody, _ := ioutil.ReadAll(r.Body)
	var updatedArticle Article 
	var article Article
	json.Unmarshal(reqBody, &updatedArticle)

	db.First(&article, params["id"])
	db.Model(&article).Updates(updatedArticle)

	json.NewEncoder(w).Encode(article)
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", homePage).Methods("GET")
	router.HandleFunc("/articles", getArticles).Methods("GET")
	router.HandleFunc("/articles", createArticle).Methods("POST")
	router.HandleFunc("/articles/{id}", updateArticle).Methods("PUT")
	router.HandleFunc("/articles/{id}", getArticle).Methods("GET")
	router.HandleFunc("/articles/{id}", deleteArticle).Methods("DELETE")
	
	log.Fatal(http.ListenAndServe(":10000", router))
}

func main() {
	dsn := "host=localhost port=5432 user=eduardoborba dbname=go-rest-api sslmode=disable password=password"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	err = db.AutoMigrate(&Article{})
	if err != nil {
		fmt.Println("Error HandleMigrate:" + err.Error())
	}
	
	handleRequests()	
}
