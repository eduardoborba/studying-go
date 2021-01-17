package main

import (
	"fmt"
	"log"
	"encoding/json"
	"net/http"
	"io/ioutil"

	"github.com/go-pg/pg/v10"
    "github.com/go-pg/pg/v10/orm"
	"github.com/gorilla/mux"
)

type Article struct {
	Id      string `json:"Id"`
    Title string `json:"Title"`
    Desc string `json:"desc"`
    Content string `json:"content"`
}

var Articles []Article

func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Welcome to the HomePage!")
    fmt.Println("Endpoint Hit: homePage")
}

func returnAllArticles(w http.ResponseWriter, r *http.Request){
    fmt.Println("Endpoint Hit: returnAllArticles")
    json.NewEncoder(w).Encode(Articles)
}

func returnSingleArticle(w http.ResponseWriter, r *http.Request){
    vars := mux.Vars(r)
    key := vars["id"]

    for _, article := range Articles {
        if article.Id == key {
            json.NewEncoder(w).Encode(article)
        }
    }
}

func createNewArticle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: createNewArticle")
    reqBody, _ := ioutil.ReadAll(r.Body)
    var article Article 
	json.Unmarshal(reqBody, &article)
	
	Articles = append(Articles, article)

    json.NewEncoder(w).Encode(article)
}

func deleteArticle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: deleteArticle")
    vars := mux.Vars(r)
    id := vars["id"]

    for index, article := range Articles {
        if article.Id == id {
            Articles = append(Articles[:index], Articles[index+1:]...)
        }
    }
}

func updateArticle(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint Hit: updateArticle")
	vars := mux.Vars(r)
    id := vars["id"]
	reqBody, _ := ioutil.ReadAll(r.Body)
	var updatedArticle Article 
	json.Unmarshal(reqBody, &updatedArticle)

	for index, article := range Articles {
        if article.Id == id {
			updatedArticle.Id = article.Id
			Articles[index] = updatedArticle
        }
	}
	
	json.NewEncoder(w).Encode(updatedArticle)
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage).Methods("GET")
	myRouter.HandleFunc("/articles", returnAllArticles).Methods("GET")
	myRouter.HandleFunc("/articles", createNewArticle).Methods("POST")
	myRouter.HandleFunc("/articles/{id}", updateArticle).Methods("PUT")
	myRouter.HandleFunc("/articles/{id}", returnSingleArticle).Methods("GET")
	myRouter.HandleFunc("/articles/{id}", deleteArticle).Methods("DELETE")
    log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
	Articles = []Article{
        Article{Id: "1", Title: "Hello", Desc: "Article Description", Content: "Article Content"},
        Article{Id: "2", Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	}
	// initDb()
    handleRequests()	
}

func initDb() {
	db := pg.Connect(&pg.Options{
		User: "eduardoborba",
		Password: "password",
		Database: "go-rest-api",
    })
    defer db.Close()

    err := createSchema(db)
    if err != nil {
        panic(err)
    }
}


func createSchema(db *pg.DB) error {
    models := []interface{}{
        (*Article)(nil),
    }

    for _, model := range models {
        err := db.Model(model).CreateTable(&orm.CreateTableOptions{})
        if err != nil {
            return err
        }
    }
    return nil
}