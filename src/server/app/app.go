package app

import (
	"alef/config"
	"alef/handlers"
	"alef/middlewares"
	"fmt"
	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strings"
)

type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

func (a *App) Initialize(config *config.Config) {
	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s",
		config.DB.Host,
		config.DB.Username,
		config.DB.DbName,
		config.DB.Password)

	db, err := gorm.Open(postgres.Open(dbUri), &gorm.Config{})
	if err != nil {
		log.Fatal("Could not connect database")
	}

	a.DB = db
	a.Router = mux.NewRouter().UseEncodedPath()
	a.setRouters()
	if strings.ToLower(config.Environment) != "production" {
		a.Router.Use(middlewares.LoggingMiddleware)
	}

	log.Println("Server started with environment:", config.Environment)
}

func (a *App) setRouters() {
	a.Get("/company-info", a.GetCompanyInfo)
	a.Get("/page/{url}", a.GetPageByUrl)
}

func (a *App) Run(host string) {
	log.Println("Listen on", host)
	log.Fatal(http.ListenAndServe(host, a.Router))
}

func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}

func (a *App) GetCompanyInfo(w http.ResponseWriter, r *http.Request) {
	handlers.GetCompanyInfo(a.DB, w)
}

func (a *App) GetPageByUrl(w http.ResponseWriter, r *http.Request) {
	handlers.GetPageByUrl(a.DB, w, r)
}
