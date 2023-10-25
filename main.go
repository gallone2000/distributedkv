package main

import (
	"distributedkv/controllers"
	"distributedkv/models"
	"distributedkv/schedulers"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	app := models.App{}
	app = app.InitApp()
	router := initRouter(app)
	httpPort := app.Viper.GetString("HTTP_PORT")

	// load nodes URLs from NODES in .env file
	// send requests to /nodes endpoint to fetch the list of active nodes
	schedulers.GetActiveNodesScheduler(&app, true)

	// send requests to /failed-nodes endpoint to fetch the list of failed nodes
	schedulers.GetFailedNodesScheduler(&app, true)

	// here I start the routines that will do few checks every after a predefined interval
	go schedulers.HeartBeatScheduler(&app)
	go schedulers.GetActiveNodesScheduler(&app, false)
	go schedulers.GetFailedNodesScheduler(&app, false)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", httpPort), router))
}

func initRouter(app models.App) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/get/{key}", func(w http.ResponseWriter, r *http.Request) {
		controllers.GetValue(w, r, &app)
	}).Methods(http.MethodGet)
	router.HandleFunc("/set", func(w http.ResponseWriter, r *http.Request) {
		controllers.SetValue(w, r, &app)
	}).Methods(http.MethodPost)
	router.HandleFunc("/confirm", func(w http.ResponseWriter, r *http.Request) {
		controllers.ConfirmSet(w, r, &app)
	}).Methods(http.MethodPost)

	router.HandleFunc("/nodes", func(w http.ResponseWriter, r *http.Request) {
		controllers.GetNodes(w, r, &app)
	}).Methods(http.MethodGet)
	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		controllers.GetHealtStatus(w, r, &app)
	}).Methods(http.MethodGet)
	router.HandleFunc("/heartbeat", func(w http.ResponseWriter, r *http.Request) {
		controllers.GetHeartBeat(w, r, &app)
	}).Methods(http.MethodGet)
	router.HandleFunc("/rejoin", func(w http.ResponseWriter, r *http.Request) {
		controllers.Rejoin(w, r, &app)
	}).Methods(http.MethodPost)
	router.HandleFunc("/failed-nodes", func(w http.ResponseWriter, r *http.Request) {
		controllers.GetFailedNodes(w, r, &app)
	}).Methods(http.MethodGet)

	return router
}
