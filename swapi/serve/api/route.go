package api

import (
	"fmt"
	"os"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
	"github.com/urfave/negroni"
	"serve/database/db"
)

func NewServer() *negroni.Negroni {

	formatter := render.New(render.Options{
		Directory:  "templates",
		Extensions: []string{".html"},
		IndentJSON: true,
	})

	n := negroni.Classic()
	mx := mux.NewRouter()

	initRoutes(mx, formatter)

	n.UseHandler(mx)
	return n
}

func initRoutes(mx *mux.Router, formatter *render.Render) {
	webRoot := os.Getenv("WEBROOT")
	if len(webRoot) == 0 {
		if root, err := os.Getwd(); err != nil {
			panic("Could not retrive working directory")
		} else {
			webRoot = root
			fmt.Println(root)
		}
	}
	db.Start("database/db/swa.db")
	mx.HandleFunc("/api/people", peopleHandler(formatter)).Methods("GET")
	mx.HandleFunc("/api/people/{id:[0-9]+}", peopleIdHandler).Methods("GET")
	mx.HandleFunc("/api/films", filmsHandler(formatter)).Methods("GET")
	mx.HandleFunc("/api/films/{id:[0-9]+}", filmsIdHandler).Methods("GET")
	mx.HandleFunc("/api/planets", planetsHandler(formatter)).Methods("GET")
	mx.HandleFunc("/api/planets/{id:[0-9]+}", planetsIdHandler).Methods("GET")
	mx.HandleFunc("/api/species", speciesHandler(formatter)).Methods("GET")
	mx.HandleFunc("/api/species/{id:[0-9]+}", speciesIdHandler).Methods("GET")
	mx.HandleFunc("/api/starships", starshipsHandler(formatter)).Methods("GET")
	mx.HandleFunc("/api/starships/{id:[0-9]+}", starshipsIdHandler).Methods("GET")
	mx.HandleFunc("/api/vehicles", vehiclesHandler(formatter)).Methods("GET")
	mx.HandleFunc("/api/vehicles/{id:[0-9]+}", vehiclesIdHandler).Methods("GET")
}
