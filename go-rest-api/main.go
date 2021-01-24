package main

import (
	"os"
)

// func homePage(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Welcome to the HomePage!")
// 	fmt.Println("Endpoint Hit: homePage")
// }

// func getArticles(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("Endpoint Hit: getArticles")
// 	var articles []Article

// 	a.DB.Find(&articles)

// 	json.NewEncoder(w).Encode(&articles)
// }

// func getArticle(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("Endpoint Hit: getArticle")
// 	params := mux.Vars(r)
// 	var article Article

// 	err := a.DB.First(&article, params["id"])

// 	if err == nil {
// 		json.NewEncoder(w).Encode(&article)
// 	} else {
// 		handleNotFound(w)
// 	}
// }

// func createArticle(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("Endpoint Hit: createArticle")
// 	fmt.Println(r.Body)
// 	reqBody, _ := ioutil.ReadAll(r.Body)
// 	var article Article
// 	json.Unmarshal(reqBody, &article)

// 	result := a.DB.Create(&article)
// 	if result.Error == nil {
// 		json.NewEncoder(w).Encode(&article)
// 	} else {
// 		w.WriteHeader(422)
// 		json.NewEncoder(w).Encode(&err)
// 	}
// }

// func deleteArticle(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("Endpoint Hit: deleteArticle")
// 	params := mux.Vars(r)
// 	var article Article

// 	a.DB.First(&article, params["id"])
// 	a.DB.Delete(&article)

// 	json.NewEncoder(w).Encode(&article)
// }

// func updateArticle(w http.ResponseWriter, r *http.Request) {
// 	fmt.Println("Endpoint Hit: updateArticle")
// 	params := mux.Vars(r)
// 	reqBody, _ := ioutil.ReadAll(r.Body)
// 	var updatedArticle Article
// 	var article Article
// 	json.Unmarshal(reqBody, &updatedArticle)

// 	a.DB.First(&article, params["id"])
// 	a.DB.Model(&article).Updates(updatedArticle)

// 	json.NewEncoder(w).Encode(article)
// }

// func handleNotFound(w http.ResponseWriter) {
// 	w.WriteHeader(404)

// 	data := make(map[string]string)
// 	data["error_code"] = "Not Found"
// 	json.NewEncoder(w).Encode(&data)
// }

// func handleRequests() {
// 	router := mux.NewRouter().StrictSlash(true)
// 	router.HandleFunc("/", homePage).Methods("GET")
// 	router.HandleFunc("/articles", getArticles).Methods("GET")
// 	router.HandleFunc("/articles", createArticle).Methods("POST")
// 	router.HandleFunc("/articles/{id}", updateArticle).Methods("PUT")
// 	router.HandleFunc("/articles/{id}", getArticle).Methods("GET")
// 	router.HandleFunc("/articles/{id}", deleteArticle).Methods("DELETE")

// 	log.Fatal(http.ListenAndServe(":10000", router))
// }

func main() {
	// dsn := "host=localhost port=5432 user=eduardoborba dbname=go-rest-api sslmode=disable password=password"
	// db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	// if err != nil {
	// 	panic("failed to connect database")
	// }

	// err = a.DB.AutoMigrate(&Article{})
	// if err != nil {
	// 	fmt.Println("Error HandleMigrate:" + err.Error())
	// }

	// handleRequests()
	a := App{}
	a.Initialize(
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"))

	a.Run(":10000")
}
