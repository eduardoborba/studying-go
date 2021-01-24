package main

import (
	"fmt"
	"log"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// App represents the application
type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

// Initialize function signature
func (a *App) Initialize(user, password, dbname string) {
	connectionString :=
		fmt.Sprintf("host=localhost user=%s password=%s dbname=%s sslmode=disable", user, password, dbname)
	var err error

	a.DB, err = gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// err = a.DB.AutoMigrate(&Article{})
	// if err != nil {
	// 	fmt.Println("Error HandleMigrate:" + err.Error())
	// }

	a.Router = mux.NewRouter()
}

// Run function signature
func (a *App) Run(addr string) {}
